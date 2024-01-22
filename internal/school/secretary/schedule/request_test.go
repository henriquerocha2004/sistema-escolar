package schedule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnErrorIfTimeFormat24HoursIsInvalid(t *testing.T) {
	scheduleRequestDto := Request{
		InitialTime: "42:00",
		FinalTime:   "25:00",
		Description: "Horario",
	}

	err := scheduleRequestDto.Validate()
	fmt.Println(err)
	assert.Error(t, err)
}

func TestShouldValidateTimeFormatWithSuccess(t *testing.T) {
	scheduleRequestDto := Request{
		InitialTime: "08:00",
		FinalTime:   "09:00",
		Description: "Manha",
		SchoolYear:  "2001",
	}

	err := scheduleRequestDto.Validate()
	assert.NoError(t, err)
}
