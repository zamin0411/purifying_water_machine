package model

type Status struct {
	ID     uint `gorm:"column:status_id"`
	Status string `gorm:"column:status"`
}
