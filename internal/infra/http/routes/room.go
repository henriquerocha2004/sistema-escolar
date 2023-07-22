package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/container"
)

func setRoomRoutes(app *fiber.App, container *container.ContainerDependency) {
	room := app.Group("room")
	room.Get("/", container.GetRoomController().FindAll)
	room.Get("/:id", container.GetRoomController().Find)
	room.Post("/", container.GetRoomController().Create)
	room.Put("/:id", container.GetRoomController().Update)
	room.Delete("/:id", container.GetRoomController().Delete)
	room.Post("sync-schedule", container.GetRoomController().SyncSchedule)
}
