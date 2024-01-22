package address

type RequestDto struct {
	Street   string `json:"street"`
	City     string `json:"city"`
	District string `json:"district"`
	State    string `json:"state"`
	ZipCode  string `json:"zip_code"`
}
