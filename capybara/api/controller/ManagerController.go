package controller

import (
	"capybara/service"

	"github.com/gin-gonic/gin"
)

func ManagerLogin(context *gin.Context) {
	service.ManagerLogin()

	context.JSON(200, map[string]interface{}{
		"message": "hello go.",
	})
}
