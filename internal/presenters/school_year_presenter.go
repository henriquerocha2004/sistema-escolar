package presenters

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/paginator"
)

type WebSchoolYearPaginationResultPresenter struct {
	Total       int                      `json:"total"`
	SchoolYears []WebSchoolYearPresenter `json:"school_years"`
}

type WebSchoolYearPresenter struct {
	Id        string `json:"id"`
	Year      string `json:"year"`
	StartedAt string `json:"start_at"`
	EndAt     string `json:"end_at"`
}

func (w *WebSchoolYearPresenter) FillSingleSchoolYearPresenter(schoolYear entities.SchoolYear) *WebSchoolYearPresenter {
	return &WebSchoolYearPresenter{
		Id:        schoolYear.Id.String(),
		Year:      schoolYear.Year,
		StartedAt: schoolYear.StartedAt.Format("2006-01-02"),
		EndAt:     schoolYear.EndAt.Format("2006-01-02"),
	}
}

func (w *WebSchoolYearPresenter) FillMultipleSchoolYearPresenter(paginationResult paginator.SchoolYearPaginationResult) WebSchoolYearPaginationResultPresenter {

	var presenter WebSchoolYearPaginationResultPresenter
	presenter.Total = paginationResult.Total

	for _, s := range paginationResult.SchoolYears {
		presenter.SchoolYears = append(presenter.SchoolYears, *w.FillSingleSchoolYearPresenter(s))
	}

	return presenter
}
