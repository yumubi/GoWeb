package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null"`
	Phone    string `gorm:"type:varchar(11);not null"`
	Password string `gorm:"type:varchar(255);not null"`
}

func main() {
	db := InitDB()
	r := gin.Default()
	r.POST("/api/auth/register", func(c *gin.Context) {
		//获取参数
		name := c.PostForm("name")
		password := c.PostForm("password")
		phone := c.PostForm("phone")
		//数据验证
		if len(phone) != 11 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "手机号必须为11位",
			})
			return
		}
		if len(password) < 6 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "密码不能小于6位",
			})
			return
		}
		if len(name) == 0 {
			name = RandomString(10)
		}
		log.Println(name, password, phone)
		//判断手机号是否存在
		if isPhoneExist(db, phone) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "手机号已存在",
			})
			return
		}
		//创建用户
		newUser := User{
			Name:     name,
			Password: password,
			Phone:    phone,
		}
		db.Debug().Create(&newUser)
		c.JSON(200, gin.H{
			"message": "注册成功",
		})
	})
	r.Run()
}

func isPhoneExist(db *gorm.DB, phone string) bool {
	var user User
	db.Debug().Where("phone = ?", phone).First(&user)
	if user.ID == 0 {
		return false
	}
	return true
}

func RandomString(i int) string {
	letters := []byte("sdnasdnauidnaodnaoicdjvdiwieujqpwugi")
	result := make([]byte, i)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func InitDB() *gorm.DB {
	url := "root:yuanxiao88110@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(url))
	if err != nil {
		fmt.Println(err)
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}
