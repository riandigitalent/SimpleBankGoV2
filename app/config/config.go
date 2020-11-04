package config

import (
	//"MVC-GO-SQL/app/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInit() * gorm.DB{
	db,err := gorm.Open(mysql.Open("root:usbw@tcp(127.0.0.1:3307)/simplebankdb?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err !=nil {
		panic("Failed to Connect DB"+err.Error())

	}
	//db.AutoMigrate(new(model.Acount),new(model.Transaction))

	DB= db

	return db


}
