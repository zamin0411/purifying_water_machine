package handler

import (
	"evg-purifying-water-machine/database"
	"evg-purifying-water-machine/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UpdateStatus(c *fiber.Ctx) error {
	db := database.DB
	var status model.WaterStatus
	if err := c.BodyParser(&status); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	//Create Water Status Transaction
	db.Transaction(func(tx *gorm.DB) error {
		if err := db.Omit("ColorIndex").Create(&status).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update", "data": err})
		}

		var id string
		db.Raw("SELECT @UUID").Scan(&id)
		status.ColorIndex.WaterStatusID = id
		status.ColorIndex.ID = id

		if err := db.Create(&status.ColorIndex).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update", "data": err})
		}
		return nil
	})

	return c.JSON(fiber.Map{"status": "success", "message": "Updated", "data": status})

}

func GetWaterStatus(c *fiber.Ctx) error {
	db := database.DB
	var status []model.WaterStatus
	db.Preload("ColorIndex").Find(&status)
	return c.JSON(fiber.Map{"status": "success", "message": "list water status", "data": status})
}
