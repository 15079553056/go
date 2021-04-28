package core

import "capybara/initialize"

//开启服务器
func RunWindowsServer() {
	//初始化redis服务

	//初始化路由
	Router := initialize.Routers()
	//Router.Static("", "")
	Router.Run(":9999")

}
