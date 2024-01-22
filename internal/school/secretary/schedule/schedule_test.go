package schedule

import (
	"github.com/google/uuid"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCreateScheduleEntityWithSuccess(t *testing.T) {
	dto := Request{
		Description: "Any Description",
		InitialTime: "08:00",
		FinalTime:   "09:00",
	}

	sch, err := New(
		dto.Description,
		dto.InitialTime,
		dto.FinalTime,
		uuid.New().String(),
	)
	assert.NoError(t, err)
	assert.Equal(t, "08:00:00", sch.StartAt())
	assert.Equal(t, "09:00:00", sch.EndAt())
}
