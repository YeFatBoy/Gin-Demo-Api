/*
@Time : 2019/5/28 17:49
@Author : SuperShuYe
@File : config.go
@Software: GoLand
*/
package config

import (
	"github.com/spf13/viper"
	"log"
)

var Yaml *viper.Viper

func init(){
	Yaml = viper.New()
	Yaml.SetConfigName("dev")
	Yaml.AddConfigPath("./config/yaml/")
	//设置配置文件类型
	Yaml.SetConfigType("yaml")

	if err := Yaml.ReadInConfig(); err != nil{
		log.Println(err)
	}
}