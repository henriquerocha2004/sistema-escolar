package registration

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/financial/service"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/classroom"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/parent"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/student"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/address"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/phone"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestForValidateStudentInRegistration(t *testing.T) {
	inputData := createInputData()
	service := getService()
	classRoom := getClassRoom()

	t.Run("should return error when student is not him self responsible but not provided parents information", func(t *testing.T) {
		data := inputData
		data.Student.Parents = []parent.RequestDto{}
		stdt := getStudent(data)
		stdt.ChangeHimSelfResponsible(false)

		reg, err := New(
			classRoom,
			inputData.Shift,
			stdt,
			service,
			inputData.MonthlyFee,
			inputData.InstallmentsQuantity,
			inputData.EnrollmentFee,
			inputData.EnrollmentDueDate,
			inputData.MonthDuration,
			inputData.PaymentDay,
		)

		err = reg.Check()
		assert.Error(t, err)
		assert.Equal(t, "information about student parents not found", err.Error())

	})

	t.Run("should return error if student is responsible himself but not provided phone information", func(t *testing.T) {
		data := inputData
		data.Student.Phones = []phone.RequestDto{}
		stdt := getStudent(data)
		reg, err := New(
			classRoom,
			inputData.Shift,
			stdt,
			service,
			inputData.MonthlyFee,
			inputData.InstallmentsQuantity,
			inputData.EnrollmentFee,
			inputData.EnrollmentDueDate,
			inputData.MonthDuration,
			inputData.PaymentDay,
		)

		err = reg.Check()
		assert.Error(t, err)
		assert.Equal(t, "phone information not found", err.Error())
	})

	t.Run("should return error if student is responsible himself but not provided address information", func(t *testing.T) {
		data := createInputData()
		data.Student.Addresses = []address.RequestDto{}
		student := getStudent(data)
		reg, err := New(
			classRoom,
			inputData.Shift,
			student,
			service,
			inputData.MonthlyFee,
			inputData.InstallmentsQuantity,
			inputData.EnrollmentFee,
			inputData.EnrollmentDueDate,
			inputData.MonthDuration,
			inputData.PaymentDay,
		)

		err = reg.Check()
		assert.Error(t, err)
		assert.Equal(t, "address information not found", err.Error())
	})

	t.Run("should return error if student is not himself responsible but not provided parents address information", func(t *testing.T) {
		data := createInputData()
		data.Student.Parents[0].Addresses = []address.RequestDto{}
		student := getStudent(data)
		student.ChangeHimSelfResponsible(false)
		err := student.AddParents(data.Student.Parents)
		reg, err := New(
			classRoom,
			inputData.Shift,
			student,
			service,
			inputData.MonthlyFee,
			inputData.InstallmentsQuantity,
			inputData.EnrollmentFee,
			inputData.EnrollmentDueDate,
			inputData.MonthDuration,
			inputData.PaymentDay,
		)

		err = reg.Check()
		assert.Error(t, err)
		assert.Equal(t, "parent address information not found", err.Error())

	})

	t.Run("should return error if student is not himself responsible but not provided phone parents information", func(t *testing.T) {
		data := inputData
		data.Student.Parents[0].Phones = []phone.RequestDto{}
		student := getStudent(data)
		student.ChangeHimSelfResponsible(false)
		reg, err := New(
			classRoom,
			inputData.Shift,
			student,
			service,
			inputData.MonthlyFee,
			inputData.InstallmentsQuantity,
			inputData.EnrollmentFee,
			inputData.EnrollmentDueDate,
			inputData.MonthDuration,
			inputData.PaymentDay,
		)

		err = reg.Check()
		assert.Error(t, err)
		assert.Equal(t, "parent phone not found", err.Error())
	})
}

