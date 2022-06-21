package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
)

func main() {
	//连接数据库
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	// 绑定模型
	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouter()

	r.Run()
}
