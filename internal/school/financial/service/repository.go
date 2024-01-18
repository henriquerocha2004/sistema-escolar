package service

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/paginator"
)

type Repository interface {
	Create(service Service) error
	Delete(id string) error
	Update(service Service) error
	FindById(id string) (*Service, error)
	FindAll(paginator paginator.Pagination) (*paginator.ServicePaginationResult, error)
}
