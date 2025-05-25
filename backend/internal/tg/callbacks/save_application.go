package callbacks

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kuzmindeniss/prost/internal/db/repository"
	"github.com/kuzmindeniss/prost/internal/tg/helpers"
	"github.com/kuzmindeniss/prost/internal/tg/initializers"
)

func SaveApplication(bot *tgbotapi.BotAPI, update *tgbotapi.Update, user *repository.GetUserTgRow) {
	applicationText := helpers.GetApplicationTextFromDraft(update.CallbackQuery.Message.Text)

	// Save application to database
	_, err := initializers.Repo.CreateApplication(context.Background(), repository.CreateApplicationParams{
		Text:     applicationText,
		UnitID:   user.UnitID,
		UserTgID: pgtype.Int8{Int64: user.ID, Valid: true},
	})

	if err != nil {
		panic(err)
	}

	msgText := fmt.Sprintf("*Заявка отправлена:* __%s__ *В подразделение:*  __%s__", applicationText, user.UnitName.String)

	msg := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, msgText)

	msg.ParseMode = "MarkdownV2"

	bot.Send(msg)
}
