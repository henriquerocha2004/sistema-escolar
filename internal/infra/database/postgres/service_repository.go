package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/common"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities"
)

type ServiceRepository struct {
	db     *sql.DB
	queues *Queries
}

func NewServiceRepository(db *sql.DB) *ServiceRepository {
	return &ServiceRepository{
		db:     db,
		queues: New(db),
	}
}

func (s *ServiceRepository) Create(service entities.Service) error {

	serviceModel := CreateServiceParams{
		ID:          service.Id,
		Description: service.Description,
		Price:       fmt.Sprintf("%f", service.Price),
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	return s.queues.CreateService(context.Background(), serviceModel)
}

func (s *ServiceRepository) Delete(id string) error {

	serviceId, _ := uuid.Parse(id)

	deleteParams := DeleteServiceParams{
		ID: serviceId,
		DeletedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	return s.queues.DeleteService(context.Background(), deleteParams)
}

func (s *ServiceRepository) Update(service entities.Service) error {

	serviceModel := UpdateServiceParams{
		ID:          service.Id,
		Description: service.Description,
		Price:       fmt.Sprintf("%f", service.Price),
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	return s.queues.UpdateService(context.Background(), serviceModel)
}

func (s *ServiceRepository) FindById(id string) (*entities.Service, error) {
	serviceId, _ := uuid.Parse(id)

	serviceModel, err := s.queues.FindServiceById(context.Background(), serviceId)

	if err != nil {
		return nil, err
	}

	price, _ := strconv.ParseFloat(serviceModel.Price, 64)

	service := entities.Service{
		Id:          serviceModel.ID,
		Description: serviceModel.Description,
		Price:       price,
	}

	return &service, nil
}

func (s *ServiceRepository) FindAll(paginator common.Pagination) (*common.ServicePaginationResult, error) {
	ctx, cancelQuery := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelQuery()

	query := "SELECT id, description, price FROM services WHERE description like $1 AND deleted_at IS NULL"
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

	var services []entities.Service

	for rows.Next() {
		var service entities.Service
		err = rows.Scan(&service.Id, &service.Description, &service.Price)
		if err != nil {
			return nil, err
		}

		services = append(services, service)
	}

	var total int
	queryCount := "SELECT COUNT(*) as total FROM services WHERE description like $1 AND deleted_at IS NULL"
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

	paginationResult := common.ServicePaginationResult{
		Total:    total,
		Services: services,
	}

	return &paginationResult, nil
}
