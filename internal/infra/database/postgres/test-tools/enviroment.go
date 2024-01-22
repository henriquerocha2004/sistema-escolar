package testtools

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func StartTestEnv() {
	rootProject, _ := os.Getwd()
	err := godotenv.Load(rootProject + "/../../../../../.env.test")
	if err != nil {
		log.Fatal("Error in read .env file")
	}
}
