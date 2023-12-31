package container

import (
	"database/sql"

	"github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres/repositories"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/http/controllers"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/financial"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/uow"
)

type ContainerDependency struct {
	db *sql.DB

	roomRepository         secretary.RoomRepository
	schoolYearRepository   secretary.SchoolYearRepository
	scheduleRepository     secretary.ScheduleRoomRepository
	classRoomRepository    secretary.ClassRoomRepository
	serviceRepository      financial.ServiceRepository
	registrationRepository secretary.RegistrationRepository
	studentRepository      secretary.StudentRepository

	roomActions         secretary.RoomActionsInterface
	schoolYearActions   secretary.SchoolYearActionsInterface
	scheduleRoomActions secretary.ScheduleActionsInterface
	classRoomActions    secretary.ClassRoomActionsInterface
	serviceActions      financial.ServiceActionsInterface
	registrationActions secretary.RegistrationActionsInterface

	roomController          *controllers.RoomController
	schoolYearController    *controllers.SchoolYearController
	scheduleClassController *controllers.ScheduleController
	classRoomController     *controllers.ClassRoomController
	serviceController       *controllers.ServiceController
	registrationController  *controllers.RegisterController

	registerUow uow.RegisterUow
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
		c.roomRepository = repositories.NewRoomRepository(
			c.GetDB(),
		)
	}

	return &c.roomRepository
}

func (c *ContainerDependency) GetSchoolYearRepository() *secretary.SchoolYearRepository {
	if c.schoolYearRepository == nil {
		c.schoolYearRepository = repositories.NewSchoolYearRepository(
			c.GetDB(),
		)
	}

	return &c.schoolYearRepository
}

func (c *ContainerDependency) GetScheduleRepository() *secretary.ScheduleRoomRepository {
	if c.scheduleRepository == nil {
		c.scheduleRepository = repositories.NewScheduleRoomRepository(
			c.GetDB(),
		)
	}

	return &c.scheduleRepository
}

func (c *ContainerDependency) GetClassRoomRepository() *secretary.ClassRoomRepository {
	if c.classRoomRepository == nil {
		c.classRoomRepository = repositories.NewClassRoomRepository(
			c.GetDB(),
		)
	}

	return &c.classRoomRepository
}

func (c *ContainerDependency) GetServiceRepository() *financial.ServiceRepository {
	if c.serviceRepository == nil {
		c.serviceRepository = repositories.NewServiceRepository(
			c.GetDB(),
		)
	}

	return &c.serviceRepository
}

func (c *ContainerDependency) GetRegisterRepository() *secretary.RegistrationRepository {
	if c.registrationRepository == nil {
		c.registrationRepository = repositories.NewRegistrationRepository(
			c.GetDB(),
		)
	}

	return &c.registrationRepository
}

func (c *ContainerDependency) GetStudentRepository() *secretary.StudentRepository {
	if c.studentRepository == nil {
		c.studentRepository = repositories.NewStudentRepository(
			c.GetDB(),
		)
	}

	return &c.studentRepository
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

func (c *ContainerDependency) GetServiceActions() financial.ServiceActionsInterface {
	if c.serviceActions == nil {
		c.serviceActions = financial.NewServiceActions(
			*c.GetServiceRepository(),
		)
	}

	return c.serviceActions
}

func (c *ContainerDependency) GetRegistrationActions() secretary.RegistrationActionsInterface {
	if c.registrationActions == nil {
		c.registrationActions = secretary.NewRegistrationActions(
			*c.GetServiceRepository(),
			*c.GetClassRoomRepository(),
			c.GetRegistrationUow(),
		)
	}

	return c.registrationActions
}

// Uow

func (c *ContainerDependency) GetRegistrationUow() uow.RegisterUow {
	if c.registerUow == nil {
		c.registerUow = repositories.NewRegistrationUow(
			c.GetDB(),
			*repositories.NewStudentRepository(c.GetDB()),
			*repositories.NewRegistrationRepository(c.GetDB()),
		)
	}

	return c.registerUow
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

func (c *ContainerDependency) GetServiceController() *controllers.ServiceController {
	if c.serviceController == nil {
		c.serviceController = controllers.NewServiceController(
			c.GetServiceActions(),
		)
	}

	return c.serviceController
}

func (c *ContainerDependency) GetRegisterController() *controllers.RegisterController {
	if c.registrationController == nil {
		c.registrationController = controllers.NewRegisterController(
			c.GetRegistrationActions(),
		)
	}

	return c.registrationController
}
