package helpers

import (
	"context"
	"log"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"github.com/kuzmindeniss/prost/internal/db/repository"
	"github.com/kuzmindeniss/prost/internal/tg/initializers"
)

var UnitsList []repository.Unit

var unitsMutex sync.RWMutex

func LoadUnits() error {
	log.Println("Loading units from database")

	// Fetch units from database
	units, err := initializers.Repo.GetUnits(context.Background())
	if err != nil {
		log.Printf("Failed to load units: %v", err)
		return err
	}

	// Initialize maps
	tempUnitMapByName := make(map[string]repository.Unit)
	tempUnitMapByID := make(map[uuid.UUID]repository.Unit)

	// Populate maps
	for _, unit := range units {
		tempUnitMapByName[unit.Name] = unit
		tempUnitMapByID[unit.ID] = unit
	}

	// Update global maps with mutex protection
	unitsMutex.Lock()
	UnitsList = units
	unitsMutex.Unlock()

	log.Printf("Loaded %d units successfully", len(units))
	return nil
}

// GetAllUnits returns a copy of all units
func GetAllUnits() []repository.Unit {
	unitsMutex.RLock()
	defer unitsMutex.RUnlock()
	unitsCopy := make([]repository.Unit, len(UnitsList))
	copy(unitsCopy, UnitsList)
	return unitsCopy
}

func CheckIfMessageTextIsUnit(message *tgbotapi.Message) (bool, repository.Unit) {
	if message == nil {
		return false, repository.Unit{}
	}

	units := GetAllUnits()
	for _, unit := range units {
		if message.Text == unit.Name {
			return true, unit
		}
	}
	return false, repository.Unit{}
}
