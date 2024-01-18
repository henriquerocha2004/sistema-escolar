package repositories

import (
	"database/sql"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/schedule"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/schoolyear"

	"github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres"
	testtools "github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres/test-tools"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/paginator"
	"log"
	"os"
	"testing"

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

type TestScheduleRoomSuit struct {
	suite.Suite
	connection           *sql.DB
	testTools            *testtools.DatabaseOperations
	repository           *ScheduleRoomRepository
	schoolYearRepository *SchoolYearRepository
}

func newTestScheduleRunRepository(connection *sql.DB, testTools *testtools.DatabaseOperations) *TestScheduleRoomSuit {
	return &TestScheduleRoomSuit{
		connection: connection,
		testTools:  testTools,
	}
}

func (s *TestScheduleRoomSuit) SetupSuite() {
	s.repository = NewScheduleRoomRepository(s.connection)
	s.schoolYearRepository = NewSchoolYearRepository(s.connection)
}

func (s *TestScheduleRoomSuit) AfterTest(suiteName, testName string) {
	s.testTools.RefreshDatabase()
}

func TestManagerScheduleRoom(t *testing.T) {
	connection := postgres.Connect()
	suite.Run(t, newTestScheduleRunRepository(connection, testtools.NewTestDatabaseOperations(connection)))
}

func (s *TestScheduleRoomSuit) TestShouldCreateSchedule() {
	sY := s.getSchoolYear()
	sch, err := schedule.New("Any Description", "08:00", "09:00", sY.Id().String())

	err = s.repository.Create(*sch)
	s.Assert().NoError(err)
}

func (s *TestScheduleRoomSuit) TestShouldUpdateSchedule() {
	sY := s.getSchoolYear()
	sch, err := schedule.New("Any Description", "08:00", "09:00", sY.Id().String())

	err = s.repository.Create(*sch)
	s.Assert().NoError(err)

	err = sch.ChangePeriod("10:00", "11:00")
	s.Assert().NoError(err)
	err = s.repository.Update(*sch)
	s.Assert().NoError(err)

	scheduleDb, err := s.repository.FindById(sch.Id().String())
	s.Assert().NoError(err)
	s.Assert().Equal(sch.StartAt(), scheduleDb.StartAt())
	s.Assert().Equal(sch.EndAt(), scheduleDb.EndAt())
}

func (s *TestScheduleRoomSuit) TestShouldDeleteScheduleRoom() {

	sY := s.getSchoolYear()
	sch, err := schedule.New("Any Description", "08:00", "09:00", sY.Id().String())

	err = s.repository.Create(*sch)
	s.Assert().NoError(err)

	err = s.repository.Delete(sch.Id().String())
	s.Assert().NoError(err)

	scheduleDb, err := s.repository.FindById(sch.Id().String())
	s.Assert().Error(err)
	s.Assert().Equal("sql: no rows in result set", err.Error())
	s.Assert().Nil(scheduleDb)
}

func (s *TestScheduleRoomSuit) TestShouldFindByDescription() {

	sY := s.getSchoolYear()
	sch, err := schedule.New("Any Description", "08:00", "09:00", sY.Id().String())

	err = s.repository.Create(*sch)
	s.Assert().NoError(err)

	pagination := paginator.Pagination{}
	pagination.ColumnSearch = append(pagination.ColumnSearch, paginator.ColumnSearch{
		Column: "description",
		Value:  "Any Description",
	})
	pagination.Limit = 1
	pagination.SetPage(1)

	schedulePaginationResult, err := s.repository.FindAll(pagination)
	s.Assert().NoError(err)
	data := schedulePaginationResult.Data.([]schedule.ScheduleClass)
	s.Assert().Equal(sch.Description(), data[0].Description)
}

func (s *TestScheduleRoomSuit) TestShouldFindScheduleById() {
	sY := s.getSchoolYear()

	sch, err := schedule.New("Any Description", "08:00", "09:00", sY.Id().String())

	err = s.repository.Create(*sch)
	s.Assert().NoError(err)

	scheduleDb, err := s.repository.FindById(sch.Id().String())
	s.Assert().NoError(err)
	s.Assert().Equal(sch.Description(), scheduleDb.Description())
}

func (s *TestScheduleRoomSuit) getSchoolYear() schoolyear.SchoolYear {
	sY, err := schoolyear.New("2001", "2001-01-01", "2001-12-30")
	s.Assert().NoError(err)
	_ = s.schoolYearRepository.Create(sY)

	return *sY
}
