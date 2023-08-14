package mocks

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/common"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities"
	"github.com/stretchr/testify/mock"
)

type RoomActionsMock struct {
	mock.Mock
}

func (r *RoomActionsMock) Create(dto dto.RoomRequestDto) error {
	args := r.Called(dto)
	return args.Error(0)
}

func (r *RoomActionsMock) Delete(id string) error {
	args := r.Called(id)
	return args.Error(0)
}

func (r *RoomActionsMock) Update(id string, dto dto.RoomRequestDto) error {
	args := r.Called(id, dto)
	return args.Error(0)
}

func (r *RoomActionsMock) FindAll(dtoRequest dto.PaginatorRequest) (*common.RoomPaginationResult, error) {
	args := r.Called(dtoRequest)
	return args.Get(0).(*common.RoomPaginationResult), args.Error(1)
}

func (r *RoomActionsMock) FindById(id string) (*entities.Room, error) {
	args := r.Called(id)
	return args.Get(0).(*entities.Room), args.Error(1)
}

func (r *RoomActionsMock) SyncSchedule(scheduleRoomDto dto.RoomScheduleDto) error {
	args := r.Called(scheduleRoomDto)
	return args.Error(0)
}
