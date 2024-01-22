package registrationService

import (
	"errors"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/financial/service"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/classroom"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/registration"

	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/student"
	"log"
)

type RegistrationActionsInterface interface {
	Create(dto registration.RequestDto) (*RegistrationResponse, error)
}

type RegistrationResponse struct {
	RegistrationCode string `json:"registration_code"`
}

type RegistrationActions struct {
	serviceRepo     service.Repository
	classRoomRepo   classroom.Repository
	registrationUow registration.RegisterUow
}

func NewRegistrationActions(
	serviceRepo service.Repository,
	classRoomRepo classroom.Repository,
	registerUow registration.RegisterUow,
) *RegistrationActions {
	return &RegistrationActions{
		serviceRepo:     serviceRepo,
		classRoomRepo:   classRoomRepo,
		registrationUow: registerUow,
	}
}

func (r *RegistrationActions) Create(dto registration.RequestDto) (*RegistrationResponse, error) {

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

	student, err := student.New(
		dto.Student.FirstName,
		dto.Student.LastName,
		dto.Student.Birthday,
		dto.Student.RgDocument,
		dto.Student.CpfDocument,
		dto.Student.Email,
		dto.Student.HimSelfResponsible,
	)

	if err != nil {
		return nil, err
	}

	student.AddAddress(dto.Student.Addresses)
	student.AddPhones(dto.Student.Phones)

	err = student.AddParents(dto.Student.Parents)
	if err != nil {
		return nil, err
	}

	err = r.registrationUow.BeginTransaction()
	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to create registration")
	}

	studentId, err := r.registrationUow.StudentAlreadyExists(string(student.Cpf()))

	if err != nil {
		_ = r.registrationUow.Rollback()
		log.Println(err)
		return nil, errors.New("failed to verify if student already exists")
	}

	if studentId == nil {
		err = r.registrationUow.CreateStudent(*student)
		if err != nil {
			_ = r.registrationUow.Rollback()
			log.Println(err)
			return nil, errors.New("failed to save student")
		}

	} else {

		studentRegistered, err := r.registrationUow.StudentAlreadyRegisterInClass(*studentId, classRoom.Id())
		if err != nil {
			_ = r.registrationUow.Rollback()
			log.Println(err)
			return nil, errors.New("failed to verify if student already registered")
		}

		if studentRegistered {
			_ = r.registrationUow.Rollback()
			log.Println(err)
			return nil, errors.New("student already registered")
		}
	}

	reg, err := registration.New(
		*classRoom,
		dto.Shift,
		*student,
		*service,
		dto.MonthlyFee,
		dto.InstallmentsQuantity,
		dto.EnrollmentFee,
		dto.EnrollmentDueDate,
		dto.MonthDuration,
		dto.PaymentDay,
	)

	err = reg.Check()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = r.registrationUow.CreateRegister(*reg)
	if err != nil {
		_ = r.registrationUow.Rollback()
		return nil, err
	}

	_ = r.registrationUow.Commit()

	registrationResponse := RegistrationResponse{
		RegistrationCode: reg.Code(),
	}

	return &registrationResponse, nil
}
