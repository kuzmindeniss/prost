package callbacks

import (
	"context"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"github.com/kuzmindeniss/prost/internal/db/repository"
	"github.com/kuzmindeniss/prost/internal/tg/helpers"
	"github.com/kuzmindeniss/prost/internal/tg/initializers"
	"github.com/kuzmindeniss/prost/internal/tg/messages"
)

func SaveUnit(bot *tgbotapi.BotAPI, update *tgbotapi.Update, user *repository.GetUserTgRow) {
	unitId := strings.Split(update.CallbackQuery.Data, ":")[1]
	initializers.Repo.UpdateUserUnitID(context.Background(), repository.UpdateUserUnitIDParams{
		UserID: user.ID,
		UnitID: uuid.MustParse(unitId),
	})

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, messages.UnitSaved)

	helpers.SendMessage(bot, &msg)
}
