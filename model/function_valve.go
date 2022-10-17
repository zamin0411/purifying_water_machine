package model

type FunctionValve struct {
	ModelID    uint
	FunctionID uint
	ValveID    uint
	Status     bool `gorm:"function_valves_status"`
}
