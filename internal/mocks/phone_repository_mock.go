package mocks

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/value_objects"
	"github.com/stretchr/testify/mock"
)

type PhoneRepository struct {
	mock.Mock
}

func (p *PhoneRepository) Create(phone value_objects.Phone) error {
	args := p.Called(phone)
	return args.Error(0)
}
