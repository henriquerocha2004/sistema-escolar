package secretary

import (
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/financial"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/uow"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/value_objects"
)

type RegistrationActionsInterface interface {
	Create(dto dto.RegistrationDto) (*RegistrationResponse, error)
}

type RegistrationResponse struct {
	RegistrationCode string `json:"registration_code"`
}

type RegistrationActions struct {
	serviceRepo     financial.ServiceRepository
	classRoomRepo   ClassRoomRepository
	registrationUow uow.RegisterUow
}

func NewRegistrationActions(
	serviceRepo financial.ServiceRepository,
	classRoomRepo ClassRoomRepository,
	registerUow uow.RegisterUow,
) *RegistrationActions {
	return &RegistrationActions{
		serviceRepo:     serviceRepo,
		classRoomRepo:   classRoomRepo,
		registrationUow: registerUow,
	}
}

func (r *RegistrationActions) Create(dto dto.RegistrationDto) (*RegistrationResponse, error) {

	service, err := r.serviceRepo.FindById(dto.ServiceId)
	if err != nil || service == nil {
		log.Println(err)
		return nil, errors.New("failed to get service information")
	}

	classRoom, err := r.classRoomRepo.FindById(dto.ClassRoomId)
	if err != nil || classRoom == nil {
		log.Println(err)
		return nil, errors.New("failed to get class room information")
	}

	student := entities.Student{}
	student.FillFromDto(dto.Student)
	student.AddAddress(dto.Student.Addresses)
	student.AddPhones(dto.Student.Phones)
	student.AddParent(dto.Student.Parents)

	r.registrationUow.BeginTransaction()

	studentId, err := r.registrationUow.StudentAlreadyExists(string(student.CPFDocument))

	if err != nil {
		r.registrationUow.Rollback()
		log.Println(err)
		return nil, errors.New("failed to verify if student already exists")
	}

	if studentId == nil {
		err = r.registrationUow.CreateStudent(student)
		if err != nil {
			r.registrationUow.Rollback()
			log.Println(err)
			return nil, errors.New("failed to save student")
		}

	} else {

		studentRegistered, err := r.registrationUow.StudentAlreadyRegisterInClass(*studentId, classRoom.Id)
		if err != nil {
			r.registrationUow.Rollback()
			log.Println(err)
			return nil, errors.New( "failed to verify if student already registered")
		}

		if studentRegistered {
			r.registrationUow.Rollback()
			log.Println(err)
			return nil, errors.New("student already registered")
		}
	}

	enrollmentDate := time.Now()
	enrollmentDueDate, _ := time.Parse("2006-02-02", dto.EnrollmentDueDate)

	registration := entities.Registration{
		Service:              *service,
		Class:                *classRoom,
		Shift:                value_objects.Shift(dto.Shift),
		Student:              student,
		MonthDuration:        dto.MonthDuration,
		EnrollmentFee:        dto.EnrollmentFee,
		InstallmentsQuantity: dto.InstallmentsQuantity,
		PaymentDay:           dto.PaymentDay,
		MonthlyFee:           dto.MonthlyFee,
		EnrollmentDate:       &enrollmentDate,
		EnrollmentDueDate:    &enrollmentDueDate,
		Id:                   uuid.New(),
	}
	registration.GenerateCode()

	err = registration.Check()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = r.registrationUow.CreateRegister(registration)
	if err != nil {
		r.registrationUow.Rollback()
		return nil, err
	}

	r.registrationUow.Commit()

	registrationResponse := RegistrationResponse{
		RegistrationCode: registration.Code,
	}

	return &registrationResponse, nil
}
