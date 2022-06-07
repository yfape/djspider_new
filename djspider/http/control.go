/*
 * @Author: your name
 * @Date: 2021-07-26 15:20:40
 * @LastEditTime: 2021-07-28 16:55:30
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider\app\http\control.go
 */
package http

import (
	"djspider/deamon"
	"os"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
)

func SwitchSpider(c *gin.Context) {
	pid := deamon.RunDae("ppid", "run")
	if pid == "" {
		c.JSON(200, gin.H{"msg": "启动定时爬虫进程：成功"})
	} else {
		if deamon.KillPid(pid, "ppid") {
			c.JSON(200, gin.H{"msg": "关闭定时爬虫进程：成功"})
		} else {
			c.JSON(522, gin.H{"msg": "关闭定时爬虫进程：失败"})
		}
	}
}

func SwitchHttp(c *gin.Context) {
	pid := deamon.RunDae("hpid", "http")
	if pid == "" {
		c.JSON(200, gin.H{"msg": "启动HTTP进程：成功"})
	} else {
		if deamon.KillPid(pid, "hpid") {
			c.JSON(200, gin.H{"msg": "关闭HTTP进程：成功"})
		} else {
			c.JSON(522, gin.H{"msg": "关闭HTTP进程：失败"})
		}
	}
}

func RunOnce(c *gin.Context) {
	args := os.Args
	for k, v := range args {
		if v == "-d" {
			args[k] = ""
		}
	}
	cmd := exec.Command(args[0], "once")
	cmd.Env = os.Environ()
	cmd.Start()
	time.Sleep(time.Second * 10)
	c.JSON(200, gin.H{"msg": "单次爬取推送进程：成功"})
}
