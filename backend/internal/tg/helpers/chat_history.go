package helpers

import (
	"regexp"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ChatID = int64

var LastBotReplyMsg = make(map[ChatID]*tgbotapi.Message)

func SendMessage(bot *tgbotapi.BotAPI, msg *tgbotapi.MessageConfig) *tgbotapi.Message {
	sentMsg, err := bot.Send(msg)
	if err != nil {
		return nil
	}

	LastBotReplyMsg[msg.ChatID] = &sentMsg
	return &sentMsg
}

func GetApplicationTextFromDraft(draft string) string {
	pattern := `(?s)Подтвердите отправку заявки:.*?\n(.*?)\n\n👷 Отправил:`

	// Compile and apply the regex
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(draft)

	var originalText string
	if len(matches) > 1 {
		originalText = strings.TrimSpace(matches[1])
	}
	return originalText
}
