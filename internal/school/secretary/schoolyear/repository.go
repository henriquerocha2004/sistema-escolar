package schoolyear

import "github.com/henriquerocha2004/sistema-escolar/internal/school/shared/paginator"

type Repository interface {
	Create(schoolYear *SchoolYear) error
	Delete(id string) error
	Update(schoolYear *SchoolYear) error
	FindById(id string) (*SchoolYear, error)
	FindByYear(year string) (*SchoolYear, error)
	FindAll(paginator paginator.Pagination) (*paginator.PaginationResult, error)
}
