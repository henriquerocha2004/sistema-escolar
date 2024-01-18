package address

import "github.com/henriquerocha2004/sistema-escolar/internal/school/value_objects"

type Repository interface {
	Create(address value_objects.Address) error
}
