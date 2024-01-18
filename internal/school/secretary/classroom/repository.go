package classroom

import "github.com/henriquerocha2004/sistema-escolar/internal/school/shared/paginator"

type Repository interface {
	Create(classRoom ClassRoom) error
	Delete(id string) error
	Update(classRoom ClassRoom) error
	FindAll(pagination paginator.Pagination) (*paginator.PaginationResult, error)
	FindById(id string) (*ClassRoom, error)
	FindByIdLock(id string) (*ClassRoom, error)
}
