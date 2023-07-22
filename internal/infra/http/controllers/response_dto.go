package controllers

type responseDto struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponseDto(status string, message string, data interface{}) responseDto {
	return responseDto{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
