package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/financial/service"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/paginator"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres/models"
)

type ServiceRepository struct {
	db     *sql.DB
	queues *models.Queries
}

type serviceSearchModel struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Price       string    `json:"price"`
	Total       int       `json:"total"`
}

func NewServiceRepository(db *sql.DB) *ServiceRepository {
	return &ServiceRepository{
		db:     db,
		queues: models.New(db),
	}
}

func (s *ServiceRepository) Create(service service.Service) error {

	serviceModel := models.CreateServiceParams{
		ID:          service.Id(),
		Description: service.Description(),
		Price:       fmt.Sprintf("%f", service.Price()),
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

	deleteParams := models.DeleteServiceParams{
		ID: serviceId,
		DeletedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	return s.queues.DeleteService(context.Background(), deleteParams)
}

func (s *ServiceRepository) Update(service service.Service) error {

	serviceModel := models.UpdateServiceParams{
		ID:          service.Id(),
		Description: service.Description(),
		Price:       fmt.Sprintf("%f", service.Price()),
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	return s.queues.UpdateService(context.Background(), serviceModel)
}

func (s *ServiceRepository) FindById(id string) (*service.Service, error) {
	serviceId, _ := uuid.Parse(id)
	serviceModel, err := s.queues.FindServiceById(context.Background(), serviceId)

	if err != nil {
		return nil, err
	}

	price, _ := strconv.ParseFloat(serviceModel.Price, 64)
	srvice, err := service.New(serviceModel.Description, price)
	err = srvice.ChangeId(serviceModel.ID.String())
	if err != nil {
		return nil, err
	}

	return srvice, nil
}

func (s *ServiceRepository) FindAll(pagination paginator.Pagination) (*paginator.PaginationResult, error) {
	ctx, cancelQuery := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelQuery()

	query := `SELECT id, description, price, COUNT(*) OVER() as total 
					FROM services 
				WHERE description like $1 AND deleted_at IS NULL`
	filters := pagination.FiltersInSql()

	if filters != "" {
		query += filters
	}

	stmt, err := s.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx,
		"%"+pagination.Search+"%",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var servicesModel []serviceSearchModel

	for rows.Next() {
		var serviceModel serviceSearchModel
		err = rows.Scan(&serviceModel.ID, &serviceModel.Description, &serviceModel.Price, &serviceModel.Total)
		if err != nil {
			return nil, err
		}

		servicesModel = append(servicesModel, serviceModel)
	}

	var services []service.Service

	for _, serviceModel := range servicesModel {

		price, err := strconv.ParseFloat(serviceModel.Price, 64)
		if err != nil {
			return nil, err
		}
		service, err := service.Load(
			serviceModel.ID.String(),
			serviceModel.Description,
			price,
		)

		if err != nil {
			return nil, err
		}

		services = append(services, *service)
	}

	paginationResult := paginator.PaginationResult{
		Total: servicesModel[0].Total,
		Data:  services,
	}

	return &paginationResult, nil
}
