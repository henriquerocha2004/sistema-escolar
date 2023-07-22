package presenters

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities"
)

type WebSchoolYearPresenter struct {
	Id        string
	Year      string
	StartedAt string
	EndAt     string
}

func (w *WebSchoolYearPresenter) FillSingleSchoolYearPresenter(schoolYear entities.SchoolYear) *WebSchoolYearPresenter {
	return &WebSchoolYearPresenter{
		Id:        schoolYear.Id.String(),
		Year:      schoolYear.Year,
		StartedAt: schoolYear.StartedAt.Format("2006-01-02"),
		EndAt:     schoolYear.EndAt.Format("2006-01-02"),
	}
}

func (w *WebSchoolYearPresenter) FillMultipleSchoolYearPresenter(schoolYears []entities.SchoolYear) *[]WebSchoolYearPresenter {

	var presenters []WebSchoolYearPresenter

	for _, s := range schoolYears {
		presenters = append(presenters, *w.FillSingleSchoolYearPresenter(s))
	}

	return &presenters
}
