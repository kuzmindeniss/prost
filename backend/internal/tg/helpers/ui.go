package helpers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// CreateProfileButtonRows generates two rows of inline keyboard buttons for profile actions:
// 1. Row with button to change name
// 2. Row with button to change unit
func CreateProfileButtonRows() [][]tgbotapi.InlineKeyboardButton {
	// Create first row with "Change Name" button
	changeNameButton := tgbotapi.NewInlineKeyboardButtonData("Изменить имя", "change_name")
	nameRow := tgbotapi.NewInlineKeyboardRow(changeNameButton)

	// Create second row with "Change Unit" button
	changeUnitButton := tgbotapi.NewInlineKeyboardButtonData("Изменить подразделение", "change_unit")
	unitRow := tgbotapi.NewInlineKeyboardRow(changeUnitButton)

	// Return both rows as a slice
	return [][]tgbotapi.InlineKeyboardButton{nameRow, unitRow}
}

// CreateProfileKeyboard creates an inline keyboard with profile action buttons
func CreateProfileKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(CreateProfileButtonRows()...)
}
