package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnErrorInValidateDto(t *testing.T) {
	dto := SchoolYearRequestDto{
		Year:      "asdasdsada",
		StartedAt: "sadsadss",
		EndAt:     "sadsadsad",
	}

	err := dto.Validate()
	assert.Error(t, err)
}

func TestShouldValidateWithSucessRequestDto(t *testing.T) {
	dto := SchoolYearRequestDto{
		Year:      "2004",
		StartedAt: "2022-01-01",
		EndAt:     "2022-02-01",
	}

	err := dto.Validate()
	assert.NoError(t, err)
}
