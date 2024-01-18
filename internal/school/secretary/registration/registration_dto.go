package registration

type AddressDto struct {
	Street   string `json:"street"`
	City     string `json:"city"`
	District string `json:"district"`
	State    string `json:"state"`
	ZipCode  string `json:"zip_code"`
}

type PhoneDto struct {
	Description string `json:"description"`
	Phone       string `json:"phone"`
}

type ParentDto struct {
	FirstName   string       `json:"first_name"`
	LastName    string       `json:"last_name"`
	BirthDay    string       `json:"birth_day"`
	Addresses   []AddressDto `json:"addresses"`
	Phones      []PhoneDto   `json:"phones"`
	RgDocument  string       `json:"rg_document"`
	CpfDocument string       `json:"cpf_document"`
	Email       string       `json:"email"`
}

type StudentDto struct {
	FirstName          string       `json:"first_name"`
	LastName           string       `json:"last_name"`
	Birthday           string       `json:"birthday"`
	RgDocument         string       `json:"rg_document"`
	CpfDocument        string       `json:"cpf_document"`
	Email              string       `json:"email"`
	HimSelfResponsible bool         `json:"him_self_responsible"`
	Addresses          []AddressDto `json:"addresses"`
	Phones             []PhoneDto   `json:"phones"`
	Parents            []ParentDto  `json:"parents"`
}

type RegistrationDto struct {
	ClassRoomId          string     `json:"class_room_id"`
	Shift                string     `json:"shift"`
	Student              StudentDto `json:"student"`
	ServiceId            string     `json:"service_id"`
	MonthlyFee           float64    `json:"monthly_fee"`
	InstallmentsQuantity int        `json:"installments_quantity"`
	EnrollmentFee        float64    `json:"enrollment_due_date"`
	EnrollmentDueDate    string     `json:"due_date"`
	MonthDuration        int        `json:"month_duration"`
	PaymentDay           string     `json:"payment_day"`
}
