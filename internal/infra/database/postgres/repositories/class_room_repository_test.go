package repositories

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres"
	testtools "github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres/test-tools"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/classroom"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/room"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/schedule"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/schoolyear"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/paginator"
	"github.com/stretchr/testify/suite"
	"testing"
)

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
	testtools.StartTestEnv()
	connection := postgres.Connect()
	suite.Run(t, newTestClassRoomSuit(connection, testtools.NewTestDatabaseOperations(connection)))
}

func (s *TestClassRoomSuit) TestShouldCreateClassRoom() {
	classRoom := s.createRoom()
	err := s.repository.Create(classRoom)
	s.Assert().NoError(err)
	classRoomDb, err := s.repository.FindById(classRoom.Id().String())
	s.Assert().NotNil(classRoomDb)
	s.Assert().NoError(err)
	s.Assert().Equal(classRoom.Localization(), classRoomDb.Localization())
	s.Assert().Equal(classRoom.VacancyQuantity(), classRoomDb.VacancyQuantity())
}

func (s *TestClassRoomSuit) TestShouldUpdateClassRoom() {
	classRoom := s.createRoom()
	err := s.repository.Create(classRoom)
	s.Assert().NoError(err)
	_ = classRoom.ChangeShift("afternoon")
	_ = classRoom.ChangeLocalization("Other Localization")
	err = s.repository.Update(classRoom)
	s.Assert().NoError(err)
	classRoomDb, err := s.repository.FindById(classRoom.Id().String())
	s.Assert().NoError(err)
	s.Assert().Equal(classRoom.Shift(), classRoomDb.Shift())
	s.Assert().Equal(classRoom.Localization(), classRoomDb.Localization())
}

func (s *TestClassRoomSuit) TestShouldDeleteClassRoom() {
	classRoom := s.createRoom()
	err := s.repository.Create(classRoom)
	s.Assert().NoError(err)
	err = s.repository.Delete(classRoom.Id().String())
	s.Assert().NoError(err)
	classRoomDb, err := s.repository.FindById(classRoom.Id().String())
	s.Assert().Nil(classRoomDb)
	s.Assert().Error(err)
}

func (s *TestClassRoomSuit) TestShouldFindClassRoomById() {
	classRoom := s.createRoom()
	err := s.repository.Create(classRoom)
	s.Assert().NoError(err)
	classRoomDb, err := s.repository.FindById(classRoom.Id().String())
	s.Assert().NotNil(classRoomDb)
	s.Assert().NoError(err)
}

func (s *TestClassRoomSuit) TestShouldFindByShift() {
	classRoom := s.createRoom()
	err := s.repository.Create(classRoom)
	s.Assert().NoError(err)

	pagination := paginator.Pagination{}
	pagination.ColumnSearch = append(pagination.ColumnSearch, paginator.ColumnSearch{
		Column: "shift",
		Value:  "morning",
	})
	pagination.Limit = 1
	pagination.SetPage(1)

	classRoomPaginationResult, err := s.repository.FindAll(pagination)
	result := classRoomPaginationResult.Data.([]classroom.ClassRoom)
	s.Assert().NoError(err)
	s.Assert().Equal(classRoom.Shift(), result[0].Shift())
}

func (s *TestClassRoomSuit) getSchoolYear() schoolyear.SchoolYear {
	schoolYear, err := schoolyear.New("2001", "2001-01-01", "2001-12-30")
	s.Assert().NoError(err)
	_ = s.schoolYearRepository.Create(schoolYear)

	return *schoolYear
}

func (s *TestClassRoomSuit) getRoom() room.Room {
	r, err := room.New("SL-07", "Sala 7", 25)
	s.Assert().NoError(err)
	_ = s.roomRepository.Create(*r)

	return *r
}

func (s *TestClassRoomSuit) getSchedule(schoolYearId uuid.UUID) schedule.ScheduleClass {
	sh, err := schedule.New("Any Description", "08:00:00", "09:00:00", schoolYearId.String())
	s.Assert().NoError(err)
	_ = s.scheduleRepository.Create(*sh)

	return *sh
}

func (s *TestClassRoomSuit) createRoom() classroom.ClassRoom {

	schoolYear := s.getSchoolYear()
	r := s.getRoom()
	sh := s.getSchedule(schoolYear.Id())

	classRoom, err := classroom.New(
		20,
		"morning",
		"1 Ano",
		"1AS",
		schoolYear.Id().String(),
		r.Id().String(),
		sh.Id().String(),
		"any location",
		"in_person",
	)

	s.Assert().NoError(err)

	return *classRoom
}
