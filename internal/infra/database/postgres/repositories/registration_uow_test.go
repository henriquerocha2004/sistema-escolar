package repositories

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities/student"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/registration"
	"log"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/value_objects"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	rootProject, _ := os.Getwd()
	err := godotenv.Load(rootProject + "/../../../../.env.test")
	if err != nil {
		log.Fatal("Error in read .env file")
	}
}

func TestUowRegistration(t *testing.T) {

	student := student.Student{
		Id:                 uuid.New(),
		FirstName:          "Pedrinho",
		LastName:           "Souza",
		BirthDay:           &time.Time{},
		RgDocument:         "123456789",
		CPFDocument:        value_objects.CPF("17515874698"),
		Email:              "teste@test.com",
		HimSelfResponsible: true,
	}

	address := []registration.AddressDto{
		{
			Street:   "Rua dos Bobos",
			City:     "SSA",
			District: "SC",
			State:    "SP",
			ZipCode:  "41500030",
		},
	}

	student.AddAddress(address)

	phone := []registration.PhoneDto{
		{
			Description: "Pessoal",
			Phone:       "71589955554",
		},
	}

	student.AddPhones(phone)
	db := postgres.Connect()
	studentRepository := *NewStudentRepository(db)
	registrationUow := NewRegistrationUow(db, studentRepository, *NewRegistrationRepository(db))

	registrationUow.BeginTransaction()
	err := registrationUow.CreateStudent(student)
	assert.NoError(t, err)

	_ = registrationUow.Rollback()

	studentDb, err := studentRepository.FindByCpf(student.CPFDocument)
	assert.Error(t, err)
	assert.Nil(t, studentDb)
}
