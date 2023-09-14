package postgres

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/value_objects"
)

type RegistrationRepository struct {
	db     *sql.DB
	queues *Queries
}

func NewRegistrationRepository(db *sql.DB) *RegistrationRepository {
	return &RegistrationRepository{
		db:     db,
		queues: New(db),
	}
}

func (r *RegistrationRepository) Create(registration entities.Registration) error {

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	r.queues = r.queues.WithTx(tx)

	classRoomModel, err := r.queues.FindClassByIdLock(context.Background(), registration.Class.Id)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if classRoomModel.VacanciesOccupied >= classRoomModel.Vacancies {
		_ = tx.Rollback()
		return errors.New("no vacancies available for this class")
	}

	err = r.createStudent(registration.Student)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	monthlyFee := strconv.FormatFloat(registration.MonthlyFee, 'f', -1, 64)
	enrollmentFee := strconv.FormatFloat(registration.EnrollmentFee, 'f', -1, 64)

	registrationModel := CreateRegistrationParams{
		ID:   registration.Id,
		Code: registration.Code,
		ClassRoomID: uuid.NullUUID{
			UUID:  classRoomModel.ID,
			Valid: true,
		},
		Shift: sql.NullString{
			String: string(registration.Shift),
			Valid:  true,
		},
		StudentID:            registration.Student.Id,
		ServiceID:            registration.Service.Id,
		MonthlyFee:           monthlyFee,
		InstallmentsQuantity: int32(registration.InstallmentsQuantity),
		EnrollmentFee: sql.NullString{
			String: enrollmentFee,
			Valid:  true,
		},
		DueDate: sql.NullTime{
			Time:  *registration.EnrollmentDueDate,
			Valid: true,
		},
		MonthDuration: sql.NullInt32{
			Int32: int32(registration.MonthDuration),
			Valid: true,
		},
		Status: registration.Status,
		EnrollmentDate: sql.NullTime{
			Time:  *registration.EnrollmentDate,
			Valid: true,
		},
		SchoolYearID: uuid.NullUUID{
			UUID:  registration.Class.SchoolYearId,
			Valid: true,
		},
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	err = r.queues.CreateRegistration(context.Background(), registrationModel)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	rowAffected, err := r.UpdateOccupiedVacancies(classRoomModel.ID.String(), classRoomModel.VacanciesOccupied)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if rowAffected < 1 {
		_ = tx.Rollback()
		return errors.New("error at update vacancies occupied")
	}

	_ = tx.Commit()

	return nil
}

func (r *RegistrationRepository) createStudent(student entities.Student) error {

	studentModel := CreateStudentParams{
		ID:        student.Id,
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Birthday:  *student.BirthDay,
		RgDocument: sql.NullString{
			String: student.RgDocument,
			Valid:  true,
		},
		CpfDocument: string(student.CPFDocument),
		Email: sql.NullString{
			String: student.Email,
			Valid:  true,
		},
		HimSelfResponsible: student.HimSelfResponsible,
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	err := r.queues.CreateStudent(context.Background(), studentModel)
	if err != nil {
		return err
	}

	err = r.createAddress(student.Addresses)
	if err != nil {
		return err
	}

	err = r.createPhone(student.Phones)
	if err != nil {
		return err
	}

	err = r.createParents(student.Parents)
	if err != nil {
		return err
	}

	return nil
}

func (r *RegistrationRepository) createParents(parents []entities.Parent) error {
	for _, parent := range parents {
		parentModel := CreateParentParams{
			ID:        parent.Id,
			FirstName: parent.FirstName,
			LastName:  parent.LastName,
			Birthday:  *parent.BirthDay,
			RgDocument: sql.NullString{
				String: parent.RgDocument,
				Valid:  true,
			},
			CpfDocument: string(parent.CpfDocument),
			StudentID:   parent.StudentId,
			Email:       parent.Email,
			CreatedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
			UpdatedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
		}

		err := r.queues.CreateParent(context.Background(), parentModel)
		if err != nil {
			return err
		}

		err = r.createAddress(parent.Addresses)
		if err != nil {
			return err
		}

		err = r.createPhone(parent.Phones)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *RegistrationRepository) createAddress(addresses []value_objects.Address) error {

	for _, address := range addresses {
		addressModel := CreateAddressParams{
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

		err := r.queues.CreateAddress(context.Background(), addressModel)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *RegistrationRepository) createPhone(phones []value_objects.Phone) error {
	for _, phone := range phones {
		phoneModel := CreatePhoneParams{
			ID:          phone.Id,
			Description: phone.Description,
			Phone:       phone.Phone,
			OwnerID:     phone.OwnerId,
		}

		err := r.queues.CreatePhone(context.Background(), phoneModel)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *RegistrationRepository) UpdateOccupiedVacancies(classroomId string, currentVacancyOccupied int32) (int, error) {

	query := `UPDATE class_room 
    			SET vacancies_occupied = $1, 
        			updated_at = $2 
				WHERE 
					id = $3 
					AND vacancies_occupied < vacancies 
					AND deleted_at IS NULL;`
	vacanciesOccupied := currentVacancyOccupied + 1

	resp, err := r.queues.db.ExecContext(
		context.Background(),
		query,
		vacanciesOccupied,
		time.Now().Format("2006-02-02 15:04:05"),
		classroomId,
	)

	if err != nil {
		return 0, err
	}

	rowsAffected, _ := resp.RowsAffected()

	return int(rowsAffected), err
}
