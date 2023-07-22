package dto

type RoomInformationDto struct {
	Room        RoomRequestDto         `json:"room"`
	SchoolYears []SchoolYearRequestDto `json:"school_years"`
}
