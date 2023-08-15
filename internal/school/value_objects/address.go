package value_objects

import "github.com/google/uuid"

type Address struct {
	Id       uuid.UUID `json:"id"`
	Street   string    `json:"street"`
	City     string    `json:"city"`
	District string    `json:"district"`
	State    string    `json:"state"`
	ZipCode  string    `json:"zip_code"`
	OwnerId  uuid.UUID `json:"owner_id"`
}
