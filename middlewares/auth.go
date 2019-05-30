/*
@Time : 2019/5/30 9:33
@Author : SuperShuYe
@File : auth.go
@Software: GoLand
*/
package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("auth")
	}
}
