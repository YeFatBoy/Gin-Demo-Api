/*
@Time : 2019/5/30 15:48
@Author : SuperShuYe
@File : category.go
@Software: GoLand
*/
package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/models"
	"web/models/category"
)

/**
	关联模型
 */
func CategoryList(c *gin.Context) {

	//验证
	type Query struct {
		Page         int    `form:"page"`
		PerPage      int    `form:"per_page"`
		CategoryName string `form:"category_name"`
	}
	var query Query
	if err := c.ShouldBind(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     err.Error(),
		})
		return
	}

	var Category category.Category
	CategoryList := models.Db.Find(&Category).Related(&Category.Topic, "CategoryTopic")
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     "success",
		"data":        CategoryList.Value,
	})

}
