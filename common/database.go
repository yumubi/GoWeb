package common

import (
	"fmt"
	"gin/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB = initDB()

func initDB() *gorm.DB {
	url := "root:yuanxiao88110@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(url))
	if err != nil {
		fmt.Println(err)
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}

func GetDB() *gorm.DB {
	return DB
}
