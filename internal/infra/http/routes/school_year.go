package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/container"
)

func setSchoolYearRoutes(app *fiber.App, container *container.ContainerDependency) {
	schoolYear := app.Group("school-year")
	schoolYear.Get("/", container.GetSchoolYearController().FindAll)
	schoolYear.Get("/:id", container.GetSchoolYearController().Find)
	schoolYear.Post("/", container.GetSchoolYearController().Create)
	schoolYear.Put("/:id", container.GetSchoolYearController().Update)
	schoolYear.Delete("/:id", container.GetSchoolYearController().Delete)
}
