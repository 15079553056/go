package controller

import (
	"capybara/service"

	"github.com/gin-gonic/gin"
)

func ManagerLogin(context *gin.Context) {
	//获取表单信息
	username:=context.PostForm("username")
	password:=context.PostForm("password")


	if err:=service.ManagerLogin(username,password);err!=nil{
		context.JSON(200, map[string]interface{}{
			"message": "hello go.",
		})
	}else{
		context.JSON(200, map[string]interface{}{
			"message": "hello go.",
		})
	}	
}