func TestRegistrationValidation(t *testing.T) {
	inputData := createInputData()
	service := getService()
	classRoom := getClassRoom()
	student := getStudent(inputData)

	t.Run("should return error if shift informed is invalid", func(t *testing.T) {
		reg, err := New(
			classRoom,
			"Invalid",
			student,
			service,
			inputData.MonthlyFee,
			inputData.InstallmentsQuantity,
			inputData.EnrollmentFee,
			inputData.EnrollmentDueDate,
			inputData.MonthDuration,
			inputData.PaymentDay,
		)
		assert.Nil(t, reg)
		assert.Error(t, err)
		assert.Equal(t, "invalid shift provided", err.Error())
	})

	t.Run("should return error if enrollment due date is after than current date", func(t *testing.T) {

		enrollmentDueDateBefore := time.Now().AddDate(0, 0, -2).Format("2006-01-02")

		reg, err := New(
			classRoom,
			inputData.Shift,
			student,
			service,
			inputData.MonthlyFee,
			inputData.InstallmentsQuantity,
			inputData.EnrollmentFee,
			enrollmentDueDateBefore,
			inputData.MonthDuration,
			inputData.PaymentDay,
		)

		err = reg.Check()

		assert.Error(t, err)
		assert.Equal(t, "enrollment due date cant be before today", err.Error())
	})

	t.Run("should return error if installments total is less than total service", func(t *testing.T) {
		_ = service.ChangePrice(5000.00)
		enrollmentDueDate := time.Now().AddDate(0, 0, 2).Format("2006-01-02")
		reg, err := New(
			classRoom,
			inputData.Shift,
			student,
			service,
			inputData.MonthlyFee,
			inputData.InstallmentsQuantity,
			inputData.EnrollmentFee,
			enrollmentDueDate,
			inputData.MonthDuration,
			inputData.PaymentDay,
		)

		err = reg.Check()
		assert.Error(t, err)
		assert.Equal(t, "total paid not match with service price", err.Error())
	})
}

func getService() service.Service {
	svc, _ := service.New("Ensino Fundamental", 5000.00)
	return *svc
}

func getClassRoom() classroom.ClassRoom {

	clr, _ := classroom.New(
		10,
		"morning",
		"Jardim",
		"TUR-001",
		uuid.New().String(),
		uuid.New().String(),
		uuid.New().String(),
		"ANY",
		"remote",
	)

	return *clr
}

func getStudent(inputData RequestDto) student.Student {
	std, _ := student.New(
		inputData.Student.FirstName,
		inputData.Student.LastName,
		inputData.Student.Birthday,
		inputData.Student.RgDocument,
		inputData.Student.CpfDocument,
		inputData.Student.Email,
		inputData.Student.HimSelfResponsible,
	)

	std.AddPhones(inputData.Student.Phones)
	std.AddAddress(inputData.Student.Addresses)
	_ = std.AddParents(inputData.Student.Parents)

	return *std
}

func createInputData() RequestDto {

	addresses := []address.RequestDto{
		{
			Street:   "Rua dos Bobos",
			City:     "Marbule",
			District: "El Nildo",
			State:    "CC",
			ZipCode:  "47854-552",
		},
	}

	phones := []phone.RequestDto{
		{
			Description: "Pessoal",
			Phone:       "7199542-3264",
		},
	}

	parents := []parent.RequestDto{
		{
			FirstName:   "Marcha",
			LastName:    "Marbule",
			BirthDay:    "1970-05-05",
			Addresses:   addresses,
			Phones:      phones,
			RgDocument:  "745698885",
			CpfDocument: "624.499.720-48",
			Email:       "parent@mail.com",
		},
	}

	student := student.RequestDto{
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

	registration := RequestDto{
		ClassRoomId:          "76d0f9b7-3d7d-4cfc-892a-94c0704b4deb",
		Shift:                "morning",
		Student:              student,
		ServiceId:            "a1c003f8-8a56-4a49-892a-825b364cc076",
		MonthlyFee:           400.00,
		InstallmentsQuantity: 12,
		EnrollmentFee:        60.00,
		EnrollmentDueDate:    time.Now().Format("2006-01-02"),
		MonthDuration:        12,
		PaymentDay:           "16",
	}

	return registration
}
