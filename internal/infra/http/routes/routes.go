package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/container"
)

func GetRoutes(app *fiber.App) {
	var di = &container.ContainerDependency{}
	setRoomRoutes(app, di)
	setSchoolYearRoutes(app, di)
	setSchedulesRoutes(app, di)
	setClassRoomRoutes(app, di)
	setServiceRoutes(app, di)
}
