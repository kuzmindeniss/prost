package helpers

import (
	"testing"
)

func TestGetApplicationTextFromDraft(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "standard format",
			input: `
🛠️ *Подтвердите отправку заявки:* 🛠️
__Нужно починить дверь__

👷 *Отправил:*
__Иван Иванов__

🏗️ *Подразделение:*
__Техотдел__
`,
			expected: "Нужно починить дверь",
		},
		{
			name: "multiline application text",
			input: `
🛠️ *Подтвердите отправку заявки:* 🛠️
__Нужно починить дверь
И заменить лампочку
И покрасить стену__

👷 *Отправил:*
__Иван Иванов__

🏗️ *Подразделение:*
__Техотдел__
`,
			expected: "Нужно починить дверь\nИ заменить лампочку\nИ покрасить стену",
		},
		{
			name: "with additional markdown",
			input: `
🛠️ *Подтвердите отправку заявки:* 🛠️
__*Срочно* нужно починить дверь__

👷 *Отправил:*
__Иван Иванов__

🏗️ *Подразделение:*
__Техотдел__
`,
			expected: "*Срочно* нужно починить дверь",
		},
		{
			name: "missing closing marker",
			input: `
🛠️ *Подтвердите отправку заявки:* 🛠️
__Нужно починить дверь__

👷 *Отправил:*
`,
			expected: "Нужно починить дверь",
		},
		{
			name: "missing opening marker",
			input: `
__Нужно починить дверь__

👷 *Отправил:*
__Иван Иванов__
`,
			expected: "Нужно починить дверь",
		},
		{
			name: "without underscores",
			input: `
🛠️ *Подтвердите отправку заявки:* 🛠️
Нужно починить дверь

👷 *Отправил:*
__Иван Иванов__
`,
			expected: "Нужно починить дверь",
		},
		{
			name:     "completely different format",
			input:    "Some random text that doesn't match the format",
			expected: "Some random text that doesn't match the format",
		},
		{
			name:     "empty input",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetApplicationTextFromDraft(tt.input)
			if result != tt.expected {
				t.Errorf("GetApplicationTextFromDraft(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
