package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/container"
)

func setServiceRoutes(app *fiber.App, container *container.ContainerDependency) {
	service := app.Group("service")
	service.Get("/", container.GetServiceController().FindAll)
	service.Get("/:id", container.GetServiceController().FindById)
	service.Post("/", container.GetServiceController().Create)
	service.Put("/:id", container.GetServiceController().Update)
	service.Delete("/:id", container.GetServiceController().Delete)
}
