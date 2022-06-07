/*
 * @Author: yfape
 * @Date: 2021-07-12 23:07:48
 * @LastEditTime: 2021-07-20 15:57:24
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider\app\db_connect.go
 */
package db

import (
	"djspider/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func MysqlConnect() *sqlx.DB {
	source := config.Mysql.Username + ":" + config.Mysql.Password +
		"@tcp(" + config.Mysql.Host + ":" + config.Mysql.Port + ")/" + config.Mysql.Database +
		"?charset=utf8&parseTime=True"
	database, err := sqlx.Open("mysql", source)
	if err != nil {
		panic(err)
	}

	return database
}

func MysqlClose(Db *sqlx.DB) {
	Db.Close()
}
