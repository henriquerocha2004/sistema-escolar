package postgres

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
	"log"
	"time"

	"github.com/henriquerocha2004/sistema-escolar/internal/school/common"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities"
)

type RoomRepository struct {
	db     *sql.DB
	queues *Queries
}

func NewRoomRepository(db *sql.DB) *RoomRepository {
	return &RoomRepository{
		db:     db,
		queues: New(db),
	}
}

func (r *RoomRepository) Create(room entities.Room) error {
	roomModel := CreateRoomParams{
		ID:          room.Id,
		Code:        room.Code,
		Description: room.Description,
		Capacity:    int32(room.Capacity),
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

	deleteParams := DeleteRoomParams{
		DeletedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		ID: idDelete,
	}

	err = r.queues.DeleteRoom(context.Background(), deleteParams)
	return err
}

func (r *RoomRepository) Update(room entities.Room) error {
	roomModel := &UpdateRoomParams{
		Code:        room.Code,
		Description: room.Description,
		Capacity:    int32(room.Capacity),
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		ID: room.Id,
	}

	err := r.queues.UpdateRoom(context.Background(), *roomModel)

	return err
}

func (r *RoomRepository) FindById(id string) (*entities.Room, error) {
	roomId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	roomModel, err := r.queues.FindOne(context.Background(), roomId)
	if err != nil {
		return nil, err
	}

	room := entities.Room{
		Id:          roomModel.ID,
		Code:        roomModel.Code,
		Description: roomModel.Description,
		Capacity:    int(roomModel.Capacity),
	}

	return &room, nil
}

func (r *RoomRepository) FindByCode(code string) (*entities.Room, error) {
	roomModel, err := r.queues.FindByCode(context.Background(), code)
	if err != nil {
		return nil, err
	}

	room := entities.Room{
		Id:          roomModel.ID,
		Code:        roomModel.Code,
		Description: roomModel.Description,
		Capacity:    int(roomModel.Capacity),
	}

	return &room, nil
}

func (r *RoomRepository) FindAll(paginator common.Pagination) (*common.RoomPaginationResult, error) {

	ctx, cancelQuery := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelQuery()

	query := "SELECT id, code, description, capacity FROM rooms WHERE (code like $1 OR description like $2) AND deleted_at IS NULL"
	filters := paginator.FiltersInSql()

	if filters != "" {
		query += filters
	}

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx,
		"%"+paginator.Search+"%",
		"%"+paginator.Search+"%",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []entities.Room

	for rows.Next() {
		var room entities.Room
		err = rows.Scan(&room.Id, &room.Code, &room.Description, &room.Capacity)
		if err != nil {
			return nil, err
		}

		rooms = append(rooms, room)
	}

	var total int
	queryCount := "SELECT COUNT(*) as total FROM rooms WHERE (code like $1 OR description like $2) AND deleted_at IS NULL"
	columnFilter := paginator.GetColumnFilter()
	if columnFilter != "" {
		queryCount += columnFilter
	}

	stmtCount, err := r.db.PrepareContext(ctx, queryCount)
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

	roomPaginationResult := common.RoomPaginationResult{
		Total: total,
		Rooms: rooms,
	}

	return &roomPaginationResult, nil
}

func (r *RoomRepository) SyncSchedule(scheduleDto dto.RoomScheduleDto) error {

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	r.queues.WithTx(tx)
	roomId, _ := uuid.Parse(scheduleDto.RoomId)
	schoolYearId, _ := uuid.Parse(scheduleDto.SchoolYear)

	unbindParams := UnbindScheduleParams{
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

		bindParams := BindScheduleParams{
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
