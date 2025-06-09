package callbacks

import (
	"context"
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kuzmindeniss/prost/internal/db"
	"github.com/kuzmindeniss/prost/internal/db/repository"
	"github.com/kuzmindeniss/prost/internal/messaging"
	"github.com/kuzmindeniss/prost/internal/tg/helpers"
)

func SaveApplication(bot *tgbotapi.BotAPI, update *tgbotapi.Update, user *repository.GetUserTgRow) {
	applicationText := helpers.GetApplicationTextFromDraft(update.CallbackQuery.Message.Text)

	// Save application to database
	application, err := db.Repo.CreateApplication(context.Background(), repository.CreateApplicationParams{
		Text:     applicationText,
		UnitID:   user.UnitID,
		UserTgID: pgtype.Int8{Int64: user.ID, Valid: true},
	})

	if err != nil {
		panic(err)
	}

	msgText := fmt.Sprintf(
		"*Заявка отправлена:* __%s__ *В подразделение:*  __%s__",
		tgbotapi.EscapeText("MarkdownV2", applicationText),
		tgbotapi.EscapeText("MarkdownV2", user.UnitName.String),
	)

	msg := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, msgText)

	msg.ParseMode = "MarkdownV2"

	bot.Send(msg)

	// Publish application created message to RabbitMQ
	ctx := context.Background()
	createTime := ""
	if application.CreatedAt.Valid {
		createTime = application.CreatedAt.Time.Format(time.RFC3339)
	}

	appMessage := messaging.ApplicationMessage{
		ID:         application.ID.String(),
		Text:       application.Text,
		UserName:   user.UserName,
		UnitName:   user.UnitName.String,
		CreateTime: createTime,
	}

	// Publish asynchronously to not block the user
	go func() {
		err := messaging.PublishApplicationCreated(ctx, appMessage)
		if err != nil {
			// Just log the error, don't fail the application save
			fmt.Printf("Error publishing application created message: %v\n", err)
		}
	}()
}
