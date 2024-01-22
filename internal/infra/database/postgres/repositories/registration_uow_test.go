package repositories

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres"
	testtools "github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres/test-tools"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/student"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/address"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/phone"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUowRegistration(t *testing.T) {

	std, err := student.New(
		"Pedrinho",
		"Souza",
		"2023-12-12",
		"123456789",
		"84731086043",
		"teste@test.com",
		true,
	)

	assert.NoError(t, err)

	add := []address.RequestDto{
		{
			Street:   "Rua dos Bobos",
			City:     "SSA",
			District: "SC",
			State:    "SP",
			ZipCode:  "41500030",
		},
	}

	std.AddAddress(add)

	p := []phone.RequestDto{
		{
			Description: "Pessoal",
			Phone:       "71589955554",
		},
	}

	std.AddPhones(p)
	testtools.StartTestEnv()
	db := postgres.Connect()
	studentRepository := *NewStudentRepository(db)
	registrationUow := NewRegistrationUow(db, studentRepository, *NewRegistrationRepository(db))

	_ = registrationUow.BeginTransaction()
	err = registrationUow.CreateStudent(*std)
	assert.NoError(t, err)

	_ = registrationUow.Rollback()

	studentDb, err := studentRepository.FindByCpf(std.Cpf())
	assert.Error(t, err)
	assert.Nil(t, studentDb)
}
