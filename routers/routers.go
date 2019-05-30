/*
@Time : 2019/5/28 17:30
@Author : SuperShuYe
@File : router.go
@Software: GoLand
*/
package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/api/controller"
	"web/middlewares"
)

func InitRouter() *gin.Engine {
	/*gin.DisableConsoleColor()
	// Logging to a file.
	accessLog, _ := os.Create("./log/gin.log")
	gin.DefaultWriter = io.MultiWriter(accessLog)*/
	//create new gin
	router := gin.Default()
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"StatusNotFound": "404",
		})
	})
	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{
			"StatusBadRequest": "400",
		})
	})

	//jwt登录注册
	jwt := router.Group("/api/v1/jwt")
	jwt.Use()
	{
		jwt.GET("/login",controller.UserLogin) //用户登录
	}
	//v1版本接口
	v1 := router.Group("/api/v1")
	v1.Use(middlewares.JWTAuth())
	{
		v1.GET("/users",controller.UserList) //用户列表
		v1.POST("/users",controller.AddUser) //增加用户
		v1.PATCH("/users/:id",controller.EditUser) //修改用户
		v1.DELETE("/users/:id",controller.DeleteUser) //删除用户
	}

	return router
}