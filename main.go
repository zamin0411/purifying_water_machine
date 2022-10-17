package main

import (
	"evg-purifying-water-machine/database"
	"evg-purifying-water-machine/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	database.Connect()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))

}
