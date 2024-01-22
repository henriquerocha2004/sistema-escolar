package repositories

import (
	"context"
	"database/sql"
	"errors"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/registration"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres/models"
)

type RegistrationRepository struct {
	db     *sql.DB
	queues *models.Queries
}

func NewRegistrationRepository(db *sql.DB) *RegistrationRepository {
	return &RegistrationRepository{
		db:     db,
		queues: models.New(db),
	}
}

func (r *RegistrationRepository) SetTransaction(tx *sql.Tx) {
	r.queues = r.queues.WithTx(tx)
}

func (r *RegistrationRepository) Create(registration registration.Registration) error {
	classRoomModel, err := r.queues.FindClassByIdLock(context.Background(), registration.Class().Id())
	if err != nil {
		return err
	}

	if classRoomModel.VacanciesOccupied >= classRoomModel.Vacancies {
		return errors.New("no vacancies available for this class")
	}

	monthlyFee := strconv.FormatFloat(registration.MonthlyFee(), 'f', -1, 64)
	enrollmentFee := strconv.FormatFloat(registration.EnrollmentFee(), 'f', -1, 64)

	registrationModel := models.CreateRegistrationParams{
		ID:   registration.Id(),
		Code: registration.Code(),
		ClassRoomID: uuid.NullUUID{
			UUID:  classRoomModel.ID,
			Valid: true,
		},
		Shift: sql.NullString{
			String: string(registration.Shift()),
			Valid:  true,
		},
		StudentID:            registration.Student().Id(),
		ServiceID:            registration.Service().Id(),
		MonthlyFee:           monthlyFee,
		InstallmentsQuantity: int32(registration.InstallmentsQuantity()),
		EnrollmentFee: sql.NullString{
			String: enrollmentFee,
			Valid:  true,
		},
		DueDate: sql.NullTime{
			Time:  registration.EnrollmentDueDate(),
			Valid: true,
		},
		MonthDuration: sql.NullInt32{
			Int32: int32(registration.MonthDuration()),
			Valid: true,
		},
		Status: registration.Status(),
		EnrollmentDate: sql.NullTime{
			Time:  registration.EnrollmentDate(),
			Valid: true,
		},
		SchoolYearID: uuid.NullUUID{
			UUID:  registration.Class().Id(),
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
		log.Println(err)
		return errors.New("failed to create registration")
	}

	rowAffected, err := r.UpdateOccupiedVacancies(classRoomModel.ID.String(), classRoomModel.VacanciesOccupied)
	if err != nil {
		log.Println(err)
		return errors.New("failed to update occupied vacancies")
	}

	if rowAffected < 1 {
		return errors.New("free vacancies not available")
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
	resp, err := r.db.ExecContext(
		context.Background(),
		query,
		vacanciesOccupied,
		time.Now().Format("2006-01-02 15:04:05"),
		classroomId,
	)

	if err != nil {
		return 0, err
	}

	rowsAffected, _ := resp.RowsAffected()

	return int(rowsAffected), err
}

func (r *RegistrationRepository) SearchStudentAlreadyRegistered(studentId uuid.UUID, classRoomId uuid.UUID) (string, error) {
	searchStudentAlready := models.SearchStudentAlreadyRegisteredParams{
		StudentID: studentId,
		ClassRoomID: uuid.NullUUID{
			UUID:  classRoomId,
			Valid: true,
		},
	}

	registrationCode, err := r.queues.SearchStudentAlreadyRegistered(context.Background(), searchStudentAlready)
	if err != nil {
		return "", err
	}

	return registrationCode, nil
}
