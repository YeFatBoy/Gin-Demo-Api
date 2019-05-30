package main

import (
	"github.com/gin-gonic/gin"
	"web/routers"
)

func main() {
	gin.SetMode(gin.DebugMode)
	Router := routers.InitRouter()
	Router.Run()
}


