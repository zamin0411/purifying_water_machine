package router

import (
	"evg-purifying-water-machine/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	machine := api.Group("/machines")
	machine.Post("/update-status", handler.UpdateStatus)
	machine.Get("/", handler.GetWaterStatus)

}
