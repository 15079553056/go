package routers

import (
	"capybara/controller"
	"capybara/middleware"

	"github.com/gin-gonic/gin"
)

func InitManagerRouter(Router *gin.RouterGroup) {
	ManagerRouter := Router.Group("manager").Use(middleware.OperationRecord())
	{
		ManagerRouter.GET("login", controller.ManagerLogin)
	}
}
