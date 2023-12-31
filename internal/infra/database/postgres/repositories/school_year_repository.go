package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"

	"github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres/models"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/common"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities"
)

type SchoolYearRepository struct {
	db     *sql.DB
	queues *models.Queries
}

func NewSchoolYearRepository(db *sql.DB) *SchoolYearRepository {
	return &SchoolYearRepository{
		db:     db,
		queues: models.New(db),
	}
}

func (s *SchoolYearRepository) Create(schoolYear entities.SchoolYear) error {
	schoolYearModel := models.CreateYearSchoolParams{
		ID:      schoolYear.Id,
		Year:    schoolYear.Year,
		StartAt: *schoolYear.StartedAt,
		EndAt:   *schoolYear.EndAt,
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	err := s.queues.CreateYearSchool(context.Background(), schoolYearModel)
	if err != nil {
		return err
	}

	return nil
}

func (s *SchoolYearRepository) Delete(id string) error {
	schoolYearId, _ := uuid.Parse(id)

	deleteParams := models.DeleteYearSchoolParams{
		DeletedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		ID: schoolYearId,
	}

	err := s.queues.DeleteYearSchool(context.Background(), deleteParams)
	return err
}

func (s *SchoolYearRepository) Update(schoolYear entities.SchoolYear) error {
	schoolYearModel := models.UpdateSchoolYearParams{
		Year:    schoolYear.Year,
		StartAt: *schoolYear.StartedAt,
		EndAt:   *schoolYear.EndAt,
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		ID: schoolYear.Id,
	}

	err := s.queues.UpdateSchoolYear(context.Background(), schoolYearModel)
	return err
}

func (s *SchoolYearRepository) FindById(id string) (*entities.SchoolYear, error) {

	schoolYearId, _ := uuid.Parse(id)
	schoolYearModel, err := s.queues.FindOneSchoolYear(context.Background(), schoolYearId)
	if err != nil {
		return nil, err
	}

	schoolYear := entities.SchoolYear{
		Id:        schoolYearModel.ID,
		Year:      schoolYearModel.Year,
		StartedAt: &schoolYearModel.StartAt,
		EndAt:     &schoolYearModel.EndAt,
	}

	return &schoolYear, nil
}

func (s *SchoolYearRepository) FindByYear(year string) (*entities.SchoolYear, error) {
	schoolYearModel, err := s.queues.FindByYear(context.Background(), year)
	if err != nil {
		return nil, err
	}

	schoolYear := entities.SchoolYear{
		Id:        schoolYearModel.ID,
		Year:      schoolYearModel.Year,
		StartedAt: &schoolYearModel.StartAt,
		EndAt:     &schoolYearModel.EndAt,
	}

	return &schoolYear, nil
}

func (s *SchoolYearRepository) FindAll(paginator common.Pagination) (*common.SchoolYearPaginationResult, error) {
	ctx, cancelQuery := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelQuery()

	query := "SELECT id as id, year, start_at, end_at FROM school_year WHERE year like $1 AND deleted_at IS NULL"
	filters := paginator.FiltersInSql()

	if filters != "" {
		query += filters
	}

	stmt, err := s.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx,
		"%"+paginator.Search+"%",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schoolYears []entities.SchoolYear

	for rows.Next() {
		var schoolYear entities.SchoolYear
		err = rows.Scan(&schoolYear.Id, &schoolYear.Year, &schoolYear.StartedAt, &schoolYear.EndAt)
		if err != nil {
			return nil, err
		}

		schoolYears = append(schoolYears, schoolYear)
	}

	var total int
	queryCount := "SELECT COUNT(*) as total FROM school_year WHERE year like $1 AND deleted_at IS NULL"
	columnFilter := paginator.GetColumnFilter()
	if columnFilter != "" {
		queryCount += columnFilter
	}

	stmtCount, err := s.db.PrepareContext(ctx, queryCount)
	if err != nil {
		return nil, err
	}

	defer stmtCount.Close()

	rows, err = stmtCount.QueryContext(ctx,
		"%"+paginator.Search+"%",
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&total)
		if err != nil {
			return nil, err
		}
	}

	paginationResult := common.SchoolYearPaginationResult{
		Total:       total,
		SchoolYears: schoolYears,
	}

	return &paginationResult, nil
}
