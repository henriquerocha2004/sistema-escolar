package mocks

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities/schedule"
	dto2 "github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/schedule"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/paginator"
	"github.com/stretchr/testify/mock"
)

type ScheduleActionsMock struct {
	mock.Mock
}

func (s *ScheduleActionsMock) Create(dto dto2.ScheduleRequestDto) error {
	args := s.Called(dto)
	return args.Error(0)
}

func (s *ScheduleActionsMock) Delete(id string) error {
	args := s.Called(id)
	return args.Error(0)
}

func (s *ScheduleActionsMock) Update(id string, dto dto2.ScheduleRequestDto) error {
	args := s.Called(id, dto)
	return args.Error(0)
}

func (s *ScheduleActionsMock) FindOne(id string) (*schedule.ScheduleClass, error) {
	args := s.Called(id)
	return args.Get(0).(*schedule.ScheduleClass), args.Error(1)
}

func (s *ScheduleActionsMock) FindAll(dtoRequest paginator.PaginatorRequest) (*paginator.SchedulePaginationResult, error) {
	args := s.Called(dtoRequest)
	return args.Get(0).(*paginator.SchedulePaginationResult), args.Error(1)
}
