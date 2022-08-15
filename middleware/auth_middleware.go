package middleware

import (
	"gin/common"
	"gin/dto"
	"gin/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取authorization header
		tokenString := c.GetHeader("authorization")
		//格式验证
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 410,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 410,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		//获取token中的userId
		userId := claims.UserId
		db := common.GetDB()
		var user model.User
		db.First(&user, userId)
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 410,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		userDto := dto.ToUserDto(user)
		c.Set("user", userDto)
	}
}
