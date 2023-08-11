package mocks

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/common"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities"
	"github.com/stretchr/testify/mock"
)

type SchoolYearActionsMock struct {
	mock.Mock
}

func (r *SchoolYearActionsMock) Create(dto dto.SchoolYearRequestDto) error {
	args := r.Called(dto)
	return args.Error(0)
}

func (r *SchoolYearActionsMock) Delete(id string) error {
	args := r.Called(id)
	return args.Error(0)
}

func (r *SchoolYearActionsMock) Update(id string, dto dto.SchoolYearRequestDto) error {
	args := r.Called(id, dto)
	return args.Error(0)
}

func (r *SchoolYearActionsMock) FindAll(dtoRequest dto.PaginatorRequest) (*common.SchoolYearPaginationResult, error) {
	args := r.Called(dtoRequest)
	return args.Get(0).(*common.SchoolYearPaginationResult), args.Error(1)
}

func (r *SchoolYearActionsMock) FindOne(id string) (*entities.SchoolYear, error) {
	args := r.Called(id)
	return args.Get(0).(*entities.SchoolYear), args.Error(1)
}
