package repositories

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres"
	testtools "github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres/test-tools"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/financial/service"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/classroom"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/registration"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/room"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/schedule"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/schoolyear"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/student"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/address"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/phone"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

type TestRegistrationSuit struct {
	suite.Suite
	connection           *sql.DB
	testTools            *testtools.DatabaseOperations
	repository           *RegistrationRepository
	schoolYearRepository *SchoolYearRepository
	scheduleRepository   *ScheduleRoomRepository
	classRoomRepository  *ClassRoomRepository
	roomRepository       *RoomRepository
	serviceRepository    *ServiceRepository
}

func newRegistrationRoomSuit(connection *sql.DB, testTools *testtools.DatabaseOperations) *TestRegistrationSuit {
	return &TestRegistrationSuit{
		testTools:  testTools,
		connection: connection,
	}
}

func (s *TestRegistrationSuit) SetupSuite() {
	s.repository = NewRegistrationRepository(s.connection)
	s.schoolYearRepository = NewSchoolYearRepository(s.connection)
	s.scheduleRepository = NewScheduleRoomRepository(s.connection)
	s.classRoomRepository = NewClassRoomRepository(s.connection)
	s.roomRepository = NewRoomRepository(s.connection)
	s.serviceRepository = NewServiceRepository(s.connection)
}

func (s *TestRegistrationSuit) AfterTest(suiteName, testName string) {
	s.testTools.RefreshDatabase()
}

func TestManagerRegistration(t *testing.T) {
	testtools.StartTestEnv()
	connection := postgres.Connect()
	suite.Run(t, newRegistrationRoomSuit(connection, testtools.NewTestDatabaseOperations(connection)))
}

func (s *TestRegistrationSuit) TestShouldCreateRegistration() {

	schoolYear := s.createSchoolYear()
	schedule := s.createSchedule(schoolYear)
	room := s.createRoom()
	service := s.createService()
	classroom := s.createClassRoom(schoolYear, room, schedule)

	stdent, err := student.New(
		"Pedrinho",
		"Souza",
		"2001-10-15",
		"123456789",
		"84731086043",
		"teste@test.com",
		true,
	)

	add := []address.RequestDto{
		{
			Street:   "Rua dos Bobos",
			City:     "SSA",
			District: "SC",
			State:    "SP",
			ZipCode:  "41500030",
		},
	}

	stdent.AddAddress(add)

	p := []phone.RequestDto{
		{
			Description: "Pessoal",
			Phone:       "71589955554",
		},
	}

	stdent.AddPhones(p)

	reg, err := registration.New(
		classroom,
		"morning",
		*stdent,
		service,
		2500.00,
		2,
		50.00,
		"2023-12-21",
		2,
		"12",
	)
	s.Assert().NoError(err)
	err = s.repository.Create(*reg)
	s.Assert().NoError(err)
}

func (s *TestRegistrationSuit) createSchoolYear() uuid.UUID {
	schoolYear, _ := schoolyear.New("2021", "2021-01-01", "2021-12-30")

	err := s.schoolYearRepository.Create(schoolYear)
	if err != nil {
		log.Fatalln(err)
	}

	return schoolYear.Id()
}

func (s *TestRegistrationSuit) createRoom() uuid.UUID {
	r, _ := room.New("SL-07", "Sala 7", 25)

	err := s.roomRepository.Create(*r)
	if err != nil {
		log.Fatalln(err)
	}

	return r.Id()
}

func (s *TestRegistrationSuit) createSchedule(schoolYearId uuid.UUID) uuid.UUID {
	sch, _ := schedule.New("Any Description", "08:00:00", "09:00:00", schoolYearId.String())

	err := s.scheduleRepository.Create(*sch)
	if err != nil {
		log.Fatalln(err)
	}

	return sch.Id()
}

func (s *TestRegistrationSuit) createService() service.Service {
	srvce, _ := service.New("ANY DESCRIPTION", 5000.00)

	err := s.serviceRepository.Create(*srvce)
	if err != nil {
		log.Fatalln(err)
	}

	return *srvce
}

func (s *TestRegistrationSuit) createClassRoom(
	schoolYearId uuid.UUID,
	roomId uuid.UUID,
	scheduleId uuid.UUID,
) classroom.ClassRoom {

	cr, _ := classroom.New(
		20,
		"morning",
		"1 ANO",
		"1AS",
		schoolYearId.String(),
		roomId.String(),
		scheduleId.String(),
		"any location",
		"in_person")

	err := s.classRoomRepository.Create(*cr)
	if err != nil {
		log.Fatalln(err)
	}

	return *cr
}
