package routers

import (
	"capybara/controller"
	"capybara/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	ClientRouter := Router.Group("client").Use(middleware.OperationRecord())
	{
		ClientRouter.GET("login", controller.ClientLogin)
	}
}


