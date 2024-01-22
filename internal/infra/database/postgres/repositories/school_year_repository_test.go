package repositories

import (
	"database/sql"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres"
	testtools "github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres/test-tools"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/schoolyear"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/paginator"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestSchoolYearSuit struct {
	suite.Suite
	connection *sql.DB
	testTools  *testtools.DatabaseOperations
	repository *SchoolYearRepository
}

func newTestSchoolYearSuit(connection *sql.DB, testTools *testtools.DatabaseOperations) *TestSchoolYearSuit {
	return &TestSchoolYearSuit{
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
	testtools.StartTestEnv()
	connection := postgres.Connect()
	suite.Run(t, newTestSchoolYearSuit(connection, testtools.NewTestDatabaseOperations(connection)))
}

func (s *TestSchoolYearSuit) TestShouldCreateSchoolYear() {
	sYear, err := schoolyear.New(
		"2001",
		"2001-01-01",
		"2001-12-30",
	)
	s.Assert().NoError(err)

	err = s.repository.Create(sYear)
	s.Assert().NoError(err)
}

func (s *TestSchoolYearSuit) TestShouldUpdateSchoolYear() {
	sYear, err := schoolyear.New(
		"2001",
		"2001-01-01",
		"2001-12-30",
	)
	s.Assert().NoError(err)

	err = s.repository.Create(sYear)
	s.Assert().NoError(err)

	_ = sYear.ChangeSchoolYear("2002")
	err = s.repository.Update(sYear)
	s.Assert().NoError(err)

	schoolYearDb, err := s.repository.FindById(sYear.Id().String())
	s.Assert().NoError(err)
	s.Assert().NotNil(schoolYearDb)
	s.Assert().Equal(sYear.Id(), schoolYearDb.Id())
	s.Assert().Equal(sYear.Year(), schoolYearDb.Year())
}

func (s *TestSchoolYearSuit) TestShouldDeleteSchoolYear() {

	sYear, err := schoolyear.New(
		"2001",
		"2001-01-01",
		"2001-12-30",
	)
	s.Assert().NoError(err)

	err = s.repository.Create(sYear)
	s.Assert().NoError(err)

	err = s.repository.Delete(sYear.Id().String())
	s.Assert().NoError(err)

	schoolYearDb, err := s.repository.FindById(sYear.Id().String())
	s.Assert().Error(err)
	s.Assert().Equal("sql: no rows in result set", err.Error())
	s.Assert().Nil(schoolYearDb)
}

func (s *TestSchoolYearSuit) TestShouldFindSchoolYearById() {

	sYear, err := schoolyear.New(
		"2001",
		"2001-01-01",
		"2001-12-30",
	)
	s.Assert().NoError(err)

	err = s.repository.Create(sYear)
	s.Assert().NoError(err)

	paginator := paginator.Pagination{}
	paginator.Limit = 3
	paginator.SortField = "created_at"
	paginator.Sort = "asc"
	paginator.SetPage(1)

	paginationResult, err := s.repository.FindAll(paginator)
	s.Assert().NoError(err)
	s.Assert().Equal(1, len(paginationResult.Data.([]schoolyear.SchoolYear)))
}

func (s *TestSchoolYearSuit) TestShouldFindSchoolYearByYear() {
	sYear, err := schoolyear.New(
		"2001",
		"2001-01-01",
		"2001-12-30",
	)
	s.Assert().NoError(err)

	err = s.repository.Create(sYear)
	s.Assert().NoError(err)

	schoolYearDb, err := s.repository.FindByYear(sYear.Year())
	s.Assert().NoError(err)
	s.Assert().Equal(sYear.Year(), schoolYearDb.Year())
}
