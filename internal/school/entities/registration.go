package entities

import (
	"time"

	"github.com/google/uuid"
)

type Registration struct {
	Id                   uuid.UUID
	Code                 string
	Class                ClassRoom
	Shift                string
	Student              Student
	Service              Price
	MonthlyFee           float64
	InstallmentsQuantity int
	EnrollmentFee        float64
	DueDate              *time.Time
	MonthDuration        int
	Status               string
}
