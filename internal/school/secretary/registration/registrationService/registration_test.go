package registrationService

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/registration"
	"testing"
)

func TestShouldRegisterStudent(t *testing.T) {
	// dataInput := createInputData()
	// serviceRepo := new(mocks)

	// err := registrationActions.Create(dataInput)

}

func createInputData() registration.RegistrationDto {

	addresses := []registration.AddressDto{
		{
			Street:   "Rua dos Bobos",
			City:     "Marbule",
			District: "El Nildo",
			State:    "CC",
			ZipCode:  "47854-552",
		},
	}

	phones := []registration.PhoneDto{
		{
			Description: "Pessoal",
			Phone:       "7199542-3264",
		},
	}

	parents := []registration.ParentDto{
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

	student := registration.StudentDto{
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

	registration := registration.RegistrationDto{
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
