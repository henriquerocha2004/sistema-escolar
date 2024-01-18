package schoolYearService

import (
	"errors"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/schoolyear"

	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/paginator"
	"log"
)

type SchoolYearActionsInterface interface {
	Create(dto schoolyear.Request) error
	Delete(id string) error
	Update(id string, dto schoolyear.Request) error
	FindOne(id string) (*schoolyear.SchoolYear, error)
	FindAll(dtoRequest paginator.PaginatorRequest) (*paginator.PaginationResult, error)
}

type SchoolYearActions struct {
	repository schoolyear.Repository
}

func New(repository schoolyear.Repository) *SchoolYearActions {
	return &SchoolYearActions{
		repository: repository,
	}
}

func (s *SchoolYearActions) Create(dto schoolyear.Request) error {
	schoolYear, err := schoolyear.New(dto.Year, dto.StartedAt, dto.EndAt)
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

func (s *SchoolYearActions) Update(id string, dto schoolyear.Request) error {
	schoolYear, err := schoolyear.New(dto.Year, dto.StartedAt, dto.EndAt)
	if err != nil {
		log.Println(err)
		return errors.New("failed to update school year")
	}

	err = schoolYear.SetId(id)
	if err != nil {
		log.Println(err)
		return errors.New("failed to update school year")
	}

	err = s.repository.Update(schoolYear)
	if err != nil {
		log.Println(err)
		return errors.New("failed to update school year")
	}

	return nil
}

func (s *SchoolYearActions) FindOne(id string) (*schoolyear.SchoolYear, error) {
	schoolYear, err := s.repository.FindById(id)
	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to get school year")
	}

	return schoolYear, nil
}

func (s *SchoolYearActions) FindAll(dtoRequest paginator.PaginatorRequest) (*paginator.PaginationResult, error) {
	pg := paginator.Pagination{}
	pg.FillFromDto(dtoRequest)
	paginationResult, err := s.repository.FindAll(pg)
	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to get school years")
	}

	return paginationResult, nil
}
