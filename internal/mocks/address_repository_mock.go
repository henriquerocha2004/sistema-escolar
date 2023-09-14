package mocks

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/value_objects"
	"github.com/stretchr/testify/mock"
)

type AddressRepository struct {
	mock.Mock
}

func (a *AddressRepository) Create(address value_objects.Address) error {
	args := a.Called(address)
	return args.Error(0)
}
