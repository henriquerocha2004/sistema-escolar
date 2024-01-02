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
	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/value_objects"
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
	connection := postgres.Connect()
	suite.Run(t, newRegistrationRoomSuit(connection, testtools.NewTestDatabaseOperations(connection)))
}

func (s *TestRegistrationSuit) TestShouldCreateRegistration() {

	schoolYear := s.createSchoolYear()
	schedule := s.createSchedule(schoolYear)
	room := s.createRoom()
	service := s.createService()
	classroom := s.createClassRoom(schoolYear, room, schedule)

	student := entities.Student{
		Id:                 uuid.New(),
		FirstName:          "Pedrinho",
		LastName:           "Souza",
		BirthDay:           &time.Time{},
		RgDocument:         "123456789",
		CPFDocument:        value_objects.CPF("17515874698"),
		Email:              "teste@test.com",
		HimSelfResponsible: true,
	}

	address := []dto.AddressDto{
		{
			Street:   "Rua dos Bobos",
			City:     "SSA",
			District: "SC",
			State:    "SP",
			ZipCode:  "41500030",
		},
	}

	student.AddAddress(address)

	phone := []dto.PhoneDto{
		{
			Description: "Pessoal",
			Phone:       "71589955554",
		},
	}

	student.AddPhones(phone)

	registration := entities.Registration{
		Id:                   uuid.New(),
		Class:                classroom,
		Student:              student,
		Shift:                value_objects.Shift("morning"),
		Service:              service,
		MonthlyFee:           2500.00,
		InstallmentsQuantity: 2,
		EnrollmentFee:        50.00,
		EnrollmentDueDate:    &currentTime,
		MonthDuration:        2,
		EnrollmentDate:       &currentTime,
		PaymentDay:           "12",
	}
	registration.GenerateCode()

	err := s.repository.Create(registration)
	s.Assert().NoError(err)

}

func (s *TestRegistrationSuit) createSchoolYear() uuid.UUID {

	now := time.Now()
	schoolYear := entities.SchoolYear{
		Id:        uuid.New(),
		Year:      "2021",
		StartedAt: &now,
		EndAt:     &now,
	}

	err := s.schoolYearRepository.Create(schoolYear)
	if err != nil {
		log.Fatalln(err)
	}

	return schoolYear.Id
}

func (s *TestRegistrationSuit) createRoom() uuid.UUID {
	room := entities.Room{
		Id:          uuid.New(),
		Code:        "SL-07",
		Description: "Sala 7",
		Capacity:    25,
	}

	err := s.roomRepository.Create(room)
	if err != nil {
		log.Fatalln(err)
	}

	return room.Id
}

func (s *TestRegistrationSuit) createSchedule(schoolYearId uuid.UUID) uuid.UUID {

	schedule := entities.ScheduleClass{
		Id:          uuid.New(),
		Schedule:    "09:00-10:00",
		Description: "Any Description",
		SchoolYear:  schoolYearId,
	}

	err := s.scheduleRepository.Create(schedule)
	if err != nil {
		log.Fatalln(err)
	}

	return schedule.Id
}

func (s *TestRegistrationSuit) createService() entities.Service {

	service := entities.Service{
		Id:          uuid.New(),
		Description: "ANY DESCRIPTION",
		Price:       5000.00,
	}

	err := s.serviceRepository.Create(service)
	if err != nil {
		log.Fatalln(err)
	}

	return service
}

func (s *TestRegistrationSuit) createClassRoom(
	schoolYearId uuid.UUID,
	roomId uuid.UUID,
	scheduleId uuid.UUID,
) entities.ClassRoom {

	classRoom := entities.ClassRoom{
		Id:              uuid.New(),
		VacancyQuantity: 20,
		Shift:           "Matutino",
		OpenDate:        &currentTime,
		OccupiedVacancy: 0,
		Status:          "OPEN",
		Level:           "1 ANO",
		Identification:  "1AS",
		SchoolYearId:    scheduleId,
		RoomId: uuid.NullUUID{
			UUID:  roomId,
			Valid: false,
		},
		ScheduleId:   scheduleId,
		Localization: "any location",
	}

	err := s.classRoomRepository.Create(classRoom)
	if err != nil {
		log.Fatalln(err)
	}

	return classRoom
}
