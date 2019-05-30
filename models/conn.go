/*
@Time : 2019/5/29 10:46
@Author : SuperShuYe
@File : conn.go
@Software: GoLand
*/
package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"web/config"
)

var Db *gorm.DB

func init(){
	//读取配置文件
	Setting := config.Yaml.GetStringMap("Mysql")
	Db, _ = gorm.Open("mysql", Setting["user"].(string)+":"+Setting["password"].(string)+"@tcp("+Setting["host"].(string)+")/"+Setting["database"].(string)+"?charset=utf8mb4&parseTime=True&loc=Local")
	//defer Db.Close()
}