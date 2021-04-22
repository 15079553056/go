package controller

import "github.com/gin-gonic/gin"

type ClientController struct {
}

func (client *ClientController) Router(engine *gin.Engine) {
	engine.GET("/ClientLogin", client.ClientLogin)
}

func (client *ClientController) ClientLogin(context *gin.Context) {
	context.JSON(200,map[string]interface{}{
		"message":"hello go.",
	})
}
