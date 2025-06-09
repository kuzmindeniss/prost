package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kuzmindeniss/prost/internal"
	"github.com/kuzmindeniss/prost/internal/db"
	"github.com/kuzmindeniss/prost/internal/messaging"
	"github.com/kuzmindeniss/prost/internal/tg"
	"github.com/kuzmindeniss/prost/internal/tg_notifications"
)

var commands = []tgbotapi.BotCommand{
	{
		Command:     "start",
		Description: "Начать работу с ботом уведомлений",
	},
	{
		Command:     "subscribe",
		Description: "Подписаться на уведомления",
	},
	{
		Command:     "unsubscribe",
		Description: "Отписаться от уведомлений",
	},
	{
		Command:     "status",
		Description: "Текущий статус подписки",
	},
}

func init() {
	internal.LoadEnv()
	db.ConnectToDb()
}

func main() {
	err := messaging.InitRabbitMQ()
	if err != nil {
		panic(fmt.Sprintf("Warning: Failed to initialize RabbitMQ: %v", err))
	} else {
		defer messaging.CloseRabbitMQ()

		ctx := context.Background()
		err = messaging.ConsumeApplicationCreated(ctx, tg_notifications.HandleApplicationCreated)
		if err != nil {
			log.Printf("Warning: Failed to start consuming application created messages: %v", err)
		}
	}

	notificationBot := tg.InitBot(os.Getenv("NOTIFICATION_BOT_TOKEN"))

	config := tgbotapi.NewSetMyCommands(commands...)
	if _, err := notificationBot.Request(config); err != nil {
		panic(err)
	}

	notificationBot.Debug = true
	log.Printf("Notification bot authorized on account %s", notificationBot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := notificationBot.GetUpdatesChan(u)

	var wg sync.WaitGroup

	for update := range updates {
		wg.Add(1)
		go tg_notifications.HandleUpdate(notificationBot, &update, &wg)
	}

	wg.Wait()
}
