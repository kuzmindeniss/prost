package main

import (
	"log"
	"os"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/kuzmindeniss/prost/internal/tg"
	"github.com/kuzmindeniss/prost/internal/tg/initializers"
)

type BotCommand struct {
	Command string
	Handler func(bot *tgbotapi.BotAPI, update *tgbotapi.Update)
}

var commands = []tgbotapi.BotCommand{
	{
		Command:     "start",
		Description: "Начать общение с ботом",
	},
	{
		Command:     "send_application",
		Description: "Отправить заявку",
	},
	{
		Command:     "profile",
		Description: "Редактирование профиля",
	},
}

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDb()
}

func main() {
	godotenv.Load()

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	config := tgbotapi.NewSetMyCommands(commands...)
	if _, err := bot.Request(config); err != nil {
		panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	var wg sync.WaitGroup

	for update := range updates {
		wg.Add(1)
		go tg.HandleUpdate(bot, &update, &wg)
	}

	wg.Wait()
}
