package controller

import (
	"gin/common"
	"gin/model"
	"gin/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	db := common.GetDB()
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
		name = util.RandomString(10)
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
	newUser := model.User{
		Name:     name,
		Password: password,
		Phone:    phone,
	}
	db.Debug().Create(&newUser)
	c.JSON(200, gin.H{
		"message": "注册成功",
	})
}

func isPhoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Debug().Where("phone = ?", phone).First(&user)
	if user.ID == 0 {
		return false
	}
	return true
}
