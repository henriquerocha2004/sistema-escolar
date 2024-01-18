package registration

import (
	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/student"
)

type RegisterUow interface {
	BeginTransaction() error
	Commit() error
	Rollback() error
	CreateStudent(student student.Student) error
	CreateRegister(register Registration) error
	StudentAlreadyExists(cpf string) (*uuid.UUID, error)
	StudentAlreadyRegisterInClass(studentId uuid.UUID, classRoomId uuid.UUID) (bool, error)
}
