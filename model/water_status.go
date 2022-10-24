package model

type WaterStatus struct {
	ID             string     `gorm:"column:water_status_id" json:"id"`
	MachineID      uint       `gorm:"column:machine_id" json:"machine_id"`
	HardWaterIndex int        `gorm:"column:water_status_hard_water_index" json:"hard_water_index"`
	PHIndex        int        `gorm:"column:water_status_pH_index" json:"pH_index"`
	TurbidityIndex int        `gorm:"column:water_status_turbidity_index" json:"turbidity_index"`
	ColorIndex     ColorIndex `json:"color_index"`
}
