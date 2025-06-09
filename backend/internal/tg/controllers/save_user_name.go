package controllers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kuzmindeniss/prost/internal/db"
	"github.com/kuzmindeniss/prost/internal/db/repository"
	"github.com/kuzmindeniss/prost/internal/tg/helpers"
	"github.com/kuzmindeniss/prost/internal/tg/messages"
)

func SaveUserName(bot *tgbotapi.BotAPI, update *tgbotapi.Update, userFromDb *repository.GetUserTgRow) {
	userId := helpers.GetUserId(update)
	isUserLogged := userFromDb != nil && userFromDb.ID != 0
	newName := update.Message.Text

	if !isUserLogged {
		db.Repo.CreateUserTg(context.Background(), repository.CreateUserTgParams{
			ID:         userId,
			Name:       newName,
			TgUsername: update.Message.From.UserName,
		})
	} else {
		db.Repo.UpdateUserTgName(context.Background(), repository.UpdateUserTgNameParams{
			Name: newName,
			ID:   userId,
		})
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, messages.UserNameSaved)
	helpers.SendMessage(bot, &msg)
}
