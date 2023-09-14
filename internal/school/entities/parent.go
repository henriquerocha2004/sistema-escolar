package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/value_objects"
)

type Parent struct {
	Id          uuid.UUID               `json:"id"`
	FirstName   string                  `json:"first_name"`
	LastName    string                  `json:"last_name"`
	BirthDay    *time.Time              `json:"birth_day"`
	Addresses   []value_objects.Address `json:"addresses"`
	Phones      []value_objects.Phone   `json:"phones"`
	RgDocument  string                  `json:"rg_document"`
	CpfDocument value_objects.CPF       `json:"cpf_document"`
	StudentId   uuid.UUID               `json:"student_id"`
	Email       string                  `json:"email"`
}

func (s *Parent) AddAddress(addressDto []dto.AddressDto) {

	var addresses []value_objects.Address

	for _, address := range addressDto {

		a := value_objects.Address{
			Id:       uuid.New(),
			Street:   address.Street,
			City:     address.City,
			District: address.District,
			State:    address.State,
			ZipCode:  address.ZipCode,
			OwnerId:  s.Id,
		}

		addresses = append(addresses, a)

	}

	s.Addresses = addresses
}

func (s *Parent) AddPhones(phonesDto []dto.PhoneDto) {

	var phones []value_objects.Phone

	for _, phone := range phonesDto {
		p := value_objects.Phone{
			Id:          uuid.New(),
			Description: phone.Description,
			Phone:       phone.Description,
			OwnerId:     s.Id,
		}

		phones = append(phones, p)
	}

	s.Phones = phones
}
