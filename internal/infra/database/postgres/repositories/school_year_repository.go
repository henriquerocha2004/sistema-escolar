package repositories

import (
	"context"
	"database/sql"
	"errors"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/schedule"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/schoolyear"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/paginator"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres/models"
)

type SchoolYearRepository struct {
	db     *sql.DB
	queues *models.Queries
}

type schoolYearSearchModel struct {
	Id      string
	Year    string
	StartAt time.Time
	EndAt   time.Time
	Total   int
}

func NewSchoolYearRepository(db *sql.DB) *SchoolYearRepository {
	return &SchoolYearRepository{
		db:     db,
		queues: models.New(db),
	}
}

func (s *SchoolYearRepository) Create(schoolYear *schoolyear.SchoolYear) error {

	schoolYearModel := models.CreateYearSchoolParams{
		ID:      schoolYear.Id(),
		Year:    schoolYear.Year(),
		StartAt: *schoolYear.StartAt(),
		EndAt:   *schoolYear.EndAt(),
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	err := s.queues.CreateYearSchool(context.Background(), schoolYearModel)
	if err != nil {
		return err
	}

	return nil
}

func (s *SchoolYearRepository) Delete(id string) error {
	schoolYearId, _ := uuid.Parse(id)

	deleteParams := models.DeleteYearSchoolParams{
		DeletedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		ID: schoolYearId,
	}

	err := s.queues.DeleteYearSchool(context.Background(), deleteParams)
	return err
}

func (s *SchoolYearRepository) Update(schoolYear *schoolyear.SchoolYear) error {
	schoolYearModel := models.UpdateSchoolYearParams{
		Year:    schoolYear.Year(),
		StartAt: *schoolYear.StartAt(),
		EndAt:   *schoolYear.EndAt(),
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		ID: schoolYear.Id(),
	}

	err := s.queues.UpdateSchoolYear(context.Background(), schoolYearModel)
	return err
}

func (s *SchoolYearRepository) FindById(id string) (*schoolyear.SchoolYear, error) {

	schoolYearId, _ := uuid.Parse(id)
	schoolYearModel, err := s.queues.FindOneSchoolYear(context.Background(), schoolYearId)
	if err != nil {
		return nil, err
	}

	schoolYear, err := schoolyear.Load(
		schoolYearModel.ID.String(),
		schoolYearModel.Year,
		schoolYearModel.StartAt.Format("2006-01-02"),
		schoolYearModel.EndAt.Format("2006-01-02"))

	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to retrieve class room")
	}

	return schoolYear, nil
}

func (s *SchoolYearRepository) FindByYear(year string) (*schoolyear.SchoolYear, error) {
	schoolYearModel, err := s.queues.FindByYear(context.Background(), year)
	if err != nil {
		return nil, err
	}

	schoolYear, err := schoolyear.New(
		schoolYearModel.Year,
		schoolYearModel.StartAt.Format("2006-01-02"),
		schoolYearModel.EndAt.Format("2006-01-02"))

	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to retrieve class room")
	}

	return schoolYear, nil
}

func (s *SchoolYearRepository) FindAll(pagination paginator.Pagination) (*paginator.PaginationResult, error) {
	ctx, cancelQuery := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelQuery()

	query := "SELECT id as id, year, start_at, end_at, COUNT(*) OVER() as total FROM school_year WHERE year like $1 AND deleted_at IS NULL"
	filters := pagination.FiltersInSql()

	if filters != "" {
		query += filters
	}

	stmt, err := s.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx,
		"%"+pagination.Search+"%",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schoolYearsSearchModels []schoolYearSearchModel
	var total int

	for rows.Next() {
		var schoolYearSearchModel schoolYearSearchModel
		err = rows.Scan(
			&schoolYearSearchModel.Id,
			&schoolYearSearchModel.Year,
			&schoolYearSearchModel.StartAt,
			&schoolYearSearchModel.EndAt,
			&total)
		if err != nil {
			return nil, err
		}

		schoolYearsSearchModels = append(schoolYearsSearchModels, schoolYearSearchModel)
	}

	var schoolYears []schoolyear.SchoolYear

	for _, schoolYearModel := range schoolYearsSearchModels {
		schoolYear, err := schoolyear.New(
			schoolYearModel.Year,
			schoolYearModel.StartAt.Format("2006-01-02"),
			schoolYearModel.EndAt.Format("2006-01-02"))

		if err != nil {
			return nil, err
		}

		err = schoolYear.SetId(schoolYearModel.Id)
		if err != nil {
			return nil, err
		}

		schoolYears = append(schoolYears, *schoolYear)
	}

	paginationResult := paginator.PaginationResult{
		Total: total,
		Data:  schoolYears,
	}

	return &paginationResult, nil
}

func (r *RoomRepository) SyncSchedule(scheduleDto schedule.RoomScheduleDto) error {

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	r.queues.WithTx(tx)
	roomId, _ := uuid.Parse(scheduleDto.RoomId)
	schoolYearId, _ := uuid.Parse(scheduleDto.SchoolYear)

	unbindParams := models.UnbindScheduleParams{
		RoomID:       roomId,
		SchoolYearID: schoolYearId,
	}

	err = r.queues.UnbindSchedule(context.Background(), unbindParams)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	for _, scheduleId := range scheduleDto.ScheduleIds {

		scheduleId, _ := uuid.Parse(scheduleId)

		bindParams := models.BindScheduleParams{
			RoomID:       roomId,
			ScheduleID:   scheduleId,
			SchoolYearID: schoolYearId,
		}

		err = r.queues.BindSchedule(context.Background(), bindParams)
		if err != nil {
			log.Println(err)
			_ = tx.Rollback()
			return err
		}
	}

	_ = tx.Commit()

	return nil
}
