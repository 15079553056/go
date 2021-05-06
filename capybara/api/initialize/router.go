package initialize

import (
	"capybara/routers"

	"github.com/gin-gonic/gin"
)

//初始化总路由
func Routers() *gin.Engine {
	//创建 gin 引擎
	var engine = gin.Default()

	//使用静态文件(用户头像，音视频文件等)

	PrivateGroup := engine.Group("")

	PrivateGroup.Use()
	{
		routers.InitUserRouter(PrivateGroup)    //注册用户路由
		routers.InitManagerRouter(PrivateGroup) //注册后台管理员路由
	}

	return engine
}
