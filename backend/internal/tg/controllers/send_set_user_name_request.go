package controllers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kuzmindeniss/prost/internal/tg/helpers"
	"github.com/kuzmindeniss/prost/internal/tg/messages"
)

func SendSetUserNameRequest(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, messages.SetUserName)

	helpers.SendMessage(bot, &msg)
}
