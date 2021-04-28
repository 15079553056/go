package middleware

import "github.com/gin-gonic/gin"

//注册 gin 框架的中间件(中间件:在请求到达接口之前进行拦截，做一些操作)
func OperationRecord() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
