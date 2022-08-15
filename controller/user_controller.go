package controller

import (
	"gin/common"
	"gin/model"
	"gin/response"
	"gin/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, password, phone)
	//判断手机号是否存在
	if isPhoneExist(db, phone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号已存在")
		return
	}
	//创建用户
	//对密码进行加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "服务器错误")
		return
	}
	newUser := model.User{
		Name:     name,
		Password: string(hashedPassword),
		Phone:    phone,
	}
	db.Debug().Create(&newUser)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "注册成功",
	})
}

func Login(c *gin.Context) {
	db := common.GetDB()
	//获取参数
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	//数据验证
	if len(phone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		return
	}
	var user model.User
	db.Debug().Where("phone = ?", phone).First(&user)
	if user.ID == 0 {
		response.Response(c, http.StatusBadRequest, 400, nil, "用户不存在")
		return
	}
	//判断密码是否正确
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}
	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Println("token err", err)
		return
	}
	response.Response(c, http.StatusOK, 500, gin.H{
		"token": token,
	}, "系统异常")
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")
	response.Success(c, gin.H{
		"user": user,
	}, "登录成功")
}

func isPhoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Debug().Where("phone = ?", phone).First(&user)
	if user.ID == 0 {
		return false
	}
	return true
}
