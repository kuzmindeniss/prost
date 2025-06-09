package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kuzmindeniss/prost/internal/db"
	"github.com/kuzmindeniss/prost/internal/messaging"
	"github.com/kuzmindeniss/prost/internal/tg/initializers"
)

var commands = []tgbotapi.BotCommand{
	{
		Command:     "start",
		Description: "Начать работу с ботом уведомлений",
	},
	{
		Command:     "subscribe",
		Description: "Подписаться на уведомления",
	},
	{
		Command:     "unsubscribe",
		Description: "Отписаться от уведомлений",
	},
	{
		Command:     "status",
		Description: "Текущий статус подписки",
	},
}

// Global bot instance for use in message handlers
var notificationBot *tgbotapi.BotAPI

// adminChatIDs keeps track of all chat IDs that should receive notifications
var adminChatIDs []int64

func init() {
	initializers.LoadEnv()
	db.ConnectToDb()
}

func main() {
	// Initialize RabbitMQ
	err := messaging.InitRabbitMQ()
	if err != nil {
		panic(fmt.Sprintf("Warning: Failed to initialize RabbitMQ: %v", err))
	} else {
		defer messaging.CloseRabbitMQ()

		// Start consuming application created messages
		ctx := context.Background()
		err = messaging.ConsumeApplicationCreated(ctx, handleApplicationCreated)
		if err != nil {
			log.Printf("Warning: Failed to start consuming application created messages: %v", err)
		}
	}

	// Initialize notification bot
	notificationBot, err = tgbotapi.NewBotAPI(os.Getenv("NOTIFICATION_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	// Load admin chat IDs from environment
	loadAdminChatIDs()

	// Set bot commands
	config := tgbotapi.NewSetMyCommands(commands...)
	if _, err := notificationBot.Request(config); err != nil {
		panic(err)
	}

	notificationBot.Debug = true
	log.Printf("Notification bot authorized on account %s", notificationBot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := notificationBot.GetUpdatesChan(u)

	var wg sync.WaitGroup

	for update := range updates {
		wg.Add(1)
		go handleUpdate(notificationBot, &update, &wg)
	}

	wg.Wait()
}

// loadAdminChatIDs loads admin chat IDs from environment variable
func loadAdminChatIDs() {
	adminChatIDsStr := os.Getenv("ADMIN_CHAT_IDS")
	if adminChatIDsStr == "" {
		log.Println("Warning: ADMIN_CHAT_IDS not set, notifications will have nowhere to go")
		return
	}

	// Parse comma-separated chat IDs
	idStrings := strings.Split(adminChatIDsStr, ",")
	for _, idStr := range idStrings {
		idStr = strings.TrimSpace(idStr)
		if idStr == "" {
			continue
		}

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			log.Printf("Error parsing chat ID %s: %v", idStr, err)
			continue
		}

		adminChatIDs = append(adminChatIDs, id)
	}

	log.Printf("Loaded %d admin chat IDs for notifications", len(adminChatIDs))
}

// handleApplicationCreated is called when a new application created message is received from RabbitMQ
func handleApplicationCreated(msg messaging.ApplicationMessage) error {
	fmt.Printf("%+v\n", adminChatIDs)
	fmt.Printf("%+v\n", msg)

	if notificationBot == nil {
		return fmt.Errorf("notification bot not initialized")
	}

	if len(adminChatIDs) == 0 {
		log.Println("Warning: No admin chat IDs configured, can't send notification but acknowledging message")
		return nil
	}

	// Format message
	message := fmt.Sprintf(
		"🆕 *Новая заявка*\n\n"+
			"*Текст:* %s\n\n"+
			"*От:* %s\n"+
			"*Подразделение:* %s\n"+
			"*ID заявки:* `%s`",
		tgbotapi.EscapeText("MarkdownV2", msg.Text),
		tgbotapi.EscapeText("MarkdownV2", msg.UserName),
		tgbotapi.EscapeText("MarkdownV2", msg.UnitName),
		tgbotapi.EscapeText("MarkdownV2", msg.ID),
	)

	// Send to all admin chats
	for _, chatID := range adminChatIDs {
		tgMsg := tgbotapi.NewMessage(chatID, message)
		tgMsg.ParseMode = "MarkdownV2"

		_, err := notificationBot.Send(tgMsg)
		if err != nil {
			log.Printf("Error sending notification to chat %d: %v", chatID, err)
			// Continue sending to other chats even if one fails
		}
	}

	return nil
}

func handleUpdate(bot *tgbotapi.BotAPI, update *tgbotapi.Update, wg *sync.WaitGroup) {
	defer wg.Done()

	if update.Message == nil {
		return
	}

	// Get the chat ID
	chatID := update.Message.Chat.ID

	// Handle commands
	switch update.Message.Command() {
	case "start":
		msg := tgbotapi.NewMessage(chatID, "Добро пожаловать в бот уведомлений!\n\nИспользуйте /subscribe чтобы подписаться на уведомления о новых заявках.")
		bot.Send(msg)

	case "subscribe":
		// Add this chat to the subscribers
		adminChatIDsStr := os.Getenv("ADMIN_CHAT_IDS")
		chatIDStr := strconv.FormatInt(chatID, 10)

		if adminChatIDsStr == "" {
			os.Setenv("ADMIN_CHAT_IDS", chatIDStr)
			adminChatIDs = []int64{chatID}
		} else if !contains(adminChatIDsStr, chatIDStr) {
			newAdminChatIDsStr := adminChatIDsStr + "," + chatIDStr
			os.Setenv("ADMIN_CHAT_IDS", newAdminChatIDsStr)
			adminChatIDs = append(adminChatIDs, chatID)
		}

		msg := tgbotapi.NewMessage(chatID, "✅ Вы успешно подписаны на уведомления о новых заявках!")
		bot.Send(msg)

	case "unsubscribe":
		// Remove this chat from subscribers
		removeAdminChatID(chatID)

		msg := tgbotapi.NewMessage(chatID, "❌ Вы отписаны от уведомлений.")
		bot.Send(msg)

	case "status":
		// Check if this chat is in the subscribers list
		isSubscribed := false
		for _, id := range adminChatIDs {
			if id == chatID {
				isSubscribed = true
				break
			}
		}

		if isSubscribed {
			msg := tgbotapi.NewMessage(chatID, "✅ Вы подписаны на уведомления о новых заявках.")
			bot.Send(msg)
		} else {
			msg := tgbotapi.NewMessage(chatID, "❌ Вы не подписаны на уведомления.")
			bot.Send(msg)
		}
	}
}

// removeAdminChatID removes a chat ID from the admin chat IDs list
func removeAdminChatID(chatID int64) {
	// Create a new slice with the removed chat ID
	var newAdminChatIDs []int64
	for _, id := range adminChatIDs {
		if id != chatID {
			newAdminChatIDs = append(newAdminChatIDs, id)
		}
	}
	adminChatIDs = newAdminChatIDs

	// Update environment variable
	var idStrings []string
	for _, id := range adminChatIDs {
		idStrings = append(idStrings, strconv.FormatInt(id, 10))
	}
	os.Setenv("ADMIN_CHAT_IDS", strings.Join(idStrings, ","))
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return strconv.Itoa(indexOf(s, substr)) != "-1"
}

// Helper function to find the index of a substring
func indexOf(s, substr string) int {
	for i := 0; i < len(s); i++ {
		if i+len(substr) <= len(s) && s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
