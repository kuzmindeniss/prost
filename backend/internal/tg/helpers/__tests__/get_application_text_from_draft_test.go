package helpers

import (
	"testing"

	"github.com/kuzmindeniss/prost/internal/tg/helpers"
)

func TestGetApplicationTextFromDraft(t *testing.T) {
	draft := `🛠️ Подтвердите отправку заявки: 🛠️
Hi

👷 Отправил:
Denis

🏗️ Подразделение:
ЦНС`

	expected := "Hi"

	result := helpers.GetApplicationTextFromDraft(draft)

	if result != expected {
		t.Errorf("GetApplicationTextFromDraft(draft) = %s; want %s", result, expected)
	}
}
