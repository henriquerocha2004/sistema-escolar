package mocks

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/financial/service"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/paginator"
	"github.com/stretchr/testify/mock"
)

type ServiceRepository struct {
	mock.Mock
}

func (s *ServiceRepository) Create(service service.Service) error {
	args := s.Called(service)
	return args.Error(0)
}

func (s *ServiceRepository) Delete(id string) error {
	args := s.Called(id)
	return args.Error(0)
}

func (s *ServiceRepository) Update(service service.Service) error {
	args := s.Called(service)
	return args.Error(0)
}

func (s *ServiceRepository) FindById(id string) (*service.Service, error) {
	args := s.Called(id)
	return args.Get(0).(*service.Service), args.Error(1)
}

func (s *ServiceRepository) FindAll(pagination paginator.Pagination) (*paginator.PaginationResult, error) {
	args := s.Called(pagination)
	return args.Get(0).(*paginator.PaginationResult), args.Error(1)
}
