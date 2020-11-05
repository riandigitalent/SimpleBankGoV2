package config

import (
	"MVC-GOLANG-ORM/app/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func DBInit() * gorm.DB{
	db,err := gorm.Open(mysql.Open("root:usbw@tcp(127.0.0.1:3307)/simplebankdb?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect DB" + err.Error())
	}

	// automigrate here
	err = db.AutoMigrate(new(model.Account), new(model.Transaction))
	if err != nil {
		panic(err)
	}

	return db
}