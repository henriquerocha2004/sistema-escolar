package entities

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/value_objects"
	"github.com/stretchr/testify/assert"
)

func TestForValidateStudentInRegistration(t *testing.T) {
	inputData := createInputData()
	service := getService()
	classRoom := getClassRoom()

	enrollmentDate := time.Now()
	enrollmentDueDate, _ := time.Parse("2006-02-02", inputData.EnrollmentDueDate)

	t.Run("should return error when student is not him self responsible but not provided parents information", func(t *testing.T) {
		student := getStudent(inputData)
		student.Parents = []Parent{}
		student.HimSelfResponsible = false

		registration := Registration{
			Class:                classRoom,
			Service:              service,
			Student:              student,
			MonthlyFee:           inputData.MonthlyFee,
			InstallmentsQuantity: inputData.InstallmentsQuantity,
			EnrollmentFee:        inputData.EnrollmentFee,
			Shift:                value_objects.Shift(inputData.Shift),
			EnrollmentDueDate:    &enrollmentDueDate,
			EnrollmentDate:       &enrollmentDate,
			MonthDuration:        inputData.MonthDuration,
			PaymentDay:           inputData.PaymentDay,
			Paid:                 false,
			Id:                   uuid.New(),
		}

		err := registration.Check()
		assert.Error(t, err)
		assert.Equal(t, "information about student parents not found", err.Error())

	})

	t.Run("should return error if cpf document is invalid", func(t *testing.T) {
		student := getStudent(inputData)
		student.CPFDocument = "145885566998"
		registration := Registration{
			Class:                classRoom,
			Service:              service,
			Student:              student,
			MonthlyFee:           inputData.MonthlyFee,
			InstallmentsQuantity: inputData.InstallmentsQuantity,
			EnrollmentFee:        inputData.EnrollmentFee,
			Shift:                value_objects.Shift(inputData.Shift),
			EnrollmentDueDate:    &enrollmentDueDate,
			EnrollmentDate:       &enrollmentDate,
			MonthDuration:        inputData.MonthDuration,
			PaymentDay:           inputData.PaymentDay,
			Paid:                 false,
			Id:                   uuid.New(),
		}

		err := registration.Check()
		assert.Error(t, err)
		assert.Equal(t, "invalid cpf", err.Error())
	})

	t.Run("should return error if student is responsible himself but not provided phone information", func(t *testing.T) {
		student := getStudent(inputData)
		student.Phones = []value_objects.Phone{}

		registration := Registration{
			Class:                classRoom,
			Service:              service,
			Student:              student,
			MonthlyFee:           inputData.MonthlyFee,
			InstallmentsQuantity: inputData.InstallmentsQuantity,
			EnrollmentFee:        inputData.EnrollmentFee,
			Shift:                value_objects.Shift(inputData.Shift),
			EnrollmentDueDate:    &enrollmentDueDate,
			EnrollmentDate:       &enrollmentDate,
			MonthDuration:        inputData.MonthDuration,
			PaymentDay:           inputData.PaymentDay,
			Paid:                 false,
			Id:                   uuid.New(),
		}

		err := registration.Check()
		assert.Error(t, err)
		assert.Equal(t, "phone information not found", err.Error())
	})

	t.Run("should return error if student is responsible himself but not provided address information", func(t *testing.T) {
		student := getStudent(inputData)
		student.Addresses = []value_objects.Address{}

		registration := Registration{
			Class:                classRoom,
			Service:              service,
			Student:              student,
			MonthlyFee:           inputData.MonthlyFee,
			InstallmentsQuantity: inputData.InstallmentsQuantity,
			EnrollmentFee:        inputData.EnrollmentFee,
			Shift:                value_objects.Shift(inputData.Shift),
			EnrollmentDueDate:    &enrollmentDueDate,
			EnrollmentDate:       &enrollmentDate,
			MonthDuration:        inputData.MonthDuration,
			PaymentDay:           inputData.PaymentDay,
			Paid:                 false,
			Id:                   uuid.New(),
		}

		err := registration.Check()
		assert.Error(t, err)
		assert.Equal(t, "address information not found", err.Error())
	})

	t.Run("should return error if student is not himself responsible but not provided parents address information", func(t *testing.T) {
		student := getStudent(inputData)
		student.Parents[0].Addresses = []value_objects.Address{}
		student.HimSelfResponsible = false
		registration := Registration{
			Class:                classRoom,
			Service:              service,
			Student:              student,
			MonthlyFee:           inputData.MonthlyFee,
			InstallmentsQuantity: inputData.InstallmentsQuantity,
			EnrollmentFee:        inputData.EnrollmentFee,
			Shift:                value_objects.Shift(inputData.Shift),
			EnrollmentDueDate:    &enrollmentDueDate,
			EnrollmentDate:       &enrollmentDate,
			MonthDuration:        inputData.MonthDuration,
			PaymentDay:           inputData.PaymentDay,
			Paid:                 false,
			Id:                   uuid.New(),
		}

		err := registration.Check()
		assert.Error(t, err)
		assert.Equal(t, "parent address information not found", err.Error())

	})

	t.Run("should return error if student is not himself responsible but not provided phone parents information", func(t *testing.T) {
		student := getStudent(inputData)
		student.Parents[0].Phones = []value_objects.Phone{}
		student.HimSelfResponsible = false
		registration := Registration{
			Class:                classRoom,
			Service:              service,
			Student:              student,
			MonthlyFee:           inputData.MonthlyFee,
			InstallmentsQuantity: inputData.InstallmentsQuantity,
			EnrollmentFee:        inputData.EnrollmentFee,
			Shift:                value_objects.Shift(inputData.Shift),
			EnrollmentDueDate:    &enrollmentDueDate,
			EnrollmentDate:       &enrollmentDate,
			MonthDuration:        inputData.MonthDuration,
			PaymentDay:           inputData.PaymentDay,
			Paid:                 false,
			Id:                   uuid.New(),
		}

		err := registration.Check()
		assert.Error(t, err)
		assert.Equal(t, "parent phone not found", err.Error())
	})
}

