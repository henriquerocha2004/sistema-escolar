package serviceActions

import (
	"errors"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/financial/service"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/paginator"

	"log"
)

type ActionsServiceInterface interface {
	Create(dto service.Request) error
	Update(id string, dto service.Request) error
	Delete(id string) error
	FindById(id string) (*service.Service, error)
	FindAll(dtoRequest paginator.PaginatorRequest) (*paginator.PaginationResult, error)
}

type ServiceActions struct {
	serviceRepository service.Repository
}

func New(repository service.Repository) *ServiceActions {
	return &ServiceActions{
		serviceRepository: repository,
	}
}

func (s *ServiceActions) Create(dto service.Request) error {
	serv, err := service.New(dto.Description, dto.Value)

	err = s.serviceRepository.Create(*serv)
	if err != nil {
		log.Println(err)
		return errors.New("failed to create service")
	}

	return nil
}

func (s *ServiceActions) Update(id string, dto service.Request) error {

	serv, err := service.New(dto.Description, dto.Value)
	if err != nil {
		return err
	}

	err = serv.ChangeId(id)
	if err != nil {
		return err
	}

	err = s.serviceRepository.Update(*serv)
	if err != nil {
		log.Println(err)
		return errors.New("failed to update service")
	}

	return nil
}

func (s *ServiceActions) Delete(id string) error {
	err := s.serviceRepository.Delete(id)

	if err != nil {
		log.Println(err)
		return errors.New("failed to delete service")
	}

	return nil
}

func (s *ServiceActions) FindById(id string) (*service.Service, error) {

	serv, err := s.serviceRepository.FindById(id)
	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to get service")
	}

	return serv, nil
}

func (s *ServiceActions) FindAll(dtoRequest paginator.PaginatorRequest) (*paginator.PaginationResult, error) {
	pg := paginator.Pagination{}
	pg.FillFromDto(dtoRequest)

	paginationResult, err := s.serviceRepository.FindAll(pg)
	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to get services")
	}

	return paginationResult, nil
}
