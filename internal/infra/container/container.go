package container

import (
	"database/sql"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/database/postgres/repositories"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/http/controllers"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/financial/service"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/financial/service/serviceActions"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/classroom"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/classroom/classRoomService"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/registration"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/registration/registrationService"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/room"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/room/roomService"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/schedule"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/schedule/scheduleService"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/schoolyear"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/schoolyear/schoolYearService"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/student"
)

type ContainerDependency struct {
	db *sql.DB

	roomRepository         room.Repository
	schoolYearRepository   schoolyear.Repository
	scheduleRepository     schedule.Repository
	classRoomRepository    classroom.Repository
	serviceRepository      service.Repository
	registrationRepository registration.Repository
	studentRepository      student.Repository

	roomActions         roomService.ServiceRoomInterface
	schoolYearActions   schoolYearService.SchoolYearActionsInterface
	scheduleRoomActions scheduleService.ServiceScheduleInterface
	classRoomActions    classRoomService.ServiceClassRoomInterface
	serviceActions      serviceActions.ActionsServiceInterface
	registrationActions registrationService.RegistrationActionsInterface

	roomController          *controllers.RoomController
	schoolYearController    *controllers.SchoolYearController
	scheduleClassController *controllers.ScheduleController
	classRoomController     *controllers.ClassRoomController
	serviceController       *controllers.ServiceController
	registrationController  *controllers.RegisterController

	registerUow registration.RegisterUow
}

func (c *ContainerDependency) GetDB() *sql.DB {
	if c.db == nil {
		c.db = postgres.Connect()
	}

	return c.db
}

// Repository

func (c *ContainerDependency) GetRoomRepository() *room.Repository {
	if c.roomRepository == nil {
		c.roomRepository = repositories.NewRoomRepository(
			c.GetDB(),
		)
	}

	return &c.roomRepository
}

func (c *ContainerDependency) GetSchoolYearRepository() *schoolyear.Repository {
	if c.schoolYearRepository == nil {
		c.schoolYearRepository = repositories.NewSchoolYearRepository(
			c.GetDB(),
		)
	}

	return &c.schoolYearRepository
}

func (c *ContainerDependency) GetScheduleRepository() *schedule.Repository {
	if c.scheduleRepository == nil {
		c.scheduleRepository = repositories.NewScheduleRoomRepository(
			c.GetDB(),
		)
	}

	return &c.scheduleRepository
}

func (c *ContainerDependency) GetClassRoomRepository() *classroom.Repository {
	if c.classRoomRepository == nil {
		c.classRoomRepository = repositories.NewClassRoomRepository(
			c.GetDB(),
		)
	}

	return &c.classRoomRepository
}

func (c *ContainerDependency) GetServiceRepository() *service.Repository {
	if c.serviceRepository == nil {
		c.serviceRepository = repositories.NewServiceRepository(
			c.GetDB(),
		)
	}

	return &c.serviceRepository
}

func (c *ContainerDependency) GetRegisterRepository() *registration.Repository {
	if c.registrationRepository == nil {
		c.registrationRepository = repositories.NewRegistrationRepository(
			c.GetDB(),
		)
	}

	return &c.registrationRepository
}

func (c *ContainerDependency) GetStudentRepository() *student.Repository {
	if c.studentRepository == nil {
		c.studentRepository = repositories.NewStudentRepository(
			c.GetDB(),
		)
	}

	return &c.studentRepository
}

// Actions

func (c *ContainerDependency) GetRoomActions() roomService.ServiceRoomInterface {
	if c.roomActions == nil {
		c.roomActions = roomService.New(
			*c.GetRoomRepository(),
		)
	}

	return c.roomActions
}

func (c *ContainerDependency) GetSchoolYearActions() schoolYearService.SchoolYearActionsInterface {
	if c.schoolYearActions == nil {
		c.schoolYearActions = schoolYearService.New(
			*c.GetSchoolYearRepository(),
		)
	}

	return c.schoolYearActions
}

func (c *ContainerDependency) GetScheduleRoomActions() scheduleService.ServiceScheduleInterface {
	if c.scheduleRoomActions == nil {
		c.scheduleRoomActions = scheduleService.New(
			*c.GetScheduleRepository(),
			*c.GetSchoolYearRepository(),
		)
	}

	return c.scheduleRoomActions
}

func (c *ContainerDependency) GetClassRoomActions() classRoomService.ServiceClassRoomInterface {
	if c.classRoomActions == nil {
		c.classRoomActions = classRoomService.New(
			*c.GetClassRoomRepository(),
		)
	}

	return c.classRoomActions
}

func (c *ContainerDependency) GetServiceActions() serviceActions.ActionsServiceInterface {
	if c.serviceActions == nil {
		c.serviceActions = serviceActions.New(
			*c.GetServiceRepository(),
		)
	}

	return c.serviceActions
}

func (c *ContainerDependency) GetRegistrationActions() registrationService.RegistrationActionsInterface {
	if c.registrationActions == nil {
		c.registrationActions = registrationService.NewRegistrationActions(
			*c.GetServiceRepository(),
			*c.GetClassRoomRepository(),
			c.GetRegistrationUow(),
		)
	}

	return c.registrationActions
}

// Uow

func (c *ContainerDependency) GetRegistrationUow() registration.RegisterUow {
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
