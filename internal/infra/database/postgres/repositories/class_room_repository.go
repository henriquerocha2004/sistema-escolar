package repositories

import (
	"context"
	"database/sql"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/classroom"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/paginator"
	"time"

	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres/models"
)

type classRoomSearchModel struct {
	ID                uuid.UUID      `json:"id"`
	Status            string         `json:"status"`
	Active            bool           `json:"active"`
	Identification    string         `json:"identification"`
	Vacancies         int32          `json:"vacancies"`
	VacanciesOccupied int32          `json:"vacancies_occupied"`
	Shift             string         `json:"shift"`
	Level             string         `json:"level"`
	Localization      sql.NullString `json:"localization"`
	OpenDate          time.Time      `json:"open_date"`
	SchoolYearID      uuid.UUID      `json:"school_year_id"`
	RoomID            uuid.NullUUID  `json:"room_id"`
	ScheduleID        uuid.UUID      `json:"schedule_id"`
	CreatedAt         sql.NullTime   `json:"created_at"`
	UpdatedAt         sql.NullTime   `json:"updated_at"`
	Type              string         `json:"type"`
	Total             int
}

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

func (c *ClassRoomRepository) Create(classRoom classroom.ClassRoom) error {

	classRoomModel := models.CreateClassParams{
		ID:             classRoom.Id(),
		SchoolYearID:   classRoom.SchoolYearId(),
		ScheduleID:     classRoom.ScheduleId(),
		RoomID:         classRoom.RoomId(),
		Status:         classRoom.Status(),
		Identification: classRoom.Identification(),
		Level:          classRoom.Level(),
		Localization: sql.NullString{
			String: classRoom.Localization(),
			Valid:  true,
		},
		OpenDate:  classRoom.OpenDate(),
		Shift:     classRoom.Shift(),
		Vacancies: int32(classRoom.VacancyQuantity()),
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		Type: classRoom.TypeClass(),
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

func (c *ClassRoomRepository) Update(classRoom classroom.ClassRoom) error {
	classRoomModel := models.UpdateClassParams{
		ID:             classRoom.Id(),
		SchoolYearID:   classRoom.SchoolYearId(),
		ScheduleID:     classRoom.ScheduleId(),
		RoomID:         classRoom.RoomId(),
		Status:         classRoom.Status(),
		Identification: classRoom.Identification(),
		Level:          classRoom.Level(),
		Localization: sql.NullString{
			String: classRoom.Localization(),
			Valid:  true,
		},
		OpenDate:  time.Now(),
		Shift:     classRoom.Shift(),
		Vacancies: int32(classRoom.VacancyQuantity()),
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	err := c.queues.UpdateClass(context.Background(), classRoomModel)
	return err
}

func (c *ClassRoomRepository) FindById(id string) (*classroom.ClassRoom, error) {
	classId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	classRoomModel, err := c.queues.FindClassById(context.Background(), classId)
	if err != nil {
		return nil, err
	}

	classRoom, err := classroom.Load(
		classRoomModel.ID.String(),
		classRoomModel.Active,
		classRoomModel.Status,
		int(classRoomModel.VacanciesOccupied),
		int(classRoomModel.Vacancies),
		classRoomModel.OpenDate.String(),
		classRoomModel.Shift,
		classRoomModel.Level,
		classRoomModel.Identification,
		classRoomModel.SchoolYearID.String(),
		classRoomModel.RoomID.UUID.String(),
		classRoomModel.ScheduleID.String(),
		classRoomModel.Localization.String,
		classRoomModel.Type,
	)

	if err != nil {
		return nil, err
	}

	return classRoom, nil
}

// FindByIdLock FindByIdLock: Funcao que busca pelo ID da classe, porém ela faz o lock do registro no banco
// para evitar problemas de race conditions
func (c *ClassRoomRepository) FindByIdLock(id string) (*classroom.ClassRoom, error) {
	classId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	classRoomModel, err := c.queues.FindClassByIdLock(context.Background(), classId)
	if err != nil {
		return nil, err
	}

	classRoom, err := classroom.Load(
		classRoomModel.ID.String(),
		classRoomModel.Active,
		classRoomModel.Status,
		int(classRoomModel.VacanciesOccupied),
		int(classRoomModel.Vacancies),
		classRoomModel.OpenDate.String(),
		classRoomModel.Shift,
		classRoomModel.Level,
		classRoomModel.Identification,
		classRoomModel.SchoolYearID.String(),
		classRoomModel.RoomID.UUID.String(),
		classRoomModel.ScheduleID.String(),
		classRoomModel.Localization.String,
		classRoomModel.Type,
	)

	return classRoom, nil
}

func (c *ClassRoomRepository) FindAll(pagination paginator.Pagination) (*paginator.PaginationResult, error) {

	ctx, cancelQuery := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelQuery()

	query := `		
		SELECT 	id,status, active, identification,vacancies,
       			vacancies_occupied,shift,level,localization,
       			open_date,school_year_id,room_id,schedule_id,type,
       			COUNT(*) OVER() as total
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

	var classRoomsModel []classRoomSearchModel

	for rows.Next() {
		var classRoomModel classRoomSearchModel
		err = rows.Scan(
			&classRoomModel.ID,
			&classRoomModel.Status,
			&classRoomModel.Active,
			&classRoomModel.Identification,
			&classRoomModel.Vacancies,
			&classRoomModel.VacanciesOccupied,
			&classRoomModel.Shift,
			&classRoomModel.Level,
			&classRoomModel.Localization,
			&classRoomModel.OpenDate,
			&classRoomModel.SchoolYearID,
			&classRoomModel.RoomID,
			&classRoomModel.ScheduleID,
			&classRoomModel.Type,
			&classRoomModel.Total,
		)
		if err != nil {
			return nil, err
		}

		classRoomsModel = append(classRoomsModel, classRoomModel)
	}

	var classRooms []classroom.ClassRoom

	for _, classRoomModel := range classRoomsModel {
		classRoom, err := classroom.Load(
			classRoomModel.ID.String(),
			classRoomModel.Active,
			classRoomModel.Status,
			int(classRoomModel.VacanciesOccupied),
			int(classRoomModel.Vacancies),
			classRoomModel.OpenDate.String(),
			classRoomModel.Shift,
			classRoomModel.Level,
			classRoomModel.Identification,
			classRoomModel.SchoolYearID.String(),
			classRoomModel.RoomID.UUID.String(),
			classRoomModel.ScheduleID.String(),
			classRoomModel.Localization.String,
			classRoomModel.Type,
		)

		if err != nil {
			return nil, err
		}

		classRooms = append(classRooms, *classRoom)
	}

	classRoomPaginationResult := paginator.PaginationResult{
		Total: classRoomsModel[0].Total,
		Data:  classRooms,
	}

	return &classRoomPaginationResult, nil
}
