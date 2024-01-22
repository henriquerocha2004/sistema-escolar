package registration

import "github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/student"

type RequestDto struct {
	ClassRoomId          string             `json:"class_room_id"`
	Shift                string             `json:"shift"`
	Student              student.RequestDto `json:"student"`
	ServiceId            string             `json:"service_id"`
	MonthlyFee           float64            `json:"monthly_fee"`
	InstallmentsQuantity int                `json:"installments_quantity"`
	EnrollmentFee        float64            `json:"enrollment_due_date"`
	EnrollmentDueDate    string             `json:"due_date"`
	MonthDuration        int                `json:"month_duration"`
	PaymentDay           string             `json:"payment_day"`
}
