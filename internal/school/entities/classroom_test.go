package entities

import (
	"github.com/google/uuid"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var currentTime time.Time = time.Now()

func TestShouldChangeOccupiedVacancyWithSuccess(t *testing.T) {

	classRoom := ClassRoom{
		Id:              uuid.New(),
		VacancyQuantity: 20,
		Shift:           "Matutino",
		Status:          "OPEN",
		Level:           "2 Ano Fundamental",
		OpenDate:        &currentTime,
		Identification:  "TUR-A123",
	}

	classRoom.SetOccupiedVacancies(5)
	assert.Equal(t, 5, classRoom.GetOccupiedVacancies())
}

func TestShouldReturnErrorIfQuantityOfOccupiedVacancyIsGreaterThanVacancy(t *testing.T) {
	classRoom := ClassRoom{
		Id:              uuid.New(),
		VacancyQuantity: 20,
		Shift:           "Matutino",
		Status:          "OPEN",
		Level:           "2 Ano Fundamental",
		OpenDate:        &currentTime,
		Identification:  "TUR-A123",
	}

	err := classRoom.SetOccupiedVacancies(21)
	assert.Error(t, err)
}

func TestShouldReturnErrorIfChangeOccupiedVacancyInClassRoomClosed(t *testing.T) {
	classRoom := ClassRoom{
		Id:              uuid.New(),
		VacancyQuantity: 20,
		Shift:           "Matutino",
		Status:          "CLOSED",
		Level:           "2 Ano Fundamental",
		OpenDate:        &currentTime,
		Identification:  "TUR-A123",
	}

	err := classRoom.SetOccupiedVacancies(21)
	assert.Error(t, err)
}

func TestShouldReturnErrorIfQuantityRequestedIsGreaterThanVacancyQuantity(t *testing.T) {
	classRoom := ClassRoom{
		Id:              uuid.New(),
		VacancyQuantity: 20,
		Shift:           "Matutino",
		Status:          "CLOSED",
		Level:           "2 Ano Fundamental",
		OpenDate:        &currentTime,
		Identification:  "TUR-A123",
	}

	err := classRoom.SetOccupiedVacancies(50)
	assert.Error(t, err)
}
