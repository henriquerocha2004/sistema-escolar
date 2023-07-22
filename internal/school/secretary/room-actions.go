package secretary

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/common"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities"
)

type RoomActionsInterface interface {
	Create(dto dto.RoomRequestDto) error
	Delete(id string) error
	Update(id string, dto dto.RoomRequestDto) error
	FindAll(dtoRequest dto.PaginatorRequest) (*[]entities.Room, error)
	FindById(id string) (*entities.Room, error)
	SyncSchedule(scheduleRoomDto dto.RoomScheduleDto) error
}

type RoomActions struct {
	repository           RoomRepository
	schoolYearRepository SchoolYearRepository
}

func NewRoomActions(repo RoomRepository) RoomActionsInterface {
	return &RoomActions{
		repository: repo,
	}
}

func (ra *RoomActions) Create(dto dto.RoomRequestDto) error {
	room := entities.Room{}
	room.FillFromDto(dto)
	room.Id = uuid.New()

	roomDuplicated, _ := ra.repository.FindByCode(room.Code)

	if roomDuplicated != nil {
		err := errors.New("room already exists")
		log.Println(err)
		return err
	}

	err := ra.repository.Create(room)
	if err != nil {
		log.Println(err)
		err := errors.New("failed to create room")
		return err
	}

	return nil
}

func (ra *RoomActions) Delete(id string) error {
	err := ra.repository.Delete(id)
	if err != nil {
		log.Println(err)
		err = errors.New("error in delete room")
	}

	return err
}

func (ra *RoomActions) Update(id string, dto dto.RoomRequestDto) error {
	room := entities.Room{}
	room.FillFromDto(dto)
	roomId, err := uuid.Parse(id)
	if err != nil {
		log.Println(err)
		return errors.New("invalid id provided")
	}
	room.Id = roomId
	err = ra.repository.Update(room)

	if err != nil {
		log.Println(err)
		err = errors.New("failed to update room")
	}

	return err
}

func (ra *RoomActions) FindAll(dtoRequest dto.PaginatorRequest) (*[]entities.Room, error) {

	paginator := common.Pagination{}
	paginator.FillFromDto(dtoRequest)
	rooms, err := ra.repository.FindAll(paginator)

	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (ra *RoomActions) FindById(id string) (*entities.Room, error) {
	room, err := ra.repository.FindById(id)
	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to retrieve room information")
	}

	return room, nil
}

func (ra *RoomActions) SyncSchedule(scheduleRoomDto dto.RoomScheduleDto) error {
	err := ra.repository.SyncSchedule(scheduleRoomDto)
	if err != nil {
		log.Println(err)
		return errors.New("failed at sync schedules with room")
	}

	return nil
}
