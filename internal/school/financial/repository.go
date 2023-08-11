package financial

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/common"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities"
)

type ServiceRepository interface {
	Create(service entities.Service) error
	Delete(id string) error
	Update(service entities.Service) error
	FindById(id string) (*entities.Service, error)
	FindAll(paginator common.Pagination) (*common.ServicePaginationResult, error)
}
