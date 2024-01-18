package phone

import "github.com/henriquerocha2004/sistema-escolar/internal/school/value_objects"

type Repository interface {
	Create(phone value_objects.Phone) error
}
