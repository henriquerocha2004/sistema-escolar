package container

import (
	"database/sql"

	"github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/http/controllers"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary"
)

type ContainerDependency struct {
	db *sql.DB

	roomRepository       secretary.RoomRepository
	schoolYearRepository secretary.SchoolYearRepository
	scheduleRepository   secretary.ScheduleRoomRepository
	classRoomRepository  secretary.ClassRoomRepository

	roomActions         secretary.RoomActionsInterface
	schoolYearActions   secretary.SchoolYearActionsInterface
	scheduleRoomActions secretary.ScheduleActionsInterface
	classRoomActions    secretary.ClassRoomActionsInterface

	roomController          *controllers.RoomController
	schoolYearController    *controllers.SchoolYearController
	scheduleClassController *controllers.ScheduleController
	classRoomController     *controllers.ClassRoomController
}

func (c *ContainerDependency) GetDB() *sql.DB {
	if c.db == nil {
		c.db = postgres.Connect()
	}

	return c.db
}

// Repository

func (c *ContainerDependency) GetRoomRepository() *secretary.RoomRepository {
	if c.roomRepository == nil {
		c.roomRepository = postgres.NewRoomRepository(
			c.GetDB(),
		)
	}

	return &c.roomRepository
}

func (c *ContainerDependency) GetSchoolYearRepository() *secretary.SchoolYearRepository {
	if c.schoolYearRepository == nil {
		c.schoolYearRepository = postgres.NewSchoolYearRepository(
			c.GetDB(),
		)
	}

	return &c.schoolYearRepository
}

func (c *ContainerDependency) GetScheduleRepository() *secretary.ScheduleRoomRepository {
	if c.scheduleRepository == nil {
		c.scheduleRepository = postgres.NewScheduleRoomRepository(
			c.GetDB(),
		)
	}

	return &c.scheduleRepository
}

func (c *ContainerDependency) GetClassRoomRepository() *secretary.ClassRoomRepository {
	if c.classRoomRepository == nil {
		c.classRoomRepository = postgres.NewClassRoomRepository(
			c.GetDB(),
		)
	}

	return &c.classRoomRepository
}

// Actions

func (c *ContainerDependency) GetRoomActions() secretary.RoomActionsInterface {
	if c.roomActions == nil {
		c.roomActions = secretary.NewRoomActions(
			*c.GetRoomRepository(),
		)
	}

	return c.roomActions
}

func (c *ContainerDependency) GetSchoolYearActions() secretary.SchoolYearActionsInterface {
	if c.schoolYearActions == nil {
		c.schoolYearActions = secretary.NewSchoolYearActions(
			*c.GetSchoolYearRepository(),
		)
	}

	return c.schoolYearActions
}

func (c *ContainerDependency) GetScheduleRoomActions() secretary.ScheduleActionsInterface {
	if c.scheduleRoomActions == nil {
		c.scheduleRoomActions = secretary.NewScheduleClassActions(
			*c.GetScheduleRepository(),
			*c.GetSchoolYearRepository(),
		)
	}

	return c.scheduleRoomActions
}

func (c *ContainerDependency) GetClassRoomActions() secretary.ClassRoomActionsInterface {
	if c.classRoomActions == nil {
		c.classRoomActions = secretary.NewClassRoomActions(
			*c.GetClassRoomRepository(),
		)
	}

	return c.classRoomActions
}

// Controllers

func (c *ContainerDependency) GetRoomController() *controllers.RoomController {
	if c.roomController == nil {
		c.roomController = controllers.NewRoomController(
			c.GetRoomActions(),
		)
	}

	return c.roomController
}

func (c *ContainerDependency) GetSchoolYearController() *controllers.SchoolYearController {
	if c.schoolYearController == nil {
		c.schoolYearController = controllers.NewSchoolYearController(
			c.GetSchoolYearActions(),
		)
	}

	return c.schoolYearController
}

func (c *ContainerDependency) GetScheduleRoomController() *controllers.ScheduleController {
	if c.scheduleClassController == nil {
		c.scheduleClassController = controllers.NewScheduleController(
			c.GetScheduleRoomActions(),
		)
	}

	return c.scheduleClassController
}

func (c *ContainerDependency) GetClassRoomController() *controllers.ClassRoomController {
	if c.classRoomController == nil {
		c.classRoomController = controllers.NewClassRoomController(
			c.GetClassRoomActions(),
		)
	}

	return c.classRoomController
}