func TestRegistrationValidation(t *testing.T) {
	inputData := createInputData()
	service := getService()
	classRoom := getClassRoom()

	enrollmentDate := time.Now()
	enrollmentDueDate, _ := time.Parse("2006-02-02", "2023-01-01")
	student := getStudent(inputData)

	t.Run("should return error if shift informed is invalid", func(t *testing.T) {
		registration := Registration{
			Class:                classRoom,
			Service:              service,
			Student:              student,
			MonthlyFee:           inputData.MonthlyFee,
			InstallmentsQuantity: inputData.InstallmentsQuantity,
			EnrollmentFee:        inputData.EnrollmentFee,
			Shift:                value_objects.Shift("INVALID"),
			EnrollmentDueDate:    &enrollmentDueDate,
			EnrollmentDate:       &enrollmentDate,
			MonthDuration:        inputData.MonthDuration,
			PaymentDay:           inputData.PaymentDay,
			Paid:                 false,
			Id:                   uuid.New(),
		}

		err := registration.Check()
		assert.Error(t, err)
		assert.Equal(t, "invalid shift provided", err.Error())
	})

	t.Run("should return error if enrollment due date is after than current date", func(t *testing.T) {

		enrollmentDueDateBefore := time.Now().AddDate(0, 0, -2)

		registration := Registration{
			Class:                classRoom,
			Service:              service,
			Student:              student,
			MonthlyFee:           inputData.MonthlyFee,
			InstallmentsQuantity: inputData.InstallmentsQuantity,
			EnrollmentFee:        inputData.EnrollmentFee,
			Shift:                value_objects.Shift(inputData.Shift),
			EnrollmentDueDate:    &enrollmentDueDateBefore,
			EnrollmentDate:       &enrollmentDate,
			MonthDuration:        inputData.MonthDuration,
			PaymentDay:           inputData.PaymentDay,
			Paid:                 false,
			Id:                   uuid.New(),
		}

		err := registration.Check()
		assert.Error(t, err)
		assert.Equal(t, "enrollment due date cant be before today", err.Error())
	})

	t.Run("should return error if installments total is less than total service", func(t *testing.T) {
		service.Price = 400.00
		enrollmentDueDate := time.Now().AddDate(0, 0, 2)
		registration := Registration{
			Class:                classRoom,
			Service:              service,
			Student:              student,
			MonthlyFee:           200.00,
			InstallmentsQuantity: 1,
			EnrollmentFee:        inputData.EnrollmentFee,
			Shift:                value_objects.Shift(inputData.Shift),
			EnrollmentDueDate:    &enrollmentDueDate,
			EnrollmentDate:       &enrollmentDate,
			MonthDuration:        inputData.MonthDuration,
			PaymentDay:           inputData.PaymentDay,
			Paid:                 false,
			Id:                   uuid.New(),
		}

		err := registration.Check()
		assert.Error(t, err)
		assert.Equal(t, "total paid not match with service price", err.Error())
	})
}

