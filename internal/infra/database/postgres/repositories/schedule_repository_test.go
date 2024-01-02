package repositories

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres"
	testtools "github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres/test-tools"
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
	schoolYear := s.getSchoolYear()

	schedule := entities.ScheduleClass{
		Id:          uuid.New(),
		Description: "Any Description",
		Schedule:    "8:00-9:00",
		SchoolYear:  schoolYear.Id,
	}

	err := s.repository.Create(schedule)
	s.Assert().NoError(err)
}

func (s *TestScheduleRoomSuit) TestShouldUpdateSchedule() {
	schoolYear := s.getSchoolYear()

	schedule := entities.ScheduleClass{
		Id:          uuid.New(),
		Description: "Any Description",
		Schedule:    "8:00-9:00",
		SchoolYear:  schoolYear.Id,
	}

	err := s.repository.Create(schedule)
	s.Assert().NoError(err)

	schedule.Schedule = "10:00-11:00"
	err = s.repository.Update(schedule)
	s.Assert().NoError(err)

	scheduleDb, err := s.repository.FindById(schedule.Id.String())
	s.Assert().NoError(err)
	s.Assert().Equal(schedule.Schedule, scheduleDb.Schedule)
}

func (s *TestScheduleRoomSuit) TestShouldDeleteScheduleRoom() {

	schoolYear := s.getSchoolYear()

	schedule := entities.ScheduleClass{
		Id:          uuid.New(),
		Description: "Any Description",
		Schedule:    "8:00-9:00",
		SchoolYear:  schoolYear.Id,
	}

	err := s.repository.Create(schedule)
	s.Assert().NoError(err)

	err = s.repository.Delete(schedule.Id.String())
	s.Assert().NoError(err)

	scheduleDb, err := s.repository.FindById(schedule.Id.String())
	s.Assert().Error(err)
	s.Assert().Equal("sql: no rows in result set", err.Error())
	s.Assert().Nil(scheduleDb)
}

func (s *TestScheduleRoomSuit) TestShouldFindByDescription() {

	schoolYear := s.getSchoolYear()

	schedule := entities.ScheduleClass{
		Id:          uuid.New(),
		Description: "Any Description",
		Schedule:    "8:00-9:00",
		SchoolYear:  schoolYear.Id,
	}

	err := s.repository.Create(schedule)
	s.Assert().NoError(err)

	pagination := common.Pagination{}
	pagination.ColumnSearch = append(pagination.ColumnSearch, dto.ColumnSearch{
		Column: "description",
		Value:  "Any Description",
	})
	pagination.Limit = 1
	pagination.SetPage(1)

	schedulePaginationResult, err := s.repository.FindAll(pagination)
	s.Assert().NoError(err)
	s.Assert().Equal(schedule.Description, schedulePaginationResult.Schedules[0].Description)
}

func (s *TestScheduleRoomSuit) TestShouldFindScheduleById() {
	schoolYear := s.getSchoolYear()

	schedule := entities.ScheduleClass{
		Id:          uuid.New(),
		Description: "Any Description",
		Schedule:    "8:00-9:00",
		SchoolYear:  schoolYear.Id,
	}

	err := s.repository.Create(schedule)
	s.Assert().NoError(err)

	scheduleDb, err := s.repository.FindById(schedule.Id.String())
	s.Assert().NoError(err)
	s.Assert().Equal(schedule.Description, scheduleDb.Description)
}

func (s *TestScheduleRoomSuit) getSchoolYear() entities.SchoolYear {

	startAt, _ := time.Parse("2006-02-02", "2001-01-01")
	endAt, _ := time.Parse("2006-02-02", "2001-12-30")

	schoolYear := entities.SchoolYear{
		Id:        uuid.New(),
		Year:      "2001",
		StartedAt: &startAt,
		EndAt:     &endAt,
	}

	_ = s.schoolYearRepository.Create(schoolYear)

	return schoolYear
}
