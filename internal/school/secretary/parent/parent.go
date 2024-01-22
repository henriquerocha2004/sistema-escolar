package parent

import (
	"encoding/json"
	"errors"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/address"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/phone"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/value_objects"
)

type Parent struct {
	id          uuid.UUID
	firstName   string
	lastName    string
	birthDay    *time.Time
	addresses   []value_objects.Address
	phones      []value_objects.Phone
	rgDocument  string
	cpfDocument value_objects.CPF
	studentId   uuid.UUID
	email       string
}

func New(firstName string, lastName string, birthDay string, rg string, cpf string, studentId string, email string) (*Parent, error) {
	p := &Parent{
		id: uuid.New(),
	}

	err := p.ChangeName(firstName, lastName)
	if err != nil {
		return nil, err
	}

	err = p.ChangeBirthDay(birthDay)
	if err != nil {
		return nil, err
	}

	p.ChangeRg(rg)

	err = p.ChangeCPF(cpf)
	if err != nil {
		return nil, err
	}

	err = p.ChangeStudentId(studentId)
	if err != nil {
		return nil, err
	}

	err = p.ChangeEmail(email)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Parent) FirstName() string {
	return p.firstName
}

func (p *Parent) LastName() string {
	return p.lastName
}

func (p *Parent) BirthDay() *time.Time {
	return p.birthDay
}

func (p *Parent) Rg() string {
	return p.rgDocument
}

func (p *Parent) Cpf() value_objects.CPF {
	return p.cpfDocument
}

func (p *Parent) StudentId() uuid.UUID {
	return p.studentId
}

func (p *Parent) Email() string {
	return p.email
}

func (p *Parent) Id() uuid.UUID {
	return p.id
}

func (p *Parent) ChangeEmail(email string) error {
	if email == "" {
		return errors.New("parent email cannot be null")
	}

	p.email = email

	return nil
}

func (p *Parent) ChangeStudentId(studentId string) error {
	if studentId == "" {
		return errors.New("student id cannot be empty")
	}

	student, err := uuid.Parse(studentId)
	if err != nil {
		log.Println(err)
		return errors.New("failed to change student id in parent")
	}

	p.studentId = student

	return nil
}

func (p *Parent) ChangeCPF(cpf string) error {
	docCpf := value_objects.CPF(cpf)

	err := docCpf.Validate()
	if err != nil {
		return err
	}

	p.cpfDocument = docCpf

	return nil
}

func (p *Parent) ChangeRg(rg string) {
	if rg == "" {
		return
	}

	p.rgDocument = rg
}

func (p *Parent) ChangeName(firstName string, lastName string) error {
	if firstName == "" || lastName == "" {
		return errors.New("parent first name and last name cannot be null")
	}

	p.firstName = firstName
	p.lastName = lastName

	return nil
}

func (p *Parent) ChangeBirthDay(birthDay string) error {
	if birthDay == "" {
		return errors.New("parent birthday cannot be empty")
	}

	b, err := time.Parse("2006-01-02", birthDay)
	if err != nil {
		log.Println(err)
		return errors.New("failed to change parent birthday")
	}

	p.birthDay = &b

	return nil
}

func (p *Parent) ChangeId(id string) error {
	if id == "" {
		return errors.New("parent id cannot be empty")
	}

	parentId, err := uuid.Parse(id)
	if err != nil {
		log.Println(err)
		return errors.New("failed to change parent id")
	}

	p.id = parentId

	return nil
}

func (p *Parent) Addresses() []value_objects.Address {
	return p.addresses
}

func (p *Parent) AddAddress(addressDto []address.RequestDto) {

	var addresses []value_objects.Address

	for _, address := range addressDto {

		a := value_objects.Address{
			Id:       uuid.New(),
			Street:   address.Street,
			City:     address.City,
			District: address.District,
			State:    address.State,
			ZipCode:  address.ZipCode,
			OwnerId:  p.Id(),
		}

		addresses = append(addresses, a)

	}

	p.addresses = addresses
}

func (p *Parent) Phones() []value_objects.Phone {
	return p.phones
}

func (p *Parent) AddPhones(phonesDto []phone.RequestDto) {

	var phones []value_objects.Phone

	for _, phone := range phonesDto {
		p := value_objects.Phone{
			Id:          uuid.New(),
			Description: phone.Description,
			Phone:       phone.Description,
			OwnerId:     p.Id(),
		}

		phones = append(phones, p)
	}

	p.phones = phones
}

func (p *Parent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id          string                  `json:"id"`
		FirstName   string                  `json:"first_name"`
		LastName    string                  `json:"last_name"`
		BirthDay    string                  `json:"birth_day"`
		Addresses   []value_objects.Address `json:"addresses"`
		Phones      []value_objects.Phone   `json:"phones"`
		RgDocument  string                  `json:"rg_document"`
		CpfDocument value_objects.CPF       `json:"cpf_document"`
		StudentId   string                  `json:"student_id"`
		Email       string                  `json:"email"`
	}{
		Id:          p.Id().String(),
		FirstName:   p.FirstName(),
		LastName:    p.LastName(),
		BirthDay:    p.BirthDay().Format("2006-01-02"),
		Addresses:   p.Addresses(),
		Phones:      p.Phones(),
		RgDocument:  p.Rg(),
		CpfDocument: p.Cpf(),
		StudentId:   p.StudentId().String(),
		Email:       p.Email(),
	})
}
