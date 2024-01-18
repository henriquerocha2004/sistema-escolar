package classroom

import (
	"github.com/google/uuid"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldChangeOccupiedVacancyWithSuccess(t *testing.T) {
	classRoom, err := New(
		20,
		"morning",
		"OPEN",
		"TUR-A123",
		uuid.New().String(),
		"",
		uuid.New().String(),
		"Terreo",
		"in_person",
	)

	assert.NoError(t, err)
	err = classRoom.SetOccupiedVacancies(5)
	assert.NoError(t, err)
	assert.Equal(t, 5, classRoom.OccupiedVacancies())
}

func TestShouldReturnErrorIfQuantityOfOccupiedVacancyIsGreaterThanVacancy(t *testing.T) {
	classRoom, err := New(
		20,
		"morning",
		"OPEN",
		"TUR-A123",
		uuid.New().String(),
		"",
		uuid.New().String(),
		"Terreo",
		"in_person",
	)

	assert.NoError(t, err)

	err = classRoom.SetOccupiedVacancies(21)
	assert.Error(t, err)
}

func TestShouldReturnErrorIfChangeOccupiedVacancyInClassRoomClosed(t *testing.T) {
	classRoom, err := New(
		20,
		"morning",
		"OPEN",
		"TUR-A123",
		uuid.New().String(),
		"",
		uuid.New().String(),
		"Terreo",
		"in_person",
	)

	assert.NoError(t, err)

	err = classRoom.SetOccupiedVacancies(21)
	assert.Error(t, err)
}

func TestShouldReturnErrorIfQuantityRequestedIsGreaterThanVacancyQuantity(t *testing.T) {
	classRoom, err := New(
		20,
		"morning",
		"OPEN",
		"TUR-A123",
		uuid.New().String(),
		"",
		uuid.New().String(),
		"Terreo",
		"in_person",
	)

	assert.NoError(t, err)

	err = classRoom.SetOccupiedVacancies(50)
	assert.Error(t, err)
}
