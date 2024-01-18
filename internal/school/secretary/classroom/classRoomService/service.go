package classRoomService

import (
	"errors"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/classroom"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/paginator"
	"log"
)

type ServiceClassRoom struct {
	repository classroom.Repository
}

type ServiceClassRoomInterface interface {
	Create(dto classroom.Request) error
	Delete(id string) error
	Update(id string, requestDto classroom.Request) error
	Find(id string) (*classroom.ClassRoom, error)
	FindAll(dtoRequest paginator.PaginatorRequest) (*paginator.PaginationResult, error)
}

func New(repository classroom.Repository) *ServiceClassRoom {
	return &ServiceClassRoom{
		repository: repository,
	}
}

func (c *ServiceClassRoom) Create(dto classroom.Request) error {
	classRoom, err := classroom.New(dto.VacancyQuantity,
		dto.Shift,
		dto.Level,
		dto.Identification,
		dto.SchoolYearId,
		dto.RoomId,
		dto.ScheduleId,
		dto.Localization,
		dto.Type)

	if err != nil {
		return err
	}

	err = c.repository.Create(*classRoom)
	if err != nil {
		log.Println(err)
		return errors.New("failed to save class")
	}

	return nil
}

func (c *ServiceClassRoom) Update(id string, dto classroom.Request) error {
	classRoom, err := classroom.New(dto.VacancyQuantity,
		dto.Shift,
		dto.Level,
		dto.Identification,
		dto.SchoolYearId,
		dto.RoomId,
		dto.ScheduleId,
		dto.Localization,
		dto.Type)

	if err != nil {
		return err
	}

	err = classRoom.ChangeRoomId(id)
	if err != nil {
		return err
	}

	err = c.repository.Update(*classRoom)
	if err != nil {
		log.Println(err)
		return errors.New("failed to update class room")
	}

	return nil
}

func (c *ServiceClassRoom) Delete(id string) error {
	err := c.repository.Delete(id)
	if err != nil {
		log.Println(err)
		return errors.New("failed to delete class room")
	}

	return nil
}

func (c *ServiceClassRoom) Find(id string) (*classroom.ClassRoom, error) {
	classRoom, err := c.repository.FindById(id)
	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to retrieve class room")
	}

	return classRoom, nil
}

func (c *ServiceClassRoom) FindAll(dtoRequest paginator.PaginatorRequest) (*paginator.PaginationResult, error) {
	pg := paginator.Pagination{}
	pg.FillFromDto(dtoRequest)
	classRooms, err := c.repository.FindAll(pg)
	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to retrieve class rooms")
	}

	return classRooms, nil
}
