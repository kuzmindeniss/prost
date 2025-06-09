package notifications

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kuzmindeniss/prost/internal/db/repository"
)

var notificationBot *tgbotapi.BotAPI
var adminChatIDs []int64

// InitNotificationBot initializes the notification bot
func InitNotificationBot() {
	// Get notification bot token from environment
	token := os.Getenv("NOTIFICATION_BOT_TOKEN")
	if token == "" {
		log.Println("Warning: NOTIFICATION_BOT_TOKEN not set, notifications will be disabled")
		return
	}

	// Parse admin chat IDs from environment
	// Format: ADMIN_CHAT_IDS=123456789,987654321
	adminChatIDsStr := os.Getenv("ADMIN_CHAT_IDS")
	if adminChatIDsStr == "" {
		log.Println("Warning: ADMIN_CHAT_IDS not set, notifications will have nowhere to go")
	} else {
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

		log.Printf("Configured %d admin chat IDs for notifications", len(adminChatIDs))
	}

	// Initialize bot if token is provided
	var err error
	notificationBot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Printf("Error initializing notification bot: %v", err)
		return
	}

	log.Printf("Notification bot authorized: @%s", notificationBot.Self.UserName)
}

// SendNewApplicationNotification sends a notification about a new application
func SendNewApplicationNotification(application repository.Application, user *repository.GetUserTgRow) {
	if notificationBot == nil {
		log.Println("Notification bot not initialized, skipping notification")
		return
	}

	if len(adminChatIDs) == 0 {
		log.Println("No admin chat IDs configured, skipping notification")
		return
	}

	// Format message
	message := fmt.Sprintf(
		"🆕 *Новая заявка*\n\n"+
			"*Текст:* %s\n\n"+
			"*От:* %s\n"+
			"*Подразделение:* %s\n"+
			"*ID заявки:* `%s`",
		tgbotapi.EscapeText("MarkdownV2", application.Text),
		tgbotapi.EscapeText("MarkdownV2", user.UserName),
		tgbotapi.EscapeText("MarkdownV2", user.UnitName.String),
		tgbotapi.EscapeText("MarkdownV2", application.ID.String()),
	)

	// Send to all admin chats
	for _, chatID := range adminChatIDs {
		msg := tgbotapi.NewMessage(chatID, message)
		msg.ParseMode = "MarkdownV2"

		_, err := notificationBot.Send(msg)
		if err != nil {
			log.Printf("Error sending notification to chat %d: %v", chatID, err)
		}
	}
}
