package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/container"
)

func setRegisterRoutes(app *fiber.App, container *container.ContainerDependency) {
	register := app.Group("register")
	register.Post("/", container.GetRegisterController().Create)
}
