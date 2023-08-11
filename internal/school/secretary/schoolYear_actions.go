package secretary

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/common"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities"
)

type SchoolYearActionsInterface interface {
	Create(dto dto.SchoolYearRequestDto) error
	Delete(id string) error
	Update(id string, dto dto.SchoolYearRequestDto) error
	FindOne(id string) (*entities.SchoolYear, error)
	FindAll(dtoRequest dto.PaginatorRequest) (*common.SchoolYearPaginationResult, error)
}

type SchoolYearActions struct {
	repository SchoolYearRepository
}

func NewSchoolYearActions(repository SchoolYearRepository) *SchoolYearActions {
	return &SchoolYearActions{
		repository: repository,
	}
}

func (s *SchoolYearActions) Create(dto dto.SchoolYearRequestDto) error {
	schoolYear := entities.SchoolYear{}
	schoolYear.FillFromDto(dto)
	schoolYear.Id = uuid.New()

	err := schoolYear.CheckPeriod()
	if err != nil {
		return err
	}

	err = s.repository.Create(schoolYear)
	if err != nil {
		log.Println(err)
		return errors.New("failed to create school year")
	}

	return nil
}

func (s *SchoolYearActions) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		log.Println(err)
		return errors.New("failed to delete school year")
	}

	return nil
}

func (s *SchoolYearActions) Update(id string, dto dto.SchoolYearRequestDto) error {
	schoolYear := entities.SchoolYear{}
	schoolYear.FillFromDto(dto)
	schoolYearId, err := uuid.Parse(id)
	if err != nil {
		log.Println(err)
		return errors.New("invalid id provided")
	}
	schoolYear.Id = schoolYearId
	err = schoolYear.CheckPeriod()
	if err != nil {
		return err
	}

	err = s.repository.Update(schoolYear)
	if err != nil {
		log.Println(err)
		return errors.New("failed to update school year")
	}

	return nil
}

func (s *SchoolYearActions) FindOne(id string) (*entities.SchoolYear, error) {
	schoolYear, err := s.repository.FindById(id)
	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to get school year")
	}

	return schoolYear, nil
}

func (s *SchoolYearActions) FindAll(dtoRequest dto.PaginatorRequest) (*common.SchoolYearPaginationResult, error) {
	paginator := common.Pagination{}
	paginator.FillFromDto(dtoRequest)
	paginationResult, err := s.repository.FindAll(paginator)
	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to get school years")
	}

	return paginationResult, nil
}
