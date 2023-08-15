package entities

import (
	"time"

	"github.com/google/uuid"
)

type Registration struct {
	Id                   uuid.UUID  `json:"id"`
	Code                 string     `json:"code"`
	Class                ClassRoom  `json:"class"`
	Shift                string     `json:"shift"`
	Student              Student    `json:"student"`
	Service              Service    `json:"service"`
	MonthlyFee           float64    `json:"monthly_fee"`
	InstallmentsQuantity int        `json:"installments_quantity"`
	EnrollmentFee        float64    `json:"enrollment_fee"`
	DueDate              *time.Time `json:"due_date"`
	MonthDuration        int        `json:"month_duration"`
	Status               string     `json:"status"`
	EnrollmentDate       *time.Time `json:"enrollment_date"`
}
