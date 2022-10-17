package model

type User struct {
	ID       string     `gorm:"column:user_id"`
	Password string     `gorm:"column:user_password"`
	Name     string     `gorm:"column:user_name"`
	Machines []*Machine `gorm:"many2many:access_rights;"`
}
