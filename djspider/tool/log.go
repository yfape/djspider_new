/*
 * @Author: your name
 * @Date: 2021-07-13 09:18:11
 * @LastEditTime: 2021-07-14 09:49:59
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider\app\action\log.go
 */
package tool

import (
	"djspider/config"
)

func SaveLog(filename string, context string) {
	logaddr := config.File.Log + "/" + DateToday()
	AutoDir(logaddr)
}
