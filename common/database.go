package common

import (
	"fmt"
	"gin/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	charset := viper.GetString("datasource.charset")
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", username, password, host, port, database, charset)
	println(url)
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
