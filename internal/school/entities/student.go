package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/value_objects"
)

type Student struct {
	Id                 uuid.UUID               `json:"id"`
	FirstName          string                  `json:"first_name"`
	LastName           string                  `json:"last_name"`
	BirthDay           *time.Time              `json:"birth_day"`
	RgDocument         string                  `json:"rg_document"`
	CPFDocument        value_objects.CPF       `json:"cpf_document"`
	Email              string                  `json:"email"`
	HimSelfResponsible bool                    `json:"him_self_responsible"`
	Addresses          []value_objects.Address `json:"addresses"`
	Phones             []value_objects.Phone   `json:"phones"`
	Parents            []Parent                `json:"parents"`
}
