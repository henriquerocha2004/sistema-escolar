package student

import (
	"encoding/json"
	"errors"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/parent"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/registration"
	"log"

	"time"

	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/value_objects"
)

type Student struct {
	id                 uuid.UUID
	firstName          string
	lastName           string
	birthDay           *time.Time
	rg                 string
	cpf                value_objects.CPF
	email              string
	himSelfResponsible bool
	addresses          []value_objects.Address
	phones             []value_objects.Phone
	parents            []parent.Parent
}

func New(firstName string, lastName string, birthDay string, rg string, cpf string, email string, himselfResponsible bool) (*Student, error) {
	s := &Student{
		id: uuid.New(),
	}

	err := s.ChangeName(firstName, lastName)
	if err != nil {
		return nil, err
	}

	err = s.ChangeBirthDay(birthDay)
	if err != nil {
		return nil, err
	}

	s.ChangeRg(rg)
	s.ChangeHimSelfResponsible(himselfResponsible)

	err = s.ChangeCPF(cpf)
	if err != nil {
		return nil, err
	}

	err = s.ChangeEmail(email)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func Load(id string, firstName string, lastName string, birthDay string, rg string, cpf string, email string, himselfResponsible bool) (*Student, error) {
	sdt, err := New(firstName, lastName, birthDay, rg, cpf, email, himselfResponsible)
	if err != nil {
		return nil, err
	}

	err = sdt.ChangeId(id)
	if err != nil {
		return nil, err
	}

	return sdt, nil
}

func (s *Student) ChangeId(id string) error {
	if id == "" {
		return errors.New("student id cannot be empty")
	}

	parentId, err := uuid.Parse(id)
	if err != nil {
		log.Println(err)
		return errors.New("failed to change student id")
	}

	s.id = parentId

	return nil
}

func (s *Student) Parents() []parent.Parent {
	return s.parents
}

func (s *Student) Addresses() []value_objects.Address {
	return s.addresses
}

func (s *Student) Phones() []value_objects.Phone {
	return s.phones
}

func (s *Student) Email() string {
	return s.email
}

func (s *Student) Id() uuid.UUID {
	return s.id
}

func (s *Student) FirstName() string {
	return s.firstName
}

func (s *Student) LastName() string {
	return s.lastName
}

func (s *Student) BirthDay() *time.Time {
	return s.birthDay
}

func (s *Student) Rg() string {
	return s.rg
}

func (s *Student) Cpf() value_objects.CPF {
	return s.cpf
}

func (s *Student) HimSelfResponsible() bool { return s.himSelfResponsible }

func (s *Student) ChangeHimSelfResponsible(responsible bool) {
	s.himSelfResponsible = responsible
}

func (s *Student) ChangeEmail(email string) error {
	if email == "" {
		return errors.New("student email cannot be null")
	}

	s.email = email

	return nil
}

func (s *Student) ChangeCPF(cpf string) error {
	docCpf := value_objects.CPF(cpf)

	err := docCpf.Validate()
	if err != nil {
		return err
	}

	s.cpf = docCpf

	return nil
}

func (s *Student) ChangeRg(rg string) {
	if rg == "" {
		return
	}

	s.rg = rg
}

func (s *Student) ChangeBirthDay(birthDay string) error {
	if birthDay == "" {
		return errors.New("student birthday cannot be empty")
	}

	b, err := time.Parse("2006-01-02", birthDay)
	if err != nil {
		log.Println(err)
		return errors.New("failed to change parent birthday")
	}

	s.birthDay = &b

	return nil
}

func (s *Student) ChangeName(firstName string, lastName string) error {
	if firstName == "" || lastName == "" {
		return errors.New("student first name and last name cannot be null")
	}

	s.firstName = firstName
	s.lastName = lastName

	return nil
}

func (s *Student) AddAddress(addressDto []registration.AddressDto) {

	var addresses []value_objects.Address

	for _, address := range addressDto {

		a := value_objects.Address{
			Id:       uuid.New(),
			Street:   address.Street,
			City:     address.City,
			District: address.District,
			State:    address.State,
			ZipCode:  address.ZipCode,
			OwnerId:  s.Id(),
		}

		addresses = append(addresses, a)

	}

	s.addresses = addresses
}

func (s *Student) AddPhones(phonesDto []registration.PhoneDto) {

	var phones []value_objects.Phone

	for _, phone := range phonesDto {
		p := value_objects.Phone{
			Id:          uuid.New(),
			Description: phone.Description,
			Phone:       phone.Description,
			OwnerId:     s.Id(),
		}

		phones = append(phones, p)
	}

	s.phones = phones
}

func (s *Student) AddParents(parentsDto []registration.ParentDto) error {

	var parents []parent.Parent

	for _, parentDto := range parentsDto {

		p, err := parent.New(parentDto.FirstName,
			parentDto.LastName,
			parentDto.BirthDay,
			parentDto.RgDocument,
			parentDto.CpfDocument,
			s.Id().String(),
			parentDto.Email)

		if err != nil {
			return err
		}

		p.AddAddress(parentDto.Addresses)
		p.AddPhones(parentDto.Phones)

		parents = append(parents, *p)
	}

	s.parents = parents

	return nil
}

func (s *Student) Validate() error {
	err := s.ValidateResponsible()
	if err != nil {
		return err
	}

	err = s.cpf.Validate()
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

	if s.himSelfResponsible {
		return nil
	}

	if len(s.parents) < 1 {
		return errors.New("information about student parents not found")
	}

	return nil
}

func (s *Student) ValidateContactInformation() error {
	if !s.himSelfResponsible {
		return nil
	}

	if len(s.Addresses()) < 1 {
		return errors.New("address information not found")
	}

	if len(s.Phones()) < 1 {
		return errors.New("phone information not found")
	}

	return nil
}

func (s *Student) ValidateParentsInformation() error {
	if s.himSelfResponsible {
		return nil
	}

	for _, studentParent := range s.Parents() {
		if len(studentParent.Addresses()) < 1 {
			return errors.New("parent address information not found")
		}

		if len(studentParent.Phones()) < 1 {
			return errors.New("parent phone not found")
		}
	}

	return nil
}

func (s *Student) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id        string                  `json:"id"`
		FirstName string                  `json:"first_name"`
		LastName  string                  `json:"last_name"`
		BirthDay  string                  `json:"birth_day"`
		Addresses []value_objects.Address `json:"addresses"`
		Phones    []value_objects.Phone   `json:"phones"`
		Rg        string                  `json:"rg"`
		Cpf       string                  `json:"cpf"`
		StudentId string                  `json:"student_id"`
		Email     string                  `json:"email"`
		Parents   []parent.Parent         `json:"parents"`
	}{
		Id:        s.Id().String(),
		FirstName: s.FirstName(),
		LastName:  s.LastName(),
		BirthDay:  s.BirthDay().Format("2006-01-02"),
		Addresses: s.Addresses(),
		Phones:    s.Phones(),
		Rg:        s.Rg(),
		Cpf:       string(s.Cpf()),
		Email:     s.Email(),
		Parents:   s.Parents(),
	})
}
