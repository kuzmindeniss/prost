package controllers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kuzmindeniss/prost/internal/db/repository"
	"github.com/kuzmindeniss/prost/internal/tg/helpers"
	"github.com/kuzmindeniss/prost/internal/tg/messages"
)

func SendApplicationRequest(bot *tgbotapi.BotAPI, update *tgbotapi.Update, user *repository.GetUserTgRow) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, messages.SendApplication)
	helpers.SendMessage(bot, &msg)
}
