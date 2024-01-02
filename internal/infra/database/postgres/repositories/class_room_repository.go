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

type ClassRoomRepository struct {
	db     *sql.DB
	queues *models.Queries
}

func NewClassRoomRepository(db *sql.DB) *ClassRoomRepository {
	return &ClassRoomRepository{
		db:     db,
		queues: models.New(db),
	}
}

func (c *ClassRoomRepository) Create(classRoom entities.ClassRoom) error {

	classRoomModel := models.CreateClassParams{
		ID:             classRoom.Id,
		SchoolYearID:   classRoom.SchoolYearId,
		ScheduleID:     classRoom.ScheduleId,
		RoomID:         classRoom.RoomId,
		Status:         classRoom.Status,
		Identification: classRoom.Identification,
		Level:          classRoom.Level,
		Localization: sql.NullString{
			String: classRoom.Localization,
			Valid:  true,
		},
		OpenDate:  *classRoom.OpenDate,
		Shift:     classRoom.Shift,
		Vacancies: int32(classRoom.VacancyQuantity),
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		Type: classRoom.Type,
	}

	err := c.queues.CreateClass(context.Background(), classRoomModel)
	if err != nil {
		return err
	}

	return nil
}

func (c *ClassRoomRepository) Delete(id string) error {
	classId, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	deleteParams := models.DeleteClassParams{
		DeletedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		ID: classId,
	}

	err = c.queues.DeleteClass(context.Background(), deleteParams)
	return err
}

func (c *ClassRoomRepository) Update(classRoom entities.ClassRoom) error {
	classRoomModel := models.UpdateClassParams{
		ID:             classRoom.Id,
		SchoolYearID:   classRoom.SchoolYearId,
		ScheduleID:     classRoom.ScheduleId,
		RoomID:         classRoom.RoomId,
		Status:         classRoom.Status,
		Identification: classRoom.Identification,
		Level:          classRoom.Level,
		Localization: sql.NullString{
			String: classRoom.Localization,
			Valid:  true,
		},
		OpenDate:  time.Now(),
		Shift:     classRoom.Shift,
		Vacancies: int32(classRoom.VacancyQuantity),
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	err := c.queues.UpdateClass(context.Background(), classRoomModel)
	return err
}

func (c *ClassRoomRepository) FindById(id string) (*entities.ClassRoom, error) {
	classId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	classRoomModel, err := c.queues.FindClassById(context.Background(), classId)
	if err != nil {
		return nil, err
	}

	classRoom := entities.ClassRoom{
		Id:              classRoomModel.ID,
		VacancyQuantity: int(classRoomModel.Vacancies),
		Shift:           classRoomModel.Shift,
		OpenDate:        &classRoomModel.OpenDate,
		Status:          classRoomModel.Status,
		Level:           classRoomModel.Level,
		Identification:  classRoomModel.Identification,
		SchoolYearId:    classRoomModel.SchoolYearID,
		RoomId:          classRoomModel.RoomID,
		ScheduleId:      classRoomModel.ScheduleID,
		Localization:    classRoomModel.Localization.String,
		Type:            classRoomModel.Type,
	}

	return &classRoom, nil
}

// FindByIdLock: Funcao que busca pelo ID da classe, porém ela faz o lock do registro no banco
// para evitar problemas de race conditions
func (c *ClassRoomRepository) FindByIdLock(id string) (*entities.ClassRoom, error) {
	classId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	classRoomModel, err := c.queues.FindClassByIdLock(context.Background(), classId)
	if err != nil {
		return nil, err
	}

	classRoom := entities.ClassRoom{
		Id:              classRoomModel.ID,
		VacancyQuantity: int(classRoomModel.Vacancies),
		Shift:           classRoomModel.Shift,
		OpenDate:        &classRoomModel.OpenDate,
		Status:          classRoomModel.Status,
		Level:           classRoomModel.Level,
		Identification:  classRoomModel.Identification,
		SchoolYearId:    classRoomModel.SchoolYearID,
		RoomId:          classRoomModel.RoomID,
		ScheduleId:      classRoomModel.ScheduleID,
		Localization:    classRoomModel.Localization.String,
		Type:            classRoomModel.Type,
	}

	return &classRoom, nil
}

func (c *ClassRoomRepository) FindAll(pagination common.Pagination) (*common.ClassRoomPaginationResult, error) {
	ctx, cancelQuery := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelQuery()

	query := `		
		SELECT 	id,status,identification,vacancies,
       			vacancies_occupied,shift,level,localization,
       			open_date,school_year_id,room_id,schedule_id,type
			FROM class_room
    	WHERE localization like $1 AND deleted_at IS NULL 
	`
	filters := pagination.FiltersInSql()

	if filters != "" {
		query += filters
	}

	stmt, err := c.db.PrepareContext(ctx, query)
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

	var classRooms []entities.ClassRoom

	for rows.Next() {
		var classRoom entities.ClassRoom
		err = rows.Scan(
			&classRoom.Id,
			&classRoom.Status,
			&classRoom.Identification,
			&classRoom.VacancyQuantity,
			&classRoom.OccupiedVacancy,
			&classRoom.Shift,
			&classRoom.Level,
			&classRoom.Localization,
			&classRoom.OpenDate,
			&classRoom.SchoolYearId,
			&classRoom.RoomId,
			&classRoom.ScheduleId,
			&classRoom.Type)
		if err != nil {
			return nil, err
		}

		classRooms = append(classRooms, classRoom)
	}

	var total int
	queryCount := "SELECT COUNT(*) as total FROM class_room WHERE localization like $1"
	columnFilter := pagination.GetColumnFilter()
	if columnFilter != "" {
		queryCount += columnFilter
	}

	stmtCount, err := c.db.PrepareContext(ctx, queryCount)
	if err != nil {
		return nil, err
	}

	defer stmtCount.Close()

	rows, err = stmtCount.QueryContext(ctx,
		"%"+pagination.Search+"%",
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

	classRoomPaginationResult := common.ClassRoomPaginationResult{
		Total:      total,
		ClassRooms: classRooms,
	}

	return &classRoomPaginationResult, nil
}
