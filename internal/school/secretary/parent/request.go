package parent

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/address"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/phone"
)

type RequestDto struct {
	FirstName   string               `json:"first_name"`
	LastName    string               `json:"last_name"`
	BirthDay    string               `json:"birth_day"`
	Addresses   []address.RequestDto `json:"addresses"`
	Phones      []phone.RequestDto   `json:"phones"`
	RgDocument  string               `json:"rg_document"`
	CpfDocument string               `json:"cpf_document"`
	Email       string               `json:"email"`
}
