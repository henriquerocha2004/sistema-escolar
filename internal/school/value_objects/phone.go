package value_objects

import "github.com/google/uuid"

type Phone struct {
	Id          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Phone       string    `json:"phone"`
	OwnerId     uuid.UUID `json:"owner_id"`
}
