package postgres

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	testtools "github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres/test-tools"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/common"
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

type TestSchoolYearSuit struct {
	suite.Suite
	connection *sql.DB
	testTools  *testtools.DatabaseOperations
	repository *SchoolYearRepository
}

func newTestSchoolYearSuit(connection *sql.DB, testTools *testtools.DatabaseOperations) *TestRoomSuit {
	return &TestRoomSuit{
		testTools:  testTools,
		connection: connection,
	}
}

func (s *TestSchoolYearSuit) SetupSuite() {
	s.repository = NewSchoolYearRepository(s.connection)
}

func (s *TestSchoolYearSuit) AfterTest(suiteName, testName string) {
	s.testTools.RefreshDatabase()
}

func TestManagerSchoolYear(t *testing.T) {
	connection := Connect()
	suite.Run(t, newTestSchoolYearSuit(connection, testtools.NewTestDatabaseOperations(connection)))
}

func (s *TestSchoolYearSuit) TestShouldCreateSchoolYear() {

	startAt, _ := time.Parse("2006-02-02", "2001-01-01")
	endAt, _ := time.Parse("2006-02-02", "2001-12-30")

	schoolYear := entities.SchoolYear{
		Id:        uuid.New(),
		Year:      "2001",
		StartedAt: &startAt,
		EndAt:     &endAt,
	}

	err := s.repository.Create(schoolYear)
	s.Assert().NoError(err)
}

func (s *TestSchoolYearSuit) TestShouldUpdateSchoolYear() {
	startAt, _ := time.Parse("2006-02-02", "2001-01-01")
	endAt, _ := time.Parse("2006-02-02", "2001-12-30")

	schoolYear := entities.SchoolYear{
		Id:        uuid.New(),
		Year:      "2001",
		StartedAt: &startAt,
		EndAt:     &endAt,
	}

	err := s.repository.Create(schoolYear)
	s.Assert().NoError(err)

	schoolYear.Year = "2002"
	err = s.repository.Update(schoolYear)
	s.Assert().NoError(err)

	schoolYearDb, err := s.repository.FindById(schoolYear.Id.String())
	s.Assert().NoError(err)
	s.Assert().NotNil(schoolYearDb)
	s.Assert().Equal(schoolYear.Id, schoolYearDb.Id)
	s.Assert().Equal(schoolYear.Year, schoolYearDb.Year)
}

func (s *TestSchoolYearSuit) TestShouldDeleteSchoolYear() {

	startAt, _ := time.Parse("2006-02-02", "2001-01-01")
	endAt, _ := time.Parse("2006-02-02", "2001-12-30")

	schoolYear := entities.SchoolYear{
		Id:        uuid.New(),
		Year:      "2001",
		StartedAt: &startAt,
		EndAt:     &endAt,
	}

	err := s.repository.Create(schoolYear)
	s.Assert().NoError(err)

	err = s.repository.Delete(schoolYear.Id.String())
	s.Assert().NoError(err)

	schoolYearDb, err := s.repository.FindById(schoolYear.Id.String())
	s.Assert().Error(err)
	s.Assert().Equal("sql: no rows in result set", err.Error())
	s.Assert().Nil(schoolYearDb)
}

func (s *TestSchoolYearSuit) TestShouldFindSchoolYearById() {

	startAt, _ := time.Parse("2006-02-02", "2001-01-01")
	endAt, _ := time.Parse("2006-02-02", "2001-12-30")

	schoolYear := entities.SchoolYear{
		Id:        uuid.New(),
		Year:      "2001",
		StartedAt: &startAt,
		EndAt:     &endAt,
	}

	err := s.repository.Create(schoolYear)
	s.Assert().NoError(err)

	paginator := common.Pagination{}
	paginator.Limit = 3
	paginator.SortField = "created_at"
	paginator.Sort = "asc"
	paginator.SetPage(1)

	schoolYears, err := s.repository.FindAll(paginator)
	s.Assert().NoError(err)
	s.Assert().Equal(1, len(*schoolYears))
}
