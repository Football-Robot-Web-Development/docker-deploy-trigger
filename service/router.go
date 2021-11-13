package service

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()
	register(r)
	return r
}

func register(r *gin.Engine) {
	r.GET("deploy", Deploy)
}

