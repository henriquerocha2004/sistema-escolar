package mocks

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/common"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities"
	"github.com/stretchr/testify/mock"
)

type ServiceRepository struct {
	mock.Mock
}

func (s *ServiceRepository) Create(service entities.Service) error {
	args := s.Called(service)
	return args.Error(0)
}

func (s *ServiceRepository) Delete(id string) error {
	args := s.Called(id)
	return args.Error(0)
}

func (s *ServiceRepository) Update(service entities.Service) error {
	args := s.Called(service)
	return args.Error(0)
}

func (s *ServiceRepository) FindById(id string) (*entities.Service, error) {
	args := s.Called(id)
	return args.Get(0).(*entities.Service), args.Error(1)
}

func (s *ServiceRepository) FindAll(paginator common.Pagination) (*common.ServicePaginationResult, error) {
	args := s.Called(paginator)
	return args.Get(0).(*common.ServicePaginationResult), args.Error(1)
}
