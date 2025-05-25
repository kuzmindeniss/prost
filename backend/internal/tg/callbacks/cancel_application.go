package callbacks

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kuzmindeniss/prost/internal/db/repository"
)

func CancelApplication(bot *tgbotapi.BotAPI, update *tgbotapi.Update, user *repository.GetUserTgRow) {
	editMsg := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, "Заявка отменена")

	bot.Send(editMsg)
}
