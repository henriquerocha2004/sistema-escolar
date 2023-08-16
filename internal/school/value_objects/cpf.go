package value_objects

import (
	"errors"
	"regexp"
)

type CPF string

func (c *CPF) Validate() error {
	digits := 11
	factorDigit1 := 10
	factorDigit2 := 11

	if *c == "" {
		return errors.New("empty cpf provided")
	}
	c.clean()

	if len(*c) < digits {
		return errors.New("cpf must be 11 characters")
	}

	if c.hasAllDigitsEqual() {
		return errors.New("all digits from cpf are equals")
	}

	digit1 := c.calculateCheckDigit(factorDigit1)
	digit2 := c.calculateCheckDigit(factorDigit2)
	cpfDigit := c.extractDigit()

	return nil
}

func (c *CPF) clean() {
	originalValue := *c
	regex := regexp.MustCompile(`[.\-/]`)
	cleanValue := regex.ReplaceAllString(string(originalValue), "")
	*c = CPF(cleanValue)
}

func (c *CPF) hasAllDigitsEqual() bool {
	cpf := []rune(*c)
	allEquals := true
	firstDigit := cpf[0]

	for _, digit := range cpf {
		if digit != firstDigit {
			allEquals = false
		}

		firstDigit = digit
	}

	return allEquals
}

func (c *CPF) calculateCheckDigit(factor int) int {
	total := 0
	digits := *c

	for _, digit := range digits {
		if factor > 1 {
			digitInt := int(digit)
			total += digitInt * factor
		}
	}

	rest := total % 11

	if rest < 2 {
		return 0
	}

	return (11 - rest)
}

func (c *CPF) extractDigit() string {
	cpf := *c

	return string(cpf[len(cpf)-2:])
}
