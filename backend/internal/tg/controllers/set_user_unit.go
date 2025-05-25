package controllers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kuzmindeniss/prost/internal/db/repository"
	"github.com/kuzmindeniss/prost/internal/tg/helpers"
	"github.com/kuzmindeniss/prost/internal/tg/initializers"
	"github.com/kuzmindeniss/prost/internal/tg/messages"
)

func SetUserUnit(bot *tgbotapi.BotAPI, update *tgbotapi.Update, user *repository.GetUserTgRow, unit *repository.Unit) {
	initializers.Repo.UpdateUserUnitID(context.Background(), repository.UpdateUserUnitIDParams{
		UnitID: unit.ID,
		UserID: user.ID,
	})

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, messages.UnitSet)
	helpers.SendMessage(bot, &msg)
}
