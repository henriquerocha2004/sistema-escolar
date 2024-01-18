package repositories

import (
	"context"
	"database/sql"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/room"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/paginator"
	"time"

	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres/models"
)

type RoomRepository struct {
	db     *sql.DB
	queues *models.Queries
}

type roomSearchModel struct {
	ID          uuid.UUID `json:"id"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	Capacity    int32     `json:"capacity"`
	Total       int       `json:"total"`
}

func NewRoomRepository(db *sql.DB) *RoomRepository {
	return &RoomRepository{
		db:     db,
		queues: models.New(db),
	}
}

func (r *RoomRepository) Create(room room.Room) error {
	roomModel := models.CreateRoomParams{
		ID:          room.Id(),
		Code:        room.Code(),
		Description: room.Description(),
		Capacity:    int32(room.Capacity()),
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	err := r.queues.CreateRoom(context.Background(), roomModel)

	if err != nil {
		return err
	}

	return nil
}

func (r *RoomRepository) Delete(id string) error {

	idDelete, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	deleteParams := models.DeleteRoomParams{
		DeletedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		ID: idDelete,
	}

	err = r.queues.DeleteRoom(context.Background(), deleteParams)
	return err
}

func (r *RoomRepository) Update(room room.Room) error {
	roomModel := &models.UpdateRoomParams{
		Code:        room.Code(),
		Description: room.Description(),
		Capacity:    int32(room.Capacity()),
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		ID: room.Id(),
	}

	err := r.queues.UpdateRoom(context.Background(), *roomModel)

	return err
}

func (r *RoomRepository) FindById(id string) (*room.Room, error) {
	roomId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	roomModel, err := r.queues.FindOne(context.Background(), roomId)
	if err != nil {
		return nil, err
	}

	room, err := room.Load(
		roomModel.ID.String(),
		roomModel.Code,
		roomModel.Description,
		int(roomModel.Capacity),
	)

	return room, err
}

func (r *RoomRepository) FindByCode(code string) (*room.Room, error) {
	roomModel, err := r.queues.FindByCode(context.Background(), code)
	if err != nil {
		return nil, err
	}

	room, err := room.Load(
		roomModel.ID.String(),
		roomModel.Code,
		roomModel.Description,
		int(roomModel.Capacity),
	)

	return room, err
}

func (r *RoomRepository) FindAll(pagination paginator.Pagination) (*paginator.PaginationResult, error) {

	ctx, cancelQuery := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelQuery()

	query := `SELECT id, code, description, capacity, COUNT(*) OVER() as total 
				FROM rooms 
		WHERE (code like $1 OR description like $2) 
		   AND deleted_at IS NULL`

	filters := pagination.FiltersInSql()

	if filters != "" {
		query += filters
	}

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx,
		"%"+pagination.Search+"%",
		"%"+pagination.Search+"%",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roomsModels []roomSearchModel

	for rows.Next() {
		var roomModel roomSearchModel
		err = rows.Scan(&roomModel.ID, &roomModel.Code, &roomModel.Description, &roomModel.Capacity, &roomModel.Total)
		if err != nil {
			return nil, err
		}

		roomsModels = append(roomsModels, roomModel)
	}

	var rooms []room.Room

	for _, roomModel := range roomsModels {

		room, err := room.Load(
			roomModel.ID.String(),
			roomModel.Code,
			roomModel.Description,
			int(roomModel.Capacity),
		)

		if err != nil {
			return nil, err
		}

		rooms = append(rooms, *room)
	}

	roomPaginationResult := paginator.PaginationResult{
		Total: roomsModels[0].Total,
		Data:  rooms,
	}

	return &roomPaginationResult, nil
}
