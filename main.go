package main

import (
	"gin/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {
	InitConfig()
	common.DB = common.InitDB()
	r := gin.Default()
	r = CollectRouter(r)
	port := ":" + viper.GetString("serve.port")
	r.Run(port)
}
func InitConfig() {
	workDir, _ := os.Getwd()
	//添加配置文件路径
	viper.AddConfigPath(workDir + "\\config")
	//添加配置文件名
	viper.SetConfigName("application")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
