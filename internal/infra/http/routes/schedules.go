package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/container"
)

func setSchedulesRoutes(app *fiber.App, container *container.ContainerDependency) {
	schedules := app.Group("schedule")
	schedules.Get("/", container.GetScheduleRoomController().FindAll)
	schedules.Get("/:id", container.GetScheduleRoomController().Find)
	schedules.Post("/", container.GetScheduleRoomController().Create)
	schedules.Put("/:id", container.GetScheduleRoomController().Update)
	schedules.Delete("/:id", container.GetScheduleRoomController().Delete)
	schedules.Post("sync-schedule", container.GetScheduleRoomController().SyncSchedule)
}
