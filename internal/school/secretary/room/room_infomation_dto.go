package room

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/schoolyear"
)

type RoomInformationDto struct {
	Room        RoomRequestDto                    `json:"room"`
	SchoolYears []schoolyear.SchoolYearRequestDto `json:"school_years"`
}
