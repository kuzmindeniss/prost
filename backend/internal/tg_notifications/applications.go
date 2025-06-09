package tg_notifications

import (
	"context"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kuzmindeniss/prost/internal/db"
	"github.com/kuzmindeniss/prost/internal/messaging"
	"github.com/kuzmindeniss/prost/internal/tg"
)

func HandleApplicationCreated(msg messaging.ApplicationMessage) error {
	if tg.Bot == nil {
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

		_, err := tg.Bot.Send(tgMsg)
		if err != nil {
			log.Printf("Error sending notification to chat %d: %v", chatID, err)
			// Continue sending to other chats even if one fails
		}
	}

	return nil
}
