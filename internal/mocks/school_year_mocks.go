package mocks

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities"
	dto2 "github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/schoolyear"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/paginator"
	"github.com/stretchr/testify/mock"
)

type SchoolYearActionsMock struct {
	mock.Mock
}

func (r *SchoolYearActionsMock) Create(dto dto2.SchoolYearRequestDto) error {
	args := r.Called(dto)
	return args.Error(0)
}

func (r *SchoolYearActionsMock) Delete(id string) error {
	args := r.Called(id)
	return args.Error(0)
}

func (r *SchoolYearActionsMock) Update(id string, dto dto2.SchoolYearRequestDto) error {
	args := r.Called(id, dto)
	return args.Error(0)
}

func (r *SchoolYearActionsMock) FindAll(dtoRequest paginator.PaginatorRequest) (*paginator.SchoolYearPaginationResult, error) {
	args := r.Called(dtoRequest)
	return args.Get(0).(*paginator.SchoolYearPaginationResult), args.Error(1)
}

func (r *SchoolYearActionsMock) FindOne(id string) (*entities.SchoolYear, error) {
	args := r.Called(id)
	return args.Get(0).(*entities.SchoolYear), args.Error(1)
}
