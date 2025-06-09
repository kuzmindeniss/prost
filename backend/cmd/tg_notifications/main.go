package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kuzmindeniss/prost/internal/db"
	"github.com/kuzmindeniss/prost/internal/db/repository"
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

var notificationBot *tgbotapi.BotAPI

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

// handleApplicationCreated is called when a new application created message is received from RabbitMQ
func handleApplicationCreated(msg messaging.ApplicationMessage) error {
	if notificationBot == nil {
		return fmt.Errorf("notification bot not initialized")
	}

	ids, err := db.Repo.GetUserNotificationsTgIds(context.Background())
	if err != nil {
		panic(fmt.Sprintf("Users extracting error: %v", err))
	}

	if len(ids) == 0 {
		log.Println("Warning: No users to send notification, can't send notification but acknowledging message")
		return nil
	}

	// Format message
	message := fmt.Sprintf(
		"🆕 *Новая заявка*\n\n"+
			"*Текст:* %s\n\n"+
			"*От:* %s\n"+
			"*Подразделение:* %s\n",
		tgbotapi.EscapeText("MarkdownV2", msg.Text),
		tgbotapi.EscapeText("MarkdownV2", msg.UserName),
		tgbotapi.EscapeText("MarkdownV2", msg.UnitName),
	)

	// Send to all admin chats
	for _, chatID := range ids {
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
		_, err := db.Repo.CreateUserNotificationTg(context.Background(), repository.CreateUserNotificationTgParams{
			ID: chatID,
		})

		if err != nil {
			log.Printf("Error while subscribing user: %v", err)
			msg := tgbotapi.NewMessage(chatID, "Произошла ошибка при подписке пользователя")
			bot.Send(msg)
			return
		}

		msg := tgbotapi.NewMessage(chatID, "✅ Вы успешно подписаны на уведомления о новых заявках!")
		bot.Send(msg)

	case "unsubscribe":
		err := db.Repo.DeleteUserNotificationsTg(context.Background(), chatID)
		if err != nil {
			log.Printf("Error while unsubscribing user: %v", err)
			msg := tgbotapi.NewMessage(chatID, "Произошла ошибка при отписке пользователя")
			bot.Send(msg)
			return
		}

		msg := tgbotapi.NewMessage(chatID, "❌ Вы отписаны от уведомлений.")
		bot.Send(msg)

	case "status":
		user, err := db.Repo.GetUserNotificationsTg(context.Background(), chatID)
		if err != nil || user.ID == 0 {
			msg := tgbotapi.NewMessage(chatID, "❌ Вы не подписаны на уведомления.")
			bot.Send(msg)
			return
		}

		msg := tgbotapi.NewMessage(chatID, "✅ Вы подписаны на уведомления о новых заявках.")
		bot.Send(msg)
		return
	}
}
