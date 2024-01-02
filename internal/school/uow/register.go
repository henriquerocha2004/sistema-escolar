package uow

import (
	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities"
)

type RegisterUow interface {
	BeginTransaction() error
	Commit() error
	Rollback() error
	CreateStudent(student entities.Student) error
	CreateRegister(register entities.Registration) error
	StudentAlreadyExists(cpf string) (*uuid.UUID, error)
	StudentAlreadyRegisterInClass(studentId uuid.UUID, classRoomId uuid.UUID) (bool, error)
}
