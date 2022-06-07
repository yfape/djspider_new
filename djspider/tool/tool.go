/*
 * @Author: your name
 * @Date: 2021-07-12 22:57:06
 * @LastEditTime: 2021-08-03 13:07:08
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider\tool\tool.go
 */
package tool

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"strings"

	"github.com/golang-module/carbon"
)

//自动判断创建目录
func AutoDir(addr string) {
	jc := func(a string) {
		if _, err := os.Stat(a); err != nil {
			if err := os.Mkdir(a, 0755); err != nil {
				panic("创建目录失败")
			}
		}
	}

	arr := strings.Split(addr, "/")
	if len(arr) > 1 {
		temppath := arr[0]
		for i := 1; i < len(arr); i++ {
			temppath += "/" + arr[i]
			jc(temppath)
		}
	}
}

//获取今日日期
func DateToday() string {
	return strings.Replace(carbon.Now().ToDateString(), "-", "", -1)
}

func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}
