package requestvalidator

import (
	"errors"
	"log"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

type DtoValidator interface {
	Validate() error
}

type ValidatorMessage struct {
	Param   string `json:"param"`
	Message string `json:"message"`
}

func ValidateRequest(requestDto DtoValidator) *[]ValidatorMessage {
	err := requestDto.Validate()
	if err == nil {
		return nil
	}

	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		out := make([]ValidatorMessage, len(validationErrors))
		for i, messageError := range validationErrors {
			out[i] = ValidatorMessage{
				Param:   messageError.Field(),
				Message: msgForTag(messageError),
			}
		}

		return &out
	}

	return nil
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "allowedColumn":
		return "invalid column search"
	case "dbOrder":
		return "invalid value to sort: only desc ou asc allowed"
	case "numeric":
		return "only number are allowed"
	case "date::format:yyyy-mm-dd":
		return "invalid date format"
	}
	return errors.New("undefined error").Error() // default error
}

// DbOrder Funcao que valida se o valor do campo é desc ou asc
func DbOrder(field validator.FieldLevel) bool {
	return field.Field().String() == "asc" || field.Field().String() == "desc"
}

// RoomAllowedColumnSearch Funcao que verifica as colunas permitidas a serem pesquisadas na tabela rooms
func RoomAllowedColumnSearch(field validator.FieldLevel) bool {
	allowedColumns := []string{"code", "description", "capacity"}

	for _, column := range allowedColumns {
		if column == field.Field().String() {
			return true
		}
	}

	return false
}

// ValidateHourFormat Funcao que valida o formato de hora Ex: 08:00
func ValidateHourFormat(field validator.FieldLevel) bool {
	regex, _ := regexp.Compile("^(0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]$")
	return regex.MatchString(field.Field().String())
}

// ValidateYearFormat Funcao que valida se a string é um ano valido YYYY
func ValidateYearFormat(field validator.FieldLevel) bool {
	regex, _ := regexp.Compile(`(19|20)\d{2}`)
	return regex.MatchString(field.Field().String())
}

// ValidateDateUSA Função que valida se a string está no formato de data valido yyyy-dd-dd
func ValidateDateUSA(field validator.FieldLevel) bool {
	_, err := time.Parse("2006-01-02", field.Field().String())
	log.Println(err)
	return err == nil
}

// ValidateShift Funcao que valida se o valor do campo shift da classe é permitido
func ValidateShift(field validator.FieldLevel) bool {
	allowedShifts := []string{"morning", "afternoon", "nocturnal", "full-time"}

	for _, shift := range allowedShifts {
		if shift == field.Field().String() {
			return true
		}
	}

	return false
}

// ValidateType Funcao que valida se o valor do campo type da classe é permitida
func ValidateType(field validator.FieldLevel) bool {
	allowedTypes := []string{"in_person", "remote"}

	for _, typeClass := range allowedTypes {
		if typeClass == field.Field().String() {
			return true
		}
	}

	return false
}

func ValidateClassStatus(field validator.FieldLevel) bool {
	allowedStatuses := []string{"open", "closed", "cancelled"}

	for _, status := range allowedStatuses {
		if status == field.Field().String() {
			return true
		}
	}

	return false
}
