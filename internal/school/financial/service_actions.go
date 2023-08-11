package financial

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/common"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities"
)

type ServiceActionsInterface interface {
	Create(dto dto.ServiceRequestDto) error
	Update(id string, dto dto.ServiceRequestDto) error
	Delete(id string) error
	FindById(id string) (*entities.Service, error)
	FindAll(dtoRequest dto.PaginatorRequest) (*common.ServicePaginationResult, error)
}

type ServiceActions struct {
	serviceRepository ServiceRepository
}

func NewServiceActions(repository ServiceRepository) *ServiceActions {
	return &ServiceActions{
		serviceRepository: repository,
	}
}

func (s *ServiceActions) Create(dto dto.ServiceRequestDto) error {

	service := entities.Service{}
	service.FillFromDto(dto)
	service.Id = uuid.New()

	err := s.serviceRepository.Create(service)
	if err != nil {
		log.Println(err)
		return errors.New("failed to create service")
	}

	return nil
}

func (s *ServiceActions) Update(id string, dto dto.ServiceRequestDto) error {

	serviceId, _ := uuid.Parse(id)

	service := entities.Service{}
	service.FillFromDto(dto)
	service.Id = serviceId

	err := s.serviceRepository.Update(service)
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

func (s *ServiceActions) FindById(id string) (*entities.Service, error) {

	service, err := s.serviceRepository.FindById(id)
	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to get service")
	}

	return service, nil
}

func (s *ServiceActions) FindAll(dtoRequest dto.PaginatorRequest) (*common.ServicePaginationResult, error) {
	paginator := common.Pagination{}
	paginator.FillFromDto(dtoRequest)

	paginationResult, err := s.serviceRepository.FindAll(paginator)
	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to get services")
	}

	return paginationResult, nil
}
