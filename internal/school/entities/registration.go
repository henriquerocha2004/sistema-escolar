package entities

import (
	"errors"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/value_objects"
)

type Registration struct {
	Id                   uuid.UUID           `json:"id"`
	Code                 string              `json:"code"`
	Class                ClassRoom           `json:"class"`
	Shift                value_objects.Shift `json:"shift"`
	Student              Student             `json:"student"`
	Service              Service             `json:"service"`
	MonthlyFee           float64             `json:"monthly_fee"`
	InstallmentsQuantity int                 `json:"installments_quantity"`
	EnrollmentFee        float64             `json:"enrollment_fee"`
	EnrollmentDueDate    *time.Time          `json:"enrollment_due_date"`
	MonthDuration        int                 `json:"month_duration"`
	Status               string              `json:"status"`
	EnrollmentDate       *time.Time          `json:"enrollment_date"`
	PaymentDay           string              `json:"payment_day"`
	Paid                 bool                `json:"paid"`
}

func (r *Registration) GenerateCode() {
	currentDate := time.Now().Format("20060202")
	randomNumber := rand.Intn(100000)
	code := strconv.Itoa(randomNumber)
	r.Code = currentDate + code
}

func (r *Registration) Check() error {
	log.Println(r.Student)
	err := r.Student.Validate()
	if err != nil {
		return err
	}

	err = r.Shift.Validate()
	if err != nil {
		return err
	}

	err = r.checkEnrollment()
	if err != nil {
		return err
	}

	err = r.checkPayment()
	if err != nil {
		return err
	}

	r.SetStatus()

	return nil
}

func (r *Registration) checkEnrollment() error {

	if r.EnrollmentFee == 0 {
		return nil
	}

	if r.EnrollmentDueDate.Before(time.Now()) {
		return errors.New("enrollment due date cant be before today")
	}

	return nil
}

func (r *Registration) checkPayment() error {

	total := r.MonthlyFee * float64(r.InstallmentsQuantity)

	if total < r.Service.Price {
		return errors.New("total paid not match with service price")
	}

	return nil
}

func (r *Registration) SetStatus() {

	if r.EnrollmentFee > 0 && !r.Paid {
		r.Status = "WAIT_ENROLLMENT_FEE"
	}

	r.Status = "APPROVED"
}
