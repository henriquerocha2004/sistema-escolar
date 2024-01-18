package repositories

import (
	"database/sql"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/financial/service"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/paginator"
	"log"
	"os"
	"testing"

	"github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres"
	testTools "github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres/test-tools"
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
	service, err := service.New("Ensino Fundamental", 440.00)
	s.Assert().NoError(err)
	err = s.repository.Create(*service)
	s.Assert().NoError(err)
}

func (s *TestServiceSuit) TestShouldUpdateSchoolYear() {

	service, err := service.New("Ensino Fundamental", 440.00)
	s.Assert().NoError(err)

	err = s.repository.Create(*service)
	s.Assert().NoError(err)

	err = service.ChangeDescription("Ensino Médio")
	s.Assert().NoError(err)

	err = s.repository.Update(*service)
	s.Assert().NoError(err)

	serviceDb, err := s.repository.FindById(service.Id().String())
	s.Assert().NoError(err)
	s.Assert().Equal(service.Description, serviceDb.Description)
}

func (s *TestServiceSuit) TestShouldDeleteSchoolYear() {

	service, err := service.New("Ensino Fundamental", 440.00)
	s.Assert().NoError(err)

	err = s.repository.Create(*service)
	s.Assert().NoError(err)

	err = s.repository.Delete(service.Id().String())
	s.Assert().NoError(err)

	serviceDb, err := s.repository.FindById(service.Id().String())
	s.Assert().Error(err)
	s.Assert().Nil(serviceDb)
}

func (s *TestServiceSuit) TestShouldFindByServiceDescription() {
	srvice, err := service.New("Ensino Fundamental", 440.00)
	s.Assert().NoError(err)

	err = s.repository.Create(*srvice)
	s.Assert().NoError(err)

	pagination := paginator.Pagination{}
	pagination.ColumnSearch = append(pagination.ColumnSearch, paginator.ColumnSearch{
		Column: "description",
		Value:  "Ensino Fundamental",
	})
	pagination.SetPage(1)

	paginationResult, err := s.repository.FindAll(pagination)
	s.Assert().NoError(err)
	s.Assert().Equal(1, len(paginationResult.Data.([]service.Service)))
}
