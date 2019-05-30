/*
@Time : 2019/5/29 11:04
@Author : SuperShuYe
@File : user.go
@Software: GoLand
*/
package controller

import (
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	myjwt "web/middlewares"
	"web/models"
	"web/models/user"
)

// LoginResult 登录结果结构
type LoginResult struct {
	Token string `json:"token"`
	user.User
}

func UserLogin(c *gin.Context) {
	//验证
	type Query struct {
		UserName string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
	}
	var query Query
	if err := c.ShouldBind(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     err.Error(),
		})
		return
	}

	var User user.User
	err := models.Db.Where("username = ? and uid = ?", query.UserName, query.Password).Find(&User).Error
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     "用户不存在",
		})
		return
	}

	//jwt生成
	j := &myjwt.JWT{
		[]byte("jmf"),
	}
	claims := myjwt.CustomClaims{
		User.Id,
		User.UserName,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "jmf",                           //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     err.Error(),
		})
		return
	}

	data := LoginResult{
		token,
		User,
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     "登录成功！",
		"data":        data,
	})
}

func UserList(c *gin.Context) {

	//验证
	type Query struct {
		Page     int    `form:"page"`
		PerPage  int    `form:"per_page"`
		UserName string `form:"username"`
	}
	var query Query
	if err := c.ShouldBind(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     err.Error(),
		})
		return
	}

	var User []user.User
	userList := models.Db.Find(&User)

	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     "success",
		"data":        userList.Value,
	})
}

func AddUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     "Add",
	})
}

func EditUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     "Edit",
	})
}

func DeleteUser(c *gin.Context) {

	//验证
	type Query struct {
		Id int `uri:"id" binding:"required"`
	}
	var query Query
	if err := c.ShouldBindUri(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     err.Error(),
		})
		return
	}

	err := models.Db.Where("id = ?", query.Id).Delete(&user.User{}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     "DeleteSuccess",
		"data":        true,
	})
}
