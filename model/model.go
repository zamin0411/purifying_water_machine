package model

type Model struct {
	ID       uint `gorm:"column:model_id"`
	Model    string `gorm:"column:model"`
	Valves    []Valve    `gorm:"many2many:model_valves"`
	Functions []Function `gorm:"many2many:model_functions"`
}
