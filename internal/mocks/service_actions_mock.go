package mocks

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/financial/service"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/paginator"
	"github.com/stretchr/testify/mock"
)

type ServiceActionsMock struct {
	mock.Mock
}

func (s *ServiceActionsMock) Create(dto service.Request) error {
	args := s.Called(dto)
	return args.Error(0)
}

func (s *ServiceActionsMock) Update(id string, dto service.Request) error {
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

func (s *ServiceActionsMock) FindAll(dtoRequest paginator.PaginatorRequest) (*paginator.PaginationResult, error) {
	args := s.Called(dtoRequest)
	return args.Get(0).(*paginator.PaginationResult), args.Error(1)
}
