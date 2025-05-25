package helpers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetUserId(update *tgbotapi.Update) int64 {
	if update.Message != nil {
		return update.Message.From.ID
	}

	if update.CallbackQuery != nil {
		return update.CallbackQuery.From.ID
	}

	return 0
}
