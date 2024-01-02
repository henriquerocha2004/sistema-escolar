package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"

	"github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres/models"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/common"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities"
)

type ScheduleRoomRepository struct {
	db     *sql.DB
	queues *models.Queries
}

func NewScheduleRoomRepository(connection *sql.DB) *ScheduleRoomRepository {
	return &ScheduleRoomRepository{
		db:     connection,
		queues: models.New(connection),
	}
}

func (s *ScheduleRoomRepository) Create(schedule entities.ScheduleClass) error {
	scheduleModel := models.CreateScheduleParams{
		ID:           schedule.Id,
		Description:  schedule.Description,
		SchoolYearID: schedule.SchoolYear,
		Schedule:     schedule.Schedule,
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	err := s.queues.CreateSchedule(context.Background(), scheduleModel)
	if err != nil {
		return err
	}

	return nil
}

func (s *ScheduleRoomRepository) Delete(id string) error {

	scheduleId, _ := uuid.Parse(id)

	deleteParams := models.DeleteScheduleParams{
		DeletedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		ID: scheduleId,
	}

	err := s.queues.DeleteSchedule(context.Background(), deleteParams)
	return err
}

func (s *ScheduleRoomRepository) Update(schedule entities.ScheduleClass) error {
	scheduleModel := models.UpdateScheduleParams{
		Description:  schedule.Description,
		Schedule:     schedule.Schedule,
		SchoolYearID: schedule.SchoolYear,
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		ID: schedule.Id,
	}

	err := s.queues.UpdateSchedule(context.Background(), scheduleModel)
	return err
}

func (s *ScheduleRoomRepository) FindById(id string) (*entities.ScheduleClass, error) {
	scheduleId, _ := uuid.Parse(id)

	scheduleModel, err := s.queues.FindOneSchedule(context.Background(), scheduleId)
	if err != nil {
		return nil, err
	}

	schedule := entities.ScheduleClass{
		Id:          scheduleModel.ScheduleID,
		Description: scheduleModel.Description,
		Schedule:    scheduleModel.Schedule,
		SchoolYear:  scheduleModel.ID,
	}

	return &schedule, nil
}

func (s *ScheduleRoomRepository) FindAll(paginator common.Pagination) (*common.SchedulePaginationResult, error) {
	ctx, cancelQuery := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelQuery()

	query := "SELECT class_schedule.id, description, schedule, school_year.id FROM class_schedule JOIN school_year ON school_year.id = class_schedule.school_year_id WHERE (class_schedule.description like $1 OR class_schedule.schedule = $2) AND class_schedule.deleted_at IS NULL"
	filters := paginator.FiltersInSql()

	if filters != "" {
		query += filters
	}

	stmt, err := s.db.PrepareContext(context.Background(), query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.QueryContext(context.Background(),
		"%"+paginator.Search+"%",
		"%"+paginator.Search+"%",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []entities.ScheduleClass

	for rows.Next() {
		var schedule entities.ScheduleClass
		err = rows.Scan(&schedule.Id, &schedule.Description, &schedule.Schedule, &schedule.SchoolYear)
		if err != nil {
			return nil, err
		}

		schedules = append(schedules, schedule)
	}

	var total int
	queryCount := "SELECT COUNT(*) FROM class_schedule JOIN school_year ON school_year.id = class_schedule.school_year_id WHERE (class_schedule.description like $1 OR class_schedule.schedule = $2) AND class_schedule.deleted_at IS NULL"
	columnFilter := paginator.GetColumnFilter()
	if columnFilter != "" {
		queryCount += columnFilter
	}

	stmtCount, err := s.db.PrepareContext(ctx, queryCount)
	if err != nil {
		return nil, err
	}

	defer stmtCount.Close()

	rows, err = stmtCount.QueryContext(ctx,
		"%"+paginator.Search+"%",
		"%"+paginator.Search+"%",
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&total)
		if err != nil {
			return nil, err
		}
	}

	schedulePaginationResult := common.SchedulePaginationResult{
		Total:     total,
		Schedules: schedules,
	}

	return &schedulePaginationResult, nil
}
