package postgres

import (
	"database/sql"
	"log"
	"os"
	"testing"

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

type TestRoomSuit struct {
	suite.Suite
	connection *sql.DB
	testTools  *testtools.DatabaseOperations
	repository *RoomRepository
}

func newTestRoomSuit(connection *sql.DB, testTools *testtools.DatabaseOperations) *TestRoomSuit {
	return &TestRoomSuit{
		testTools:  testTools,
		connection: connection,
	}
}

func (s *TestRoomSuit) SetupSuite() {
	s.repository = NewRoomRepository(s.connection)
}

func (s *TestRoomSuit) AfterTest(suiteName, testName string) {
	s.testTools.RefreshDatabase()
}

func (s *TestRoomSuit) TearDownSuite() {
	_ = s.connection.Close()
}

func TestManagerRoom(t *testing.T) {
	connection := Connect()
	suite.Run(t, newTestRoomSuit(connection, testtools.NewTestDatabaseOperations(connection)))
}

func (s *TestRoomSuit) TestShouldCreateRoom() {
	room := entities.Room{
		Id:          uuid.New(),
		Code:        "SL-07",
		Description: "Sala 7",
		Capacity:    25,
	}

	err := s.repository.Create(room)
	s.Assert().NoError(err)
}

func (s *TestRoomSuit) TestShouldUpdateRoom() {
	room := entities.Room{
		Id:          uuid.New(),
		Code:        "SL-07",
		Description: "Sala 7",
		Capacity:    25,
	}

	err := s.repository.Create(room)
	s.Assert().NoError(err)

	room.Capacity = 15
	err = s.repository.Update(room)
	s.Assert().NoError(err)

	roomDb, err := s.repository.FindById(room.Id.String())
	s.Assert().NoError(err)
	s.Assert().NotNil(roomDb)
	s.Assert().Equal(room.Id, roomDb.Id)
}

func (s *TestRoomSuit) TestShouldDeleteRoom() {
	room := entities.Room{
		Id:          uuid.New(),
		Code:        "SL-07",
		Description: "Sala 7",
		Capacity:    25,
	}

	err := s.repository.Create(room)
	s.Assert().NoError(err)

	err = s.repository.Delete(room.Id.String())
	s.Assert().NoError(err)

	roomDb, err := s.repository.FindById(room.Id.String())
	s.Assert().Error(err)
	s.Assert().Equal("sql: no rows in result set", err.Error())
	s.Assert().Nil(roomDb)
}

func (s *TestRoomSuit) TestShouldFindRoomById() {
	room := entities.Room{
		Id:          uuid.New(),
		Code:        "SL-07",
		Description: "Sala 7",
		Capacity:    25,
	}

	err := s.repository.Create(room)
	s.Assert().NoError(err)

	paginator := common.Pagination{}
	paginator.Limit = 3
	paginator.SortField = "created_at"
	paginator.Sort = "asc"
	paginator.SetPage(1)

	rooms, err := s.repository.FindAll(paginator)
	s.Assert().NoError(err)
	s.Assert().Equal(1, len(*rooms))
}
