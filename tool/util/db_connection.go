package util

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDBConnection() *gorm.DB {
	user := "root"
	pass := "root"
	protocol := "tcp(command-box-db:3306)"
	dbName := "command_box"

	dsn := user + ":" + pass + "@" + protocol + "/" + dbName + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	return db
}
