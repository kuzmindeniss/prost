package controllers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kuzmindeniss/prost/internal/tg/helpers"
	"github.com/kuzmindeniss/prost/internal/tg/messages"
)

func SendSetUserUnitRequest(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, messages.SetUserUnit)

	buttons := make([][]tgbotapi.InlineKeyboardButton, 0)
	for _, unit := range helpers.GetAllUnits() {
		buttons = append(
			buttons,
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(unit.Name, "save_unit:"+unit.ID.String()),
			),
		)
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(buttons...)

	helpers.SendMessage(bot, &msg)
}
