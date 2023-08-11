package entities

import (
	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/value_objects"
	"time"
)

type Parent struct {
	Id        uuid.UUID               `json:"id"`
	FirstName string                  `json:"first_name"`
	LastName  string                  `json:"last_name"`
	BirthDay  *time.Time              `json:"birth_day"`
	Addresses []value_objects.Address `json:"addresses"`
	Phones    []value_objects.Phone   `json:"phones"`
}
