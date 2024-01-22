package student

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/parent"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/address"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/phone"
)

type RequestDto struct {
	FirstName          string               `json:"first_name"`
	LastName           string               `json:"last_name"`
	Birthday           string               `json:"birthday"`
	RgDocument         string               `json:"rg_document"`
	CpfDocument        string               `json:"cpf_document"`
	Email              string               `json:"email"`
	HimSelfResponsible bool                 `json:"him_self_responsible"`
	Addresses          []address.RequestDto `json:"addresses"`
	Phones             []phone.RequestDto   `json:"phones"`
	Parents            []parent.RequestDto  `json:"parents"`
}
