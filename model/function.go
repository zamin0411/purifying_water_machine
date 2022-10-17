package model

type Function struct {
	ID       uint `gorm:"column:function_id"`
	Function string `gorm:"column:function"`
}