func getService() Service {
	return Service{
		Id:          uuid.New(),
		Description: "Ensino Fundamental",
		Price:       5000.00,
	}
}

func getClassRoom() ClassRoom {

	now := time.Now()

	return ClassRoom{
		Id:              uuid.New(),
		VacancyQuantity: 10,
		Shift:           "morning",
		OpenDate:        &now,
		OccupiedVacancy: 3,
		Status:          "OPEN",
		Identification:  "TUR-001",
		SchoolYearId:    uuid.New(),
		Level:           "Jardim",
		RoomId: uuid.NullUUID{
			UUID:  uuid.New(),
			Valid: true,
		},
		ScheduleId:   uuid.New(),
		Localization: "ANY",
		Type:         "remote",
	}
}

func getStudent(inputData dto.RegistrationDto) Student {
	student := Student{}
	student.FillFromDto(inputData.Student)
	student.AddPhones(inputData.Student.Phones)
	student.AddAddress(inputData.Student.Addresses)
	student.AddParent(inputData.Student.Parents)

	return student
}

func createInputData() dto.RegistrationDto {

	addresses := []dto.AddressDto{
		{
			Street:   "Rua dos Bobos",
			City:     "Marbule",
			District: "El Nildo",
			State:    "CC",
			ZipCode:  "47854-552",
		},
	}

	phones := []dto.PhoneDto{
		{
			Description: "Pessoal",
			Phone:       "7199542-3264",
		},
	}

	parents := []dto.ParentDto{
		{
			FirstName:   "Marcha",
			LastName:    "Marbule",
			BirthDay:    "1970-05-05",
			Addresses:   addresses,
			Phones:      phones,
			RgDocument:  "745698885",
			CpfDocument: "624.499.720-48",
		},
	}

	student := dto.StudentDto{
		FirstName:          "Henrique",
		LastName:           "Rocha",
		Birthday:           "1987-09-21",
		RgDocument:         "1452658877",
		CpfDocument:        "823.781.140-28",
		Email:              "test@test.com",
		HimSelfResponsible: true,
		Addresses:          addresses,
		Phones:             phones,
		Parents:            parents,
	}

	registration := dto.RegistrationDto{
		ClassRoomId:          "76d0f9b7-3d7d-4cfc-892a-94c0704b4deb",
		Shift:                "morning",
		Student:              student,
		ServiceId:            "a1c003f8-8a56-4a49-892a-825b364cc076",
		MonthlyFee:           400.00,
		InstallmentsQuantity: 12,
		EnrollmentFee:        60.00,
		EnrollmentDueDate:    "2023-08-30",
		MonthDuration:        12,
		PaymentDay:           "16",
	}

	return registration
}
