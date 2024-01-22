package mocks

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/schoolyear"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/paginator"
	"github.com/stretchr/testify/mock"
)

type SchoolYearActionsMock struct {
	mock.Mock
}

func (r *SchoolYearActionsMock) Create(dto schoolyear.Request) error {
	args := r.Called(dto)
	return args.Error(0)
}

func (r *SchoolYearActionsMock) Delete(id string) error {
	args := r.Called(id)
	return args.Error(0)
}

func (r *SchoolYearActionsMock) Update(id string, dto schoolyear.Request) error {
	args := r.Called(id, dto)
	return args.Error(0)
}

func (r *SchoolYearActionsMock) FindAll(dtoRequest paginator.PaginatorRequest) (*paginator.PaginationResult, error) {
	args := r.Called(dtoRequest)
	return args.Get(0).(*paginator.PaginationResult), args.Error(1)
}

func (r *SchoolYearActionsMock) FindOne(id string) (*schoolyear.SchoolYear, error) {
	args := r.Called(id)
	return args.Get(0).(*schoolyear.SchoolYear), args.Error(1)
}
