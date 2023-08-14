package mocks

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/common"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities"
	"github.com/stretchr/testify/mock"
)

type ServiceActionsMock struct {
	mock.Mock
}

func (s *ServiceActionsMock) Create(dto dto.ServiceRequestDto) error {
	args := s.Called(dto)
	return args.Error(0)
}

func (s *ServiceActionsMock) Update(id string, dto dto.ServiceRequestDto) error {
	args := s.Called(id, dto)
	return args.Error(0)
}

func (s *ServiceActionsMock) Delete(id string) error {
	args := s.Called(id)
	return args.Error(0)
}

func (s *ServiceActionsMock) FindById(id string) (*entities.Service, error) {
	args := s.Called(id)
	return args.Get(0).(*entities.Service), args.Error(1)
}

func (s *ServiceActionsMock) FindAll(dtoRequest dto.PaginatorRequest) (*common.ServicePaginationResult, error) {
	args := s.Called(dtoRequest)
	return args.Get(0).(*common.ServicePaginationResult), args.Error(1)
}
