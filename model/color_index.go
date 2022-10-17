package model

type ColorIndex struct {
	ID            string `gorm:"column:color_index_id"`
	Red           uint   `gorm:"column:color_index_red" json:"red"`
	Green         uint   `gorm:"column:color_index_green" json:"green"`
	Blue          uint   `gorm:"column:color_index_blue" json:"blue"`
	WaterStatusID string `gorm:"water_status_id"`
}
