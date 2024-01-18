package repositories

import (
	"context"
	"database/sql"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/parent"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/student"
	"time"

	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres/models"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/value_objects"
)

type StudentRepository struct {
	queues *models.Queries
	db     *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{
		queues: models.New(db),
		db:     db,
	}
}

func (s *StudentRepository) SetTransaction(tx *sql.Tx) {
	s.queues = s.queues.WithTx(tx)
}

func (s *StudentRepository) Create(student student.Student) error {

	studentModel := models.CreateStudentParams{
		ID:        student.Id(),
		FirstName: student.FirstName(),
		LastName:  student.LastName(),
		Birthday:  *student.BirthDay(),
		RgDocument: sql.NullString{
			String: student.Rg(),
			Valid:  true,
		},
		CpfDocument: string(student.Cpf()),
		Email: sql.NullString{
			String: student.Email(),
			Valid:  true,
		},
		HimSelfResponsible: student.HimSelfResponsible(),
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	err := s.queues.CreateStudent(context.Background(), studentModel)
	if err != nil {
		return err
	}

	err = s.syncAddress(student.Id(), student.Addresses())
	if err != nil {
		return err
	}

	err = s.syncPhones(student.Id(), student.Phones())
	if err != nil {
		return err
	}

	err = s.syncParents(student.Id(), student.Parents())
	if err != nil {
		return err
	}

	return nil
}

func (s *StudentRepository) syncAddress(ownerId uuid.UUID, addresses []value_objects.Address) error {

	deleteAddressParams := models.DeleteAddressByOwnerParams{
		OwnerID: ownerId,
		DeletedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	err := s.queues.DeleteAddressByOwner(context.Background(), deleteAddressParams)
	if err != nil {
		return err
	}

	for _, address := range addresses {
		addressModel := models.CreateAddressParams{
			ID:       address.Id,
			Street:   address.Street,
			City:     address.City,
			District: address.District,
			State:    address.State,
			ZipCode:  address.ZipCode,
			OwnerID:  address.OwnerId,
			CreatedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
			UpdatedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
		}

		err = s.queues.CreateAddress(context.Background(), addressModel)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *StudentRepository) syncPhones(ownerId uuid.UUID, phones []value_objects.Phone) error {

	deletePhonesParams := models.DeletePhonesByOwnerParams{
		OwnerID: ownerId,
		DeletedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	err := s.queues.DeletePhonesByOwner(context.Background(), deletePhonesParams)
	if err != nil {
		return err
	}

	for _, phone := range phones {
		phoneModel := models.CreatePhoneParams{
			ID:          phone.Id,
			Description: phone.Description,
			Phone:       phone.Phone,
			OwnerID:     phone.OwnerId,
		}

		err := s.queues.CreatePhone(context.Background(), phoneModel)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *StudentRepository) syncParents(studentId uuid.UUID, parents []parent.Parent) error {

	deleteParentsParam := models.DeleteParentsByStudentParams{
		StudentID: studentId,
		DeletedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	err := s.queues.DeleteParentsByStudent(context.Background(), deleteParentsParam)
	if err != nil {
		return err
	}

	for _, parent := range parents {
		parentModel := models.CreateParentParams{
			ID:        parent.Id(),
			FirstName: parent.FirstName(),
			LastName:  parent.LastName(),
			Birthday:  *parent.BirthDay(),
			RgDocument: sql.NullString{
				String: parent.Rg(),
				Valid:  true,
			},
			CpfDocument: string(parent.Cpf()),
			StudentID:   parent.StudentId(),
			Email:       parent.Email(),
			CreatedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
			UpdatedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
		}

		err := s.queues.CreateParent(context.Background(), parentModel)
		if err != nil {
			return err
		}

		err = s.syncAddress(parent.Id(), parent.Addresses())
		if err != nil {
			return err
		}

		err = s.syncPhones(parent.Id(), parent.Phones())
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *StudentRepository) FindByCpf(cpf value_objects.CPF) (*student.Student, error) {

	studentModel, err := s.queues.FindByCPFDocument(context.Background(), string(cpf))

	if err != nil {
		return nil, err
	}

	sdt, err := student.Load(
		studentModel.ID.String(),
		studentModel.FirstName,
		studentModel.LastName,
		studentModel.Birthday.Format("2006-01-02"),
		studentModel.RgDocument.String,
		studentModel.CpfDocument,
		studentModel.Email.String,
		studentModel.HimSelfResponsible,
	)

	return sdt, err
}
