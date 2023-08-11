package entities

import (
	"github.com/google/uuid"
	"time"
)

type Price struct {
	Id          uuid.UUID
	Description string
	Value       float64
	ValidAt     *time.Time
	ServiceId   uuid.UUID
}
