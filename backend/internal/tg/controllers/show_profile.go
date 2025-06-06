package controllers

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kuzmindeniss/prost/internal/db/repository"
	"github.com/kuzmindeniss/prost/internal/tg/helpers"
)

func ShowProfile(bot *tgbotapi.BotAPI, update *tgbotapi.Update, user repository.GetUserTgRow) {
	// Create message with profile information
	profileText := fmt.Sprintf(
		"📋 *Ваш профиль*\n\n"+
			"👤 *Имя:* %s\n"+
			"🏢 *Подразделение:* %s",
		tgbotapi.EscapeText("MarkdownV2", user.UserName),
		tgbotapi.EscapeText("MarkdownV2", user.UnitName.String),
	)

	changeNameButton := tgbotapi.NewInlineKeyboardButtonData("Изменить имя", "change_name")
	nameRow := tgbotapi.NewInlineKeyboardRow(changeNameButton)

	changeUnitButton := tgbotapi.NewInlineKeyboardButtonData("Изменить подразделение", "change_unit")
	unitRow := tgbotapi.NewInlineKeyboardRow(changeUnitButton)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, profileText)
	msg.ParseMode = "MarkdownV2"

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(nameRow, unitRow)

	helpers.LastBotReplyMsg[user.ID] = helpers.SendMessage(bot, &msg)
}
