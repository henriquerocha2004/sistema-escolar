package room

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/schoolyear"
)

type RoomInformationDto struct {
	Room        Request                 `json:"room"`
	SchoolYears []schoolyear.SchoolYear `json:"school_years"`
}
