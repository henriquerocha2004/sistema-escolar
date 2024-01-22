package mocks

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/student"
	"github.com/stretchr/testify/mock"
)

type StudentRepository struct {
	mock.Mock
}

func (s *StudentRepository) Create(student student.Student) error {
	args := s.Called(student)
	return args.Error(0)
}
