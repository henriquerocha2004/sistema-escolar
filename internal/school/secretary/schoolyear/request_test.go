package schoolyear

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnErrorInValidateDto(t *testing.T) {
	dto := Request{
		Year:      "asdasdsada",
		StartedAt: "sadsadss",
		EndAt:     "sadsadsad",
	}

	err := dto.Validate()
	assert.Error(t, err)
}

func TestShouldValidateWithSuccessRequestDto(t *testing.T) {
	dto := Request{
		Year:      "2004",
		StartedAt: "2022-01-01",
		EndAt:     "2022-02-01",
	}

	err := dto.Validate()
	assert.NoError(t, err)
}
