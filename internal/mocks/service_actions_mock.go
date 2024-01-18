package mocks

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities/service"
	service2 "github.com/henriquerocha2004/sistema-escolar/internal/school/financial/service"
	dto2 "github.com/henriquerocha2004/sistema-escolar/internal/school/shared/paginator"
	"github.com/stretchr/testify/mock"
)

type ServiceActionsMock struct {
	mock.Mock
}

func (s *ServiceActionsMock) Create(dto service2.ServiceRequestDto) error {
	args := s.Called(dto)
	return args.Error(0)
}

func (s *ServiceActionsMock) Update(id string, dto service2.ServiceRequestDto) error {
	args := s.Called(id, dto)
	return args.Error(0)
}

func (s *ServiceActionsMock) Delete(id string) error {
	args := s.Called(id)
	return args.Error(0)
}

func (s *ServiceActionsMock) FindById(id string) (*service.Service, error) {
	args := s.Called(id)
	return args.Get(0).(*service.Service), args.Error(1)
}

func (s *ServiceActionsMock) FindAll(dtoRequest dto2.PaginatorRequest) (*dto2.ServicePaginationResult, error) {
	args := s.Called(dtoRequest)
	return args.Get(0).(*dto2.ServicePaginationResult), args.Error(1)
}
