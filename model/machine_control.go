package model

type MachineControl struct {
	Id        string `gorm:"column:machine_control_id"`
	MachineID uint   `gorm:"column:machine_id"`
	Machine   Machine
	ValveID   uint `gorm:"column:valve_id"`
	Valve     Valve
	Status    bool `gorm:"column:machine_status"`
}
