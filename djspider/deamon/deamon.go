/*
 * @Author: your name
 * @Date: 2021-07-26 15:31:48
 * @LastEditTime: 2021-07-28 16:17:48
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider\deamon\deamon.go
 */
package deamon

import (
	"djspider/myredis"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/garyburd/redigo/redis"
)

func RunDae(pidname string, command string) string {
	if s := JudgePid(pidname); s != "" {
		fmt.Println("后台存在相同进程：", s)
		return s
	}

	args := os.Args
	for k, v := range args {
		if v == "-d" {
			args[k] = ""
		}
	}

	cmd := exec.Command(args[0], command)
	cmd.Env = os.Environ()
	cmd.Start()
	SavePid(pidname, cmd.Process.Pid)
	return ""
}
func JudgePid(key string) string {
	rd := myredis.Connect()
	b, err := redis.Bool((*rd).Do("EXISTS", key))
	if err != nil {
		panic(err)
	}
	defer (*rd).Close()
	if b {
		res, err := redis.String((*rd).Do("GET", key))
		if err != nil {
			panic(err)
		}
		return res
	} else {
		return ""
	}
}
func SavePid(key string, pid int) {
	rd := myredis.Connect()
	_, err := (*rd).Do("Set", key, pid)
	if err != nil {
		panic(err)
	}
	defer (*rd).Close()
}
func DeletePid(key string) {
	rd := myredis.Connect()
	_, err := (*rd).Do("DEL", key)
	if err != nil {
		panic(err)
	}
	(*rd).Do("DEL", "centre")
	(*rd).Do("DEL", "lastspider")
	defer (*rd).Close()
}
func KillPid(pid string, key string) bool {
	sysType := runtime.GOOS
	var fc string
	if sysType == "Linux" {
		fc = "kill"
	} else if sysType == "windows" {
		fc = "tskill"
	} else {
		log.Println("暂不支持此操作系统")
	}
	DeletePid(key)
	cmd := exec.Command(fc, pid)
	cmd.Env = os.Environ()
	cmd.Start()
	log.Println("停止进程：" + key)
	return true
}
