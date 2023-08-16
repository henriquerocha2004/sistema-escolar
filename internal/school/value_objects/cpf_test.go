package value_objects

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnErrorIfValueProvidedIsEmpty(t *testing.T) {
	var cpf CPF = ""
	err := cpf.Validate()
	assert.Error(t, err)
	assert.Equal(t, "empty cpf provided", err.Error())
}

func TestShouldReturnErrorIfValueProvidedIsLessThan11Characters(t *testing.T) {
	var cpf CPF = "123.1548."
	err := cpf.Validate()
	assert.Error(t, err)
	assert.Equal(t, "cpf must be 11 characters", err.Error())
}

func TestShouldReturnErrorIfAllCharactersIsEquals(t *testing.T) {

	equalsCpfs := []string{
		"111.111.111-11",
		"222.222.222-22",
		"333.333.333-33",
		"444.444.444-44",
		"555.555.555-55",
		"666.666.666-66",
		"777.777.777-77",
		"888.888.888-88",
		"999.999.999-99",
		"000.000.000-00",
	}

	for _, cpfEqual := range equalsCpfs {
		var cpf CPF = CPF(cpfEqual)
		err := cpf.Validate()
		assert.Error(t, err)
		assert.Equal(t, "all digits from cpf are equals", err.Error())
	}
}
