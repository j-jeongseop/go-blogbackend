package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/happynet78/go-blogbackend/database"
	"github.com/happynet78/go-blogbackend/routes"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":" + port)
}
