package repositories

import (
	"context"
	"database/sql"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/schedule"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/paginator"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres/models"
)

type ScheduleRoomRepository struct {
	db     *sql.DB
	queues *models.Queries
}

type scheduleSearchModel struct {
	ID           uuid.UUID `json:"id"`
	Description  string    `json:"description"`
	StartAt      time.Time `json:"start_at"`
	EndAt        time.Time `json:"end_at"`
	SchoolYearID uuid.UUID `json:"school_year_id"`
	Total        int       `json:"total"`
}

func NewScheduleRoomRepository(connection *sql.DB) *ScheduleRoomRepository {
	return &ScheduleRoomRepository{
		db:     connection,
		queues: models.New(connection),
	}
}

func (s *ScheduleRoomRepository) Create(schedule schedule.ScheduleClass) error {

	stDate, _ := s.parseToTime(schedule.StartAt())
	edDate, _ := s.parseToTime(schedule.EndAt())

	scheduleModel := models.CreateScheduleParams{
		ID:           schedule.Id(),
		Description:  schedule.Description(),
		SchoolYearID: schedule.SchoolYearId(),
		StartAt:      stDate,
		EndAt:        edDate,
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

func (s *ScheduleRoomRepository) Update(schedule schedule.ScheduleClass) error {

	stDate, _ := s.parseToTime(schedule.StartAt())
	edDate, _ := s.parseToTime(schedule.EndAt())

	scheduleModel := models.UpdateScheduleParams{
		Description:  schedule.Description(),
		StartAt:      stDate,
		EndAt:        edDate,
		SchoolYearID: schedule.SchoolYearId(),
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		ID: schedule.Id(),
	}

	err := s.queues.UpdateSchedule(context.Background(), scheduleModel)
	return err
}

func (s *ScheduleRoomRepository) FindById(id string) (*schedule.ScheduleClass, error) {
	scheduleId, _ := uuid.Parse(id)

	scheduleModel, err := s.queues.FindOneSchedule(context.Background(), scheduleId)
	if err != nil {
		return nil, err
	}

	schedule, err := schedule.Load(
		scheduleModel.ScheduleID.String(),
		scheduleModel.Description,
		scheduleModel.StartAt.Format("15:01:05"),
		scheduleModel.EndAt.Format("15:01:05"),
		scheduleModel.ID.String(),
	)

	return schedule, nil
}

func (s *ScheduleRoomRepository) FindAll(pagination paginator.Pagination) (*paginator.PaginationResult, error) {
	ctx, cancelQuery := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelQuery()

	query := `SELECT 
    			class_schedule.id, description, class_schedule.start_at, class_schedule.end_at, school_year.id, COUNT(*) OVER() as total
			FROM class_schedule 
			    JOIN school_year ON school_year.id = class_schedule.school_year_id 
			WHERE (class_schedule.description like $1) 
			  AND class_schedule.deleted_at IS NULL`
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

	var schedulesModel []scheduleSearchModel

	for rows.Next() {
		var scheduleModel scheduleSearchModel
		err = rows.Scan(
			&scheduleModel.ID,
			&scheduleModel.Description,
			&scheduleModel.StartAt,
			&scheduleModel.EndAt,
			&scheduleModel.SchoolYearID,
			&scheduleModel.Total)
		if err != nil {
			return nil, err
		}

		schedulesModel = append(schedulesModel, scheduleModel)
	}

	var schedules []schedule.ScheduleClass

	for _, scheduleModel := range schedulesModel {
		sch, err := schedule.Load(
			scheduleModel.ID.String(),
			scheduleModel.Description,
			scheduleModel.StartAt.Format("15:04:05"),
			scheduleModel.EndAt.Format("15:04:05"),
			scheduleModel.SchoolYearID.String(),
		)

		if err != nil {
			return nil, err
		}

		schedules = append(schedules, *sch)
	}

	schedulePaginationResult := paginator.PaginationResult{
		Total: schedulesModel[0].Total,
		Data:  schedules,
	}

	return &schedulePaginationResult, nil
}

func (r *ScheduleRoomRepository) SyncSchedule(scheduleDto schedule.RoomScheduleDto) error {

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

func (s *ScheduleRoomRepository) parseToTime(t string) (time.Time, error) {
	return time.Parse("15:01:05", t)
}
