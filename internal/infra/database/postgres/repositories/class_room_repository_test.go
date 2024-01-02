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

var currentTime time.Time = time.Now()

type TestClassRoomSuit struct {
	suite.Suite
	connection           *sql.DB
	testTools            *testtools.DatabaseOperations
	repository           *ClassRoomRepository
	schoolYearRepository *SchoolYearRepository
	roomRepository       *RoomRepository
	scheduleRepository   *ScheduleRoomRepository
}

func newTestClassRoomSuit(connection *sql.DB, testTools *testtools.DatabaseOperations) *TestClassRoomSuit {
	return &TestClassRoomSuit{
		testTools:  testTools,
		connection: connection,
	}
}

func (s *TestClassRoomSuit) SetupSuite() {
	s.repository = NewClassRoomRepository(s.connection)
	s.schoolYearRepository = NewSchoolYearRepository(s.connection)
	s.roomRepository = NewRoomRepository(s.connection)
	s.scheduleRepository = NewScheduleRoomRepository(s.connection)
}

func (s *TestClassRoomSuit) AfterTest(suiteName, testName string) {
	s.testTools.RefreshDatabase()
}

func TestManagerClassRoom(t *testing.T) {
	connection := postgres.Connect()
	suite.Run(t, newTestClassRoomSuit(connection, testtools.NewTestDatabaseOperations(connection)))
}

func (s *TestClassRoomSuit) TestShouldCreateClassRoom() {
	classRoom := s.createRoom()
	err := s.repository.Create(classRoom)
	s.Assert().NoError(err)
	classRoomDb, err := s.repository.FindById(classRoom.Id.String())
	s.Assert().NotNil(classRoomDb)
	s.Assert().NoError(err)
	s.Assert().Equal(classRoom.Localization, classRoomDb.Localization)
	s.Assert().Equal(classRoom.VacancyQuantity, classRoomDb.VacancyQuantity)
}

func (s *TestClassRoomSuit) TestShouldUpdateClassRoom() {
	classRoom := s.createRoom()
	err := s.repository.Create(classRoom)
	s.Assert().NoError(err)
	classRoom.Shift = "Vespertino"
	classRoom.Localization = "Outher Localization"
	err = s.repository.Update(classRoom)
	s.Assert().NoError(err)
	classRoomDb, err := s.repository.FindById(classRoom.Id.String())
	s.Assert().NoError(err)
	s.Assert().Equal(classRoom.Shift, classRoomDb.Shift)
	s.Assert().Equal(classRoom.Localization, classRoomDb.Localization)
}

func (s *TestClassRoomSuit) TestShouldDeleteClassRoom() {
	classRoom := s.createRoom()
	err := s.repository.Create(classRoom)
	s.Assert().NoError(err)
	err = s.repository.Delete(classRoom.Id.String())
	s.Assert().NoError(err)
	classRoomDb, err := s.repository.FindById(classRoom.Id.String())
	s.Assert().Nil(classRoomDb)
	s.Assert().Error(err)
}

func (s *TestClassRoomSuit) TestShouldFindClassRoomById() {
	classRoom := s.createRoom()
	err := s.repository.Create(classRoom)
	s.Assert().NoError(err)
	classRoomDb, err := s.repository.FindById(classRoom.Id.String())
	s.Assert().NotNil(classRoomDb)
	s.Assert().NoError(err)
}

func (s *TestClassRoomSuit) TestShouldFindByShift() {
	classRoom := s.createRoom()
	err := s.repository.Create(classRoom)
	s.Assert().NoError(err)

	pagination := common.Pagination{}
	pagination.ColumnSearch = append(pagination.ColumnSearch, dto.ColumnSearch{
		Column: "shift",
		Value:  "Matutino",
	})
	pagination.Limit = 1
	pagination.SetPage(1)

	classRoomPaginationResult, err := s.repository.FindAll(pagination)

	s.Assert().NoError(err)
	s.Assert().Equal(classRoom.Shift, classRoomPaginationResult.ClassRooms[0].Shift)
}

func (s *TestClassRoomSuit) getSchoolYear() entities.SchoolYear {

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

func (s *TestClassRoomSuit) getRoom() entities.Room {
	room := entities.Room{
		Id:          uuid.New(),
		Code:        "SL-07",
		Description: "Sala 7",
		Capacity:    25,
	}

	_ = s.roomRepository.Create(room)

	return room
}

func (s *TestClassRoomSuit) getSchedule(schoolYearId uuid.UUID) entities.ScheduleClass {

	schedule := entities.ScheduleClass{
		Id:          uuid.New(),
		Description: "Any Description",
		Schedule:    "8:00-9:00",
		SchoolYear:  schoolYearId,
	}

	_ = s.scheduleRepository.Create(schedule)

	return schedule
}

func (s *TestClassRoomSuit) createRoom() entities.ClassRoom {

	schoolYear := s.getSchoolYear()
	room := s.getRoom()
	schedule := s.getSchedule(schoolYear.Id)

	return entities.ClassRoom{
		Id:              uuid.New(),
		VacancyQuantity: 20,
		Shift:           "Matutino",
		OpenDate:        &currentTime,
		OccupiedVacancy: 0,
		Status:          "OPEN",
		Level:           "1 ANO",
		Identification:  "1AS",
		SchoolYearId:    schoolYear.Id,
		RoomId: uuid.NullUUID{
			UUID:  room.Id,
			Valid: false,
		},
		ScheduleId:   schedule.Id,
		Localization: "any location",
	}
}
