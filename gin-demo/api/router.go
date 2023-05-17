package api

import (
	"gin-demo/api/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	r.Use(middleware.CORS())

	r.POST("/register", register) // 注册

	r.POST("/login", login) // 登录

	r.Run(":8080")
}
