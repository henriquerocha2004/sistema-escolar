package room

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/paginator"
)

type Repository interface {
	Create(room Room) error
	Delete(id string) error
	Update(room Room) error
	FindByCode(code string) (*Room, error)
	FindAll(pagination paginator.Pagination) (*paginator.PaginationResult, error)
	FindById(id string) (*Room, error)
}
