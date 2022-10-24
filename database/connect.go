package database

import (
	"evg-purifying-water-machine/config"
	"evg-purifying-water-machine/model"
	"fmt"

	"gorm.io/driver/mysql"

	//"gorm.io/driver/postgres"
	//"strconv"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func Connect() {

	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), config.Config("DB_PORT"), config.Config("DB_NAME"))
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	//PostgreSQL connect
	// var err error
	// p := config.Config("DB_PORT")
	// port, err := strconv.ParseUint(p, 10, 32)
	// dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
	// 	NamingStrategy: schema.NamingStrategy{
	// 		SingularTable: true,
	// 	},
	// })
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&model.WaterStatus{}, &model.ColorIndex{})
	// DB.AutoMigrate(&model.Machine{}, &model.Model{}, &model.WaterStatus{}, &model.Status{}, &model.User{}, &model.Valve{}, &model.Function{}, &model.ColorIndex{}, &model.FunctionValve{}, &model.MachineControl{})
	fmt.Println("Database Migrated")
}
