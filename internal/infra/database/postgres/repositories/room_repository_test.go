package repositories

import (
	"database/sql"
	"fmt"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres"
	testtools "github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres/test-tools"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/room"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/schedule"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/schoolyear"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/paginator"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestRoomSuit struct {
	suite.Suite
	connection           *sql.DB
	testTools            *testtools.DatabaseOperations
	repository           *RoomRepository
	schoolYearRepository *SchoolYearRepository
	scheduleRepository   *ScheduleRoomRepository
}

func newTestRoomSuit(connection *sql.DB, testTools *testtools.DatabaseOperations) *TestRoomSuit {
	return &TestRoomSuit{
		testTools:  testTools,
		connection: connection,
	}
}

func (s *TestRoomSuit) SetupSuite() {
	s.repository = NewRoomRepository(s.connection)
	s.schoolYearRepository = NewSchoolYearRepository(s.connection)
	s.scheduleRepository = NewScheduleRoomRepository(s.connection)
}

func (s *TestRoomSuit) AfterTest(suiteName, testName string) {
	s.testTools.RefreshDatabase()
}

func (s *TestRoomSuit) TearDownSuite() {
	fmt.Println("tear down")
	err := s.connection.Close()
	fmt.Println(err)
}

func TestManagerRoom(t *testing.T) {
	testtools.StartTestEnv()
	connection := postgres.Connect()
	suite.Run(t, newTestRoomSuit(connection, testtools.NewTestDatabaseOperations(connection)))
}

func (s *TestRoomSuit) TestShouldCreateRoom() {
	room, err := room.New("SL-07", "Sala 7", 25)
	s.Assert().NoError(err)

	err = s.repository.Create(*room)
	s.Assert().NoError(err)
}

func (s *TestRoomSuit) TestShouldUpdateRoom() {
	room, err := room.New("SL-07", "Sala 7", 25)
	s.Assert().NoError(err)

	err = s.repository.Create(*room)
	s.Assert().NoError(err)

	err = room.ChangeCapacity(15)
	s.Assert().NoError(err)

	err = s.repository.Update(*room)
	s.Assert().NoError(err)

	roomDb, err := s.repository.FindById(room.Id().String())
	s.Assert().NoError(err)
	s.Assert().NotNil(roomDb)
	s.Assert().Equal(room.Id(), roomDb.Id())
}

func (s *TestRoomSuit) TestShouldDeleteRoom() {
	room, err := room.New("SL-07", "Sala 7", 25)
	s.Assert().NoError(err)

	err = s.repository.Create(*room)
	s.Assert().NoError(err)

	err = s.repository.Delete(room.Id().String())
	s.Assert().NoError(err)

	roomDb, err := s.repository.FindById(room.Id().String())
	s.Assert().Error(err)
	s.Assert().Equal("sql: no rows in result set", err.Error())
	s.Assert().Nil(roomDb)
}

func (s *TestRoomSuit) TestShouldFindRoomById() {
	r, err := room.New("SL-07", "Sala 7", 25)
	s.Assert().NoError(err)

	err = s.repository.Create(*r)
	s.Assert().NoError(err)

	paginator := paginator.Pagination{}
	paginator.Limit = 3
	paginator.SortField = "created_at"
	paginator.Sort = "asc"
	paginator.SetPage(1)

	rooms, err := s.repository.FindAll(paginator)
	s.Assert().NoError(err)
	s.Assert().Equal(1, len(rooms.Data.([]room.Room)))
}

func (s *TestRoomSuit) TestShouldFindByCode() {
	r, err := room.New("SL-07", "Sala 7", 25)
	s.Assert().NoError(err)

	err = s.repository.Create(*r)
	s.Assert().NoError(err)

	roomDB, err := s.repository.FindByCode(r.Code())
	s.Assert().NoError(err)
	s.Assert().Equal(r.Code(), roomDB.Code())
}

func (s *TestRoomSuit) TestShouldSyncSchedule() {
	r, err := room.New("SL-07", "Sala 7", 25)
	s.Assert().NoError(err)

	err = s.repository.Create(*r)
	s.Assert().NoError(err)

	schoolYear, err := schoolyear.New("2021", "2023-01-01", "2023-12-21")
	s.Assert().NoError(err)

	_ = s.schoolYearRepository.Create(schoolYear)

	sch, err := schedule.New("Any description", "08:00:00", "09:00:00", schoolYear.Id().String())
	s.Assert().NoError(err)

	_ = s.scheduleRepository.Create(*sch)

	roomScheduleDto := schedule.RoomScheduleDto{
		SchoolYear:  schoolYear.Id().String(),
		RoomId:      r.Id().String(),
		ScheduleIds: []string{sch.Id().String()},
	}

	err = s.repository.SyncSchedule(roomScheduleDto)
	s.Assert().NoError(err)
}
