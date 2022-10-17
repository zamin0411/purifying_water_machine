package model

type Machine struct {
	ID       string `gorm:"column:machine_id"`
	Name     string `gorm:"column:machine_name"`
	StatusID uint   `gorm:"column:status_id"`
	Status   Status
	ModelID  uint `gorm:"column:model_id"`
	Model    Model
	Location string  `gorm:"machine_location"`
	Users    []*User `gorm:"many2many:access_rights;"`
}
