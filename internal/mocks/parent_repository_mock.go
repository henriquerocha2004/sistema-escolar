package mocks

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities"
	"github.com/stretchr/testify/mock"
)

type ParentRepository struct {
	mock.Mock
}

func (p *ParentRepository) Create(parent entities.Parent) error {
	args := p.Called(parent)
	return args.Error(0)
}
