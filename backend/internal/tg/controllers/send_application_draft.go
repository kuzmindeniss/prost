package controllers

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kuzmindeniss/prost/internal/db/repository"
	"github.com/kuzmindeniss/prost/internal/tg/helpers"
)

func SendApplicationDraft(bot *tgbotapi.BotAPI, update *tgbotapi.Update, user *repository.GetUserTgRow) {
	text := fmt.Sprintf(`
🛠️ *Подтвердите отправку заявки:* 🛠️
__%s__

👷 *Отправил:*
__%s__

🏗️ *Подразделение:*
__%s__
`,
		tgbotapi.EscapeText("MarkdownV2", update.Message.Text),
		tgbotapi.EscapeText("MarkdownV2", user.UserName),
		tgbotapi.EscapeText("MarkdownV2", user.UnitName.String),
	)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	msg.ParseMode = "MarkdownV2"

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Отправить", "save_application"),
			tgbotapi.NewInlineKeyboardButtonData("Отменить", "cancel_application"),
		),
	)

	helpers.SendMessage(bot, &msg)
}
