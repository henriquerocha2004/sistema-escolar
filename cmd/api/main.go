package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/http/routes"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func init() {
	rootProject, _ := os.Getwd()
	err := godotenv.Load(rootProject + "/.env")
	if err != nil {
		log.Fatal("Error in read .env file")
	}
}

func main() {

	var data time.Time
	fmt.Println(data.String())

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))
	routes.GetRoutes(app)
	port := os.Getenv("APP_PORT")
	err := app.Listen(":" + port)
	if err != nil {
		panic(err)
	}
}
