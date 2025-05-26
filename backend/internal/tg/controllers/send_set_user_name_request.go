package controllers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kuzmindeniss/prost/internal/tg/helpers"
	"github.com/kuzmindeniss/prost/internal/tg/messages"
)

func SendSetUserNameRequest(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	var chatID int64

	if update.Message != nil {
		chatID = update.Message.Chat.ID
	} else {
		chatID = update.CallbackQuery.Message.Chat.ID
	}

	msg := tgbotapi.NewMessage(chatID, messages.SetUserName)

	helpers.SendMessage(bot, &msg)
}
