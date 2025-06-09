package tg

import (
	"context"
	"strings"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"github.com/kuzmindeniss/prost/internal/db"
	"github.com/kuzmindeniss/prost/internal/tg/callbacks"
	"github.com/kuzmindeniss/prost/internal/tg/controllers"
	"github.com/kuzmindeniss/prost/internal/tg/helpers"
	"github.com/kuzmindeniss/prost/internal/tg/messages"
)

func HandleUpdate(bot *tgbotapi.BotAPI, update *tgbotapi.Update, wg *sync.WaitGroup) {
	defer wg.Done()

	userId := helpers.GetUserId(update)

	user, err := db.Repo.GetUserTg(context.Background(), userId)
	isUserLogged := err == nil && user.ID != 0

	isUnitChosen := user.UnitID != uuid.Nil

	// Callbacks
	if update.CallbackQuery != nil {
		if update.CallbackQuery.Data == "save_application" {
			callbacks.SaveApplication(bot, update, &user)
		}

		if update.CallbackQuery.Data == "cancel_application" {
			callbacks.CancelApplication(bot, update, &user)
		}

		if update.CallbackQuery.Data == "change_name" {
			controllers.SendSetUserNameRequest(bot, update)
		}

		if update.CallbackQuery.Data == "change_unit" {
			controllers.SendSetUserUnitRequest(bot, update)
		}

		if strings.HasPrefix(update.CallbackQuery.Data, "save_unit:") {
			callbacks.SaveUnit(bot, update, &user)
		}

		return
	}

	// Commands
	isStartCommand := update.Message.Text == "/start"
	isSendApplicationCommand := update.Message.Text == "/send_application"
	isProfileCommand := update.Message.Text == "/profile"

	isCommand := isStartCommand || isSendApplicationCommand || isProfileCommand

	// User logging
	isResponseOnSetUserNameRequest := helpers.LastBotReplyMsg[userId] != nil && helpers.LastBotReplyMsg[userId].Text == messages.SetUserName

	if isResponseOnSetUserNameRequest && !isCommand {
		controllers.SaveUserName(bot, update, &user)
		if user.UnitName.String == "" {
			controllers.SendSetUserUnitRequest(bot, update)
		}
		return
	}

	if !isUserLogged {
		controllers.SendSetUserNameRequest(bot, update)
		return
	}

	if !isUnitChosen {
		controllers.SendSetUserUnitRequest(bot, update)
		return
	}

	// Commands handling
	if isStartCommand || isSendApplicationCommand {
		controllers.SendApplicationRequest(bot, update, &user)
		return
	}

	if isProfileCommand {
		controllers.ShowProfile(bot, update, user)
		return
	}

	isResponseOnApplicationRequest := helpers.LastBotReplyMsg[userId] != nil &&
		strings.Contains(helpers.LastBotReplyMsg[userId].Text, messages.SendApplication)

	if isResponseOnApplicationRequest {
		controllers.SendApplicationDraft(bot, update, &user)
		return
	}

	controllers.SendApplicationRequest(bot, update, &user)
}
