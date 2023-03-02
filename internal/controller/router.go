package controller

import (
	"github.com/gin-gonic/gin"
	docs "github.com/EinStealth/TentativeBackend/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	// sample
	r.GET("/ping", Ping)

	// sample
	r.POST("/throw", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.JSON(200, gin.H{
			"username": username,
			"password": password,
		})
	})

	// v1
	v1 := r.Group("/api/v1")
	{
		// 指定された合言葉の(ステータス)を返す
		v1.GET("/room", GetRoom)

		// 指定された合言葉の(ステータス)を返す
		v1.POST("/room", PostRoom)
	}

	// swagger ui
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}