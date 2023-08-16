package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/http/routes"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/value_objects"
	"github.com/joho/godotenv"
)

func init() {
	rootProject, _ := os.Getwd()
	err := godotenv.Load(rootProject + "/.env")
	if err != nil {
		log.Fatal("Error in read .env file")
	}
}

func main() {

	var cpf value_objects.CPF
	cpf = "035.808.175-07"
	err := cpf.Validate()
	fmt.Println(err)
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))
	routes.GetRoutes(app)
	port := os.Getenv("APP_PORT")
	err = app.Listen(":" + port)
	if err != nil {
		panic(err)
	}
}
