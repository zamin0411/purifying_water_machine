package handler

import (
	"evg-purifying-water-machine/database"
	"evg-purifying-water-machine/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func UpdateStatus(c *fiber.Ctx) error {
	db := database.DB
	// type InputStatus struct {
	// 	WaterStatus model.WaterStatus
	// 	ColorIndex  model.ColorIndex `json:"color_index"`
	// }

	var status model.WaterStatus
	// type ColorIndexInput struct {
	// 	ColorIndex model.ColorIndex `json:"color_index"`
	// }
	// var colorIndex ColorIndexInput

	if err := c.BodyParser(&status); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// if err := c.BodyParser(&colorIndex); err != nil {
	// 	return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	// }

	// colorIndex := model.ColorIndex{
	// 	Red:   status.ColorIndex.Red,
	// 	Green: status.ColorIndex.Green,
	// 	Blue:  status.ColorIndex.Blue,
	// }
	fmt.Print(status)
	if err := db.Omit("ColorIndex").Create(&status).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update", "data": err})
	}

	status.ColorIndex.ID = status.ID
	status.ColorIndex.WaterStatusID = status.ID

	// result := db.Raw("INSERT INTO purifying_water_machine.color_index (color_index_red,color_index_green,color_index_blue) VALUES (?,?,?) RETURNING color_index_id;", colorIndex.Red, colorIndex.Green, colorIndex.Blue)
	if err := db.Create(&status.ColorIndex).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update", "data": err})
	}
	// status.ID = colorIndex.ColorIndex.ID
	// fmt.Print(status.ID)
	// if err := db.Create(&status).Error; err != nil {
	// 	return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update", "data": err})
	// }

	return c.JSON(fiber.Map{"status": "success", "message": "Updated", "data": status})

}

func GetWaterStatus(c *fiber.Ctx) error {
	db := database.DB
	type OutputStatus struct {
		WaterStatus model.WaterStatus
		ColorIndex  model.ColorIndex
	}
	var status model.WaterStatus
	var colorIndex model.ColorIndex
	var output OutputStatus

	db.Find(&status).Scan(&output)
	db.Find(&colorIndex).Scan(&output)

	return c.JSON(fiber.Map{"status": "success", "message": "", "data": status})

}
