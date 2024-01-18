package registration

import (
	"encoding/json"
	"errors"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/financial/service"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/classroom"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/student"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/value_objects"
)

type Registration struct {
	id                   uuid.UUID
	code                 string
	class                classroom.ClassRoom
	shift                value_objects.Shift
	student              student.Student
	service              service.Service
	monthlyFee           float64
	installmentsQuantity int
	enrollmentFee        float64
	enrollmentDueDate    time.Time
	monthDuration        int
	status               string
	enrollmentDate       time.Time
	paymentDay           string
	paid                 bool
}

func New(class classroom.ClassRoom,
	shift string,
	student student.Student,
	service service.Service,
	monthlyFee float64,
	installmentsQuantity int,
	enrollmentFee float64,
	enrollmentDueDate string,
	monthDuration int,
	paymentDay string) (*Registration, error) {

	r := &Registration{
		id:             uuid.New(),
		student:        student,
		service:        service,
		class:          class,
		enrollmentDate: time.Now(),
	}

	r.GenerateCode()
	r.ChangeEnrollmentFee(enrollmentFee)

	err := r.ChangeShift(shift)
	if err != nil {
		return nil, err
	}

	err = r.ChangeMonthlyFee(monthlyFee)
	if err != nil {
		return nil, err
	}

	err = r.ChangeInstallmentsQuantity(installmentsQuantity)
	if err != nil {
		return nil, err
	}

	err = r.ChangeEnrollmentDueDate(enrollmentDueDate)
	if err != nil {
		return nil, err
	}

	err = r.ChangeMonthDuration(monthDuration)
	if err != nil {
		return nil, err
	}

	err = r.ChangePaymentDay(paymentDay)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r *Registration) Id() uuid.UUID {
	return r.id
}

func (r *Registration) Code() string {
	return r.code
}

func (r *Registration) Class() *classroom.ClassRoom {
	return &r.class
}

func (r *Registration) Shift() value_objects.Shift {
	return r.shift
}

func (r *Registration) Student() *student.Student {
	return &r.student
}

func (r *Registration) Service() *service.Service {
	return &r.service
}

func (r *Registration) MonthlyFee() float64 {
	return r.monthlyFee
}

func (r *Registration) InstallmentsQuantity() int {
	return r.installmentsQuantity
}

func (r *Registration) EnrollmentFee() float64 {
	return r.enrollmentFee
}

func (r *Registration) EnrollmentDate() time.Time {
	return r.enrollmentDate
}

func (r *Registration) EnrollmentDueDate() time.Time {
	return r.enrollmentDueDate
}

func (r *Registration) MonthDuration() int {
	return r.monthDuration
}

func (r *Registration) PaymentDay() string {
	return r.paymentDay
}

func (r *Registration) Status() string {
	return r.status
}

func (r *Registration) Paid() bool {
	return r.paid
}

func (r *Registration) ChangeId(id string) error {
	if id == "" {
		return errors.New("registration id cannot be empty")
	}

	regId, err := uuid.Parse(id)
	if err != nil {
		log.Println(err)
		return errors.New("failed to change registration id")
	}

	r.id = regId

	return nil
}

func (r *Registration) ChangePaymentDay(day string) error {
	if day == "" {
		return errors.New("payment day cannot be empty")
	}

	r.paymentDay = day

	return nil
}

func (r *Registration) ChangeEnrollmentDate(date string) error {
	if date == "" {
		return errors.New("enrollment date cannot be empty")
	}

	d, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Println(err)
		return errors.New("failed to change enrollment date")
	}

	r.enrollmentDate = d

	return nil
}

func (r *Registration) ChangeMonthDuration(duration int) error {
	if duration <= 0 {
		return errors.New("month duration cannot be empty")
	}

	r.monthDuration = duration

	return nil
}

func (r *Registration) ChangeEnrollmentDueDate(date string) error {
	if date == "" {
		return nil
	}

	d, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Println(err)
		return errors.New("failed to change enrollment due date")
	}

	r.enrollmentDueDate = d

	return nil
}

func (r *Registration) ChangeEnrollmentFee(fee float64) {
	if fee <= 0 {
		return
	}

	r.enrollmentFee = fee
}

func (r *Registration) ChangeInstallmentsQuantity(installments int) error {
	if installments <= 0 {
		return errors.New("installments quantity cannot be empty")
	}

	r.installmentsQuantity = installments

	return nil
}

func (r *Registration) ChangeMonthlyFee(fee float64) error {
	if fee == 0 {
		return errors.New("monthly fee cannot be empty")
	}

	r.monthlyFee = fee

	return nil
}

func (r *Registration) ChangeShift(shift string) error {
	if shift == "" {
		return errors.New("shift cannot be null")
	}

	s := value_objects.Shift(shift)
	err := s.Validate()
	if err != nil {
		return err
	}

	r.shift = s

	return nil
}

func (r *Registration) GenerateCode() {
	currentDate := time.Now().Format("20060202")
	randomNumber := rand.Intn(100000)
	code := strconv.Itoa(randomNumber)
	r.code = currentDate + code
}

func (r *Registration) Check() error {
	err := r.student.Validate()
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

	r.ChangeStatus()

	return nil
}

func (r *Registration) checkEnrollment() error {

	if r.enrollmentFee == 0 {
		return nil
	}

	if r.enrollmentDueDate.Before(time.Now()) {
		return errors.New("enrollment due date cant be before today")
	}

	return nil
}

func (r *Registration) checkPayment() error {

	total := r.monthlyFee * float64(r.installmentsQuantity)

	if total < r.service.Price() {
		return errors.New("total paid not match with service price")
	}

	return nil
}

func (r *Registration) ChangeStatus() {

	if r.enrollmentFee > 0 && !r.paid {
		r.status = "WAIT_ENROLLMENT_FEE"
	}

	r.status = "APPROVED"
}

func (r *Registration) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id                   string              `json:"id"`
		Code                 string              `json:"code"`
		Class                classroom.ClassRoom `json:"class"`
		Shift                string              `json:"shift"`
		Student              student.Student     `json:"student"`
		Service              service.Service     `json:"service"`
		MonthlyFee           float64             `json:"monthly_fee"`
		InstallmentsQuantity int                 `json:"installments_quantity"`
		EnrollmentFee        float64             `json:"enrollment_fee"`
		EnrollmentDueDate    string              `json:"enrollment_due_date"`
		MonthDuration        int                 `json:"month_duration"`
		Status               string              `json:"status"`
		EnrollmentDate       string              `json:"enrollment_date"`
		PaymentDay           string              `json:"payment_day"`
		Paid                 bool                `json:"paid"`
	}{
		Id:                   r.Id().String(),
		Code:                 r.Code(),
		Class:                *r.Class(),
		Shift:                string(r.Shift()),
		Student:              *r.Student(),
		Service:              *r.Service(),
		MonthlyFee:           r.MonthlyFee(),
		InstallmentsQuantity: r.InstallmentsQuantity(),
		EnrollmentFee:        r.EnrollmentFee(),
		EnrollmentDueDate:    r.EnrollmentDueDate().Format("2006-01-02"),
		MonthDuration:        r.MonthDuration(),
		Status:               r.Status(),
		EnrollmentDate:       r.EnrollmentDate().Format("2006-01-02"),
		PaymentDay:           r.PaymentDay(),
		Paid:                 r.Paid(),
	})
}
