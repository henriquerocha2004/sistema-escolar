package repositories

import (
	"database/sql"
	"errors"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/registration"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/student"

	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/value_objects"
)

type RegistrationUow struct {
	db               *sql.DB
	tx               *sql.Tx
	studentRepo      StudentRepository
	registrationRepo RegistrationRepository
}

func NewRegistrationUow(db *sql.DB, studentRepo StudentRepository, registerRepo RegistrationRepository) *RegistrationUow {
	return &RegistrationUow{
		db:               db,
		studentRepo:      studentRepo,
		registrationRepo: registerRepo,
	}
}

func (r *RegistrationUow) BeginTransaction() error {
	tx, err := r.db.Begin()

	if err != nil {
		return err
	}

	r.tx = tx

	return nil
}

func (r *RegistrationUow) Rollback() error {
	if r.tx == nil {
		return errors.New("failed to rollback transaction. Transaction not started")
	}

	err := r.tx.Rollback()
	if err != nil {
		return err
	}

	return nil
}

func (r *RegistrationUow) Commit() error {
	if r.tx == nil {
		return errors.New("failed in commit transaction. Transaction not started")
	}

	err := r.tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *RegistrationUow) CreateStudent(student student.Student) error {

	if r.tx == nil {
		return errors.New("failed in create student. Transaction not started")
	}

	r.studentRepo.SetTransaction(r.tx)

	return r.studentRepo.Create(student)
}

func (r *RegistrationUow) CreateRegister(register registration.Registration) error {
	if r.tx == nil {
		return errors.New("failed in register student. Transaction not started")
	}

	r.registrationRepo.SetTransaction(r.tx)

	return r.registrationRepo.Create(register)
}

func (r *RegistrationUow) StudentAlreadyExists(cpf string) (*uuid.UUID, error) {
	student, err := r.studentRepo.FindByCpf(value_objects.CPF(cpf))
	if err != nil {
		return nil, err
	}

	id := student.Id()

	return &id, nil
}

func (r *RegistrationUow) StudentAlreadyRegisterInClass(studentId uuid.UUID, classRoomId uuid.UUID) (bool, error) {

	registrationCode, err := r.registrationRepo.SearchStudentAlreadyRegistered(studentId, classRoomId)
	if err != nil {
		return false, err
	}

	if registrationCode != "" {
		return true, nil
	}

	return false, nil
}
