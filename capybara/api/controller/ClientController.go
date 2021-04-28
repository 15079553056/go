package controller

import "github.com/gin-gonic/gin"


func ClientLogin(context *gin.Context) {
	context.JSON(200, map[string]interface{}{
		"message": "hello go.",
	})
}

//
