package roomService

import (
	"errors"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/room"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/schoolyear"

	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/paginator"
	"log"
)

type ServiceRoomInterface interface {
	Create(dto room.Request) error
	Delete(id string) error
	Update(id string, dto room.Request) error
	FindAll(dtoRequest paginator.PaginatorRequest) (*paginator.PaginationResult, error)
	FindById(id string) (*room.Room, error)
}

type ServiceRoom struct {
	repository           room.Repository
	schoolYearRepository schoolyear.Repository
}

func New(repo room.Repository) *ServiceRoom {
	return &ServiceRoom{
		repository: repo,
	}
}

func (r *ServiceRoom) Create(dto room.Request) error {

	rom, err := room.New(dto.Code, dto.Description, dto.Capacity)
	if err != nil {
		return err
	}

	roomDuplicated, _ := r.repository.FindByCode(rom.Code())

	if roomDuplicated != nil {
		return errors.New("room already exists")
	}

	err = r.repository.Create(*rom)
	if err != nil {
		log.Println(err)
		return errors.New("failed to create room")
	}

	return nil
}

func (r *ServiceRoom) Delete(id string) error {
	err := r.repository.Delete(id)
	if err != nil {
		log.Println(err)
		err = errors.New("error in delete room")
	}

	return err
}

func (r *ServiceRoom) Update(id string, dto room.Request) error {
	rom, err := room.New(dto.Code, dto.Description, dto.Capacity)
	if err != nil {
		return err
	}

	err = rom.ChangeId(id)
	if err != nil {
		return err
	}

	err = r.repository.Update(*rom)

	if err != nil {
		log.Println(err)
		return errors.New("failed to update room")
	}

	return nil
}

func (r *ServiceRoom) FindAll(dtoRequest paginator.PaginatorRequest) (*paginator.PaginationResult, error) {

	pg := paginator.Pagination{}
	pg.FillFromDto(dtoRequest)
	rooms, err := r.repository.FindAll(pg)

	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (r *ServiceRoom) FindById(id string) (*room.Room, error) {
	rom, err := r.repository.FindById(id)
	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to retrieve room information")
	}

	return rom, nil
}
