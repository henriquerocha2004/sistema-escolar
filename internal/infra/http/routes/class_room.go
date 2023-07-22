package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/container"
)

func setClassRoomRoutes(app *fiber.App, container *container.ContainerDependency) {
	classRoom := app.Group("class-room")
	classRoom.Get("/", container.GetClassRoomController().FindAll)
	classRoom.Get("/:id", container.GetClassRoomController().Find)
	classRoom.Post("/", container.GetClassRoomController().Create)
	classRoom.Put("/:id", container.GetClassRoomController().Update)
	classRoom.Delete("/:id", container.GetClassRoomController().Delete)
}
