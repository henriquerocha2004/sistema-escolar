package mocks

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities"
	"github.com/stretchr/testify/mock"
)

type StudentRepository struct {
	mock.Mock
}

func (s *StudentRepository) Create(student entities.Student) error {
	args := s.Called(student)
	return args.Error(0)
}
