package model

type Valve struct {
	ID    uint `gorm:"column:valve_id"`
	Valve string `gorm:"column:valve"`
}
