package schedule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCreateScheduleEntityWithSuccess(t *testing.T) {
	dto := ScheduleRequestDto{
		Description: "Any Description",
		InitialTime: "08:00",
		FinalTime:   "09:00",
	}

	schedule := ScheduleClass{}
	err := schedule.FillFromDto(dto)
	assert.NoError(t, err)
	assert.Equal(t, "08:00-09:00", schedule.Schedule)
}
