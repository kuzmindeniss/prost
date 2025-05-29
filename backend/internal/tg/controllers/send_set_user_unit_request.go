package controllers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kuzmindeniss/prost/internal/tg/helpers"
	"github.com/kuzmindeniss/prost/internal/tg/messages"
)

func SendSetUserUnitRequest(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	var chatID int64

	if update.Message != nil {
		chatID = update.Message.Chat.ID
	} else {
		chatID = update.CallbackQuery.Message.Chat.ID
	}

	msg := tgbotapi.NewMessage(chatID, messages.SetUserUnit)

	buttons := make([][]tgbotapi.InlineKeyboardButton, 0)
	for _, unit := range helpers.GetAllUnits() {
		buttons = append(
			buttons,
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(unit.Name, "save_unit:"+unit.ID.String()),
			),
		)
	}

	if len(buttons) > 0 {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(buttons...)
	} else {
		msg.Text = "Нет доступных подразделений"
	}

	helpers.SendMessage(bot, &msg)
}
