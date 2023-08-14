package secretary

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/common"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities"
)

type ClassRoomActions struct {
	repository ClassRoomRepository
}

type ClassRoomActionsInterface interface {
	Create(dto dto.ClassRoomRequestDto) error
	Delete(id string) error
	Update(id string, requestDto dto.ClassRoomRequestDto) error
	Find(id string) (*entities.ClassRoom, error)
	FindAll(dtoRequest dto.PaginatorRequest) (*common.ClassRoomPaginationResult, error)
}

func NewClassRoomActions(repository ClassRoomRepository) *ClassRoomActions {
	return &ClassRoomActions{
		repository: repository,
	}
}

func (c *ClassRoomActions) Create(dto dto.ClassRoomRequestDto) error {
	classRoom := entities.ClassRoom{}
	classRoom.FillFromDto(dto)
	classRoom.Id = uuid.New()

	err := c.repository.Create(classRoom)
	if err != nil {
		log.Println(err)
		return errors.New("failed to save class")
	}

	return nil
}

func (c *ClassRoomActions) Update(id string, dto dto.ClassRoomRequestDto) error {
	classRoomId, _ := uuid.Parse(id)
	classRoom := entities.ClassRoom{}
	classRoom.FillFromDto(dto)
	classRoom.Id = classRoomId

	err := c.repository.Update(classRoom)
	if err != nil {
		log.Println(err)
		return errors.New("failed to update class room")
	}

	return nil
}

func (c *ClassRoomActions) Delete(id string) error {
	err := c.repository.Delete(id)
	if err != nil {
		log.Println(err)
		return errors.New("failed to delete class room")
	}

	return nil
}

func (c *ClassRoomActions) Find(id string) (*entities.ClassRoom, error) {
	classRoom, err := c.repository.FindById(id)
	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to retrieve class room")
	}

	return classRoom, nil
}

func (c *ClassRoomActions) FindAll(dtoRequest dto.PaginatorRequest) (*common.ClassRoomPaginationResult, error) {
	paginator := common.Pagination{}
	paginator.FillFromDto(dtoRequest)
	classRooms, err := c.repository.FindAll(paginator)
	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to retrieve class rooms")
	}

	return classRooms, nil
}
