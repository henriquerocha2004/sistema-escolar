package mocks

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/room"
	dto2 "github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/schedule"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/paginator"
	"github.com/stretchr/testify/mock"
)

type RoomActionsMock struct {
	mock.Mock
}

func (r *RoomActionsMock) Create(dto room.Request) error {
	args := r.Called(dto)
	return args.Error(0)
}

func (r *RoomActionsMock) Delete(id string) error {
	args := r.Called(id)
	return args.Error(0)
}

func (r *RoomActionsMock) Update(id string, dto room.Request) error {
	args := r.Called(id, dto)
	return args.Error(0)
}

func (r *RoomActionsMock) FindAll(dtoRequest paginator.PaginatorRequest) (*paginator.PaginationResult, error) {
	args := r.Called(dtoRequest)
	return args.Get(0).(*paginator.PaginationResult), args.Error(1)
}

func (r *RoomActionsMock) FindById(id string) (*room.Room, error) {
	args := r.Called(id)
	return args.Get(0).(*room.Room), args.Error(1)
}

func (r *RoomActionsMock) SyncSchedule(scheduleRoomDto dto2.RoomScheduleDto) error {
	args := r.Called(scheduleRoomDto)
	return args.Error(0)
}
