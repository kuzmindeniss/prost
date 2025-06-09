package tg_notifications

import (
	"context"
	"log"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kuzmindeniss/prost/internal/db"
	"github.com/kuzmindeniss/prost/internal/db/repository"
)

func HandleUpdate(bot *tgbotapi.BotAPI, update *tgbotapi.Update, wg *sync.WaitGroup) {
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
