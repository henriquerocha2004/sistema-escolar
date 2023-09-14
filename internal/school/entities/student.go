package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
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

func (s *Student) FillFromDto(dto dto.StudentDto) {
	s.Id = uuid.New()
	s.FirstName = dto.FirstName
	s.LastName = dto.LastName
	birthDay, _ := time.Parse("2006-02-02", dto.Birthday)
	s.BirthDay = &birthDay
	s.RgDocument = dto.RgDocument
	s.CPFDocument = value_objects.CPF(dto.CpfDocument)
	s.Email = dto.Email
	s.HimSelfResponsible = dto.HimSelfResponsible
}

func (s *Student) AddAddress(addressDto []dto.AddressDto) {

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

func (s *Student) AddPhones(phonesDto []dto.PhoneDto) {

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

func (s *Student) AddParent(parentsDto []dto.ParentDto) {

	var parents []Parent

	for _, parent := range parentsDto {

		birthDay, _ := time.Parse("2006-02-02", parent.BirthDay)

		p := Parent{
			Id:          uuid.New(),
			FirstName:   parent.FirstName,
			LastName:    parent.LastName,
			BirthDay:    &birthDay,
			RgDocument:  parent.RgDocument,
			CpfDocument: value_objects.CPF(parent.CpfDocument),
			StudentId:   s.Id,
			Email:       parent.Email,
		}

		p.AddAddress(parent.Addresses)
		p.AddPhones(parent.Phones)

		parents = append(parents, p)
	}

	s.Parents = parents
}

func (s *Student) Validate() error {
	err := s.ValidateResponsible()
	if err != nil {
		return err
	}

	err = s.CPFDocument.Validate()
	if err != nil {
		return err
	}

	err = s.ValidateContactInformation()
	if err != nil {
		return err
	}

	err = s.ValidateParentsInformation()
	if err != nil {
		return err
	}

	return nil
}

func (s *Student) ValidateResponsible() error {

	if s.HimSelfResponsible {
		return nil
	}

	if len(s.Parents) < 1 {
		return errors.New("information about student parents not found")
	}

	return nil
}

func (s *Student) ValidateContactInformation() error {
	if !s.HimSelfResponsible {
		return nil
	}

	if len(s.Addresses) < 1 {
		return errors.New("address information not found")
	}

	if len(s.Phones) < 1 {
		return errors.New("phone information not found")
	}

	return nil
}

func (s *Student) ValidateParentsInformation() error {
	if s.HimSelfResponsible {
		return nil
	}

	for _, parent := range s.Parents {
		err := parent.CpfDocument.Validate()
		if err != nil {
			return err
		}

		if len(parent.Addresses) < 1 {
			return errors.New("parent address information not found")
		}

		if len(parent.Phones) < 1 {
			return errors.New("parent phone not found")
		}
	}

	return nil
}
