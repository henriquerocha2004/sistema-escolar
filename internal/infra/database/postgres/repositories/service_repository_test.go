package repositories

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres"
	testTools "github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres/test-tools"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/common"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

func init() {
	rootProject, _ := os.Getwd()
	err := godotenv.Load(rootProject + "/../../../../.env.test")
	if err != nil {
		log.Fatal("Error in read .env file")
	}
}

type TestServiceSuit struct {
	suite.Suite
	connection *sql.DB
	testTools  *testTools.DatabaseOperations
	repository *ServiceRepository
}

func newTestServiceSuit(connection *sql.DB, testTools *testTools.DatabaseOperations) *TestServiceSuit {
	return &TestServiceSuit{
		testTools:  testTools,
		connection: connection,
	}
}

func (s *TestServiceSuit) SetupSuite() {
	s.repository = NewServiceRepository(s.connection)
}

func (s *TestServiceSuit) AfterTest(suiteName, testName string) {
	s.testTools.RefreshDatabase()
}

func TestManagerService(t *testing.T) {
	connection := postgres.Connect()
	suite.Run(t, newTestServiceSuit(connection, testTools.NewTestDatabaseOperations(connection)))
}

func (s *TestServiceSuit) TestShouldCreateService() {

	service := entities.Service{
		Id:          uuid.New(),
		Description: "Ensino Fundamental",
		Price:       440.00,
	}

	err := s.repository.Create(service)
	s.Assert().NoError(err)
}

func (s *TestServiceSuit) TestShouldUpdateSchoolYear() {
	service := entities.Service{
		Id:          uuid.New(),
		Description: "Ensino Fundamental",
		Price:       440.00,
	}

	err := s.repository.Create(service)
	s.Assert().NoError(err)

	service.Description = "Ensino Médio"
	err = s.repository.Update(service)
	s.Assert().NoError(err)

	serviceDb, err := s.repository.FindById(service.Id.String())
	s.Assert().NoError(err)
	s.Assert().Equal(service.Description, serviceDb.Description)
}

func (s *TestServiceSuit) TestShouldDeleteSchoolYear() {

	service := entities.Service{
		Id:          uuid.New(),
		Description: "Ensino Fundamental",
		Price:       440.00,
	}

	err := s.repository.Create(service)
	s.Assert().NoError(err)

	err = s.repository.Delete(service.Id.String())
	s.Assert().NoError(err)

	serviceDb, err := s.repository.FindById(service.Id.String())
	s.Assert().Error(err)
	s.Assert().Nil(serviceDb)
}

func (s *TestServiceSuit) TestShouldFindByServiceDescription() {
	service := entities.Service{
		Id:          uuid.New(),
		Description: "Ensino Fundamental",
		Price:       440.00,
	}

	err := s.repository.Create(service)
	s.Assert().NoError(err)

	pagination := common.Pagination{}
	pagination.ColumnSearch = append(pagination.ColumnSearch, dto.ColumnSearch{
		Column: "description",
		Value:  "Ensino Fundamental",
	})
	pagination.SetPage(1)

	paginationResult, err := s.repository.FindAll(pagination)
	s.Assert().NoError(err)
	s.Assert().Equal(1, len(paginationResult.Services))
}
