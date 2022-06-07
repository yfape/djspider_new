/*
 * @Author: yfape
 * @Date: 2021-07-12 23:05:00
 * @LastEditTime: 2021-07-23 15:15:02
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider\app\config.go
 */
package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	File   FileConfig
	MySql  MySqlConfig
	System SystemConfig
	Email  EmailConfig
	Auth   AuthConfig
}

type MySqlConfig struct {
	Port     string
	Host     string
	Username string
	Password string
	Database string
}

type FileConfig struct {
	Img string
	Log string
}

type SystemConfig struct {
	Period string
}

type EmailConfig struct {
	From     string
	Subject  string
	Username string
	Authcode string
}

type AuthConfig struct {
	Key       string
	ExpireDay int
	Issue     string
	Subject   string
}

//全局配置
var Mysql MySqlConfig
var File FileConfig
var System SystemConfig
var Email EmailConfig
var Auth AuthConfig

func LoadConfig() {
	var config Config
	//加载配置文件
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.Unmarshal(&config)

	Mysql = config.MySql
	File = config.File
	System = config.System
	Email = config.Email
	Auth = config.Auth
}
