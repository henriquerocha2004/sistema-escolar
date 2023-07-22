package mocks

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities"
	"github.com/stretchr/testify/mock"
)

type ScheduleActionsMock struct {
	mock.Mock
}

func (s *ScheduleActionsMock) Create(dto dto.ScheduleRequestDto) error {
	args := s.Called(dto)
	return args.Error(0)
}

func (s *ScheduleActionsMock) Delete(id string) error {
	args := s.Called(id)
	return args.Error(0)
}

func (s *ScheduleActionsMock) Update(id string, dto dto.ScheduleRequestDto) error {
	args := s.Called(id, dto)
	return args.Error(0)
}

func (s *ScheduleActionsMock) FindOne(id string) (*entities.ScheduleClass, error) {
	args := s.Called(id)
	return args.Get(0).(*entities.ScheduleClass), args.Error(1)
}

func (s *ScheduleActionsMock) FindAll(dtoRequest dto.PaginatorRequest) (*[]entities.ScheduleClass, error) {
	args := s.Called(dtoRequest)
	return args.Get(0).(*[]entities.ScheduleClass), args.Error(1)
}
