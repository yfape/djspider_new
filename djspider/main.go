/*
 * @Author: your name
 * @Date: 2021-07-19 16:07:51
 * @LastEditTime: 2021-07-28 16:39:32
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider\main.go
 */
/*
 * @Author: your name
 * @Date: 2021-07-09 15:25:44
 * @LastEditTime: 2021-07-26 15:23:17
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider\app\main.go
 */
package main

import (
	"bufio"
	"djspider/app"
	"djspider/config"
	"fmt"
	"os"
	"strings"

	"djspider/deamon"
)

var reader *bufio.Reader

func main() {
	config.LoadConfig()
	//判断命令行是否带有参数
	if len(os.Args) > 1 {
		Back()
		return
	}
	Welcome()
	reader = bufio.NewReader(os.Stdin)
mainfor:
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, " ", "", -1)
		text = string([]byte(text)[:len(text)-1])
		switch text {
		case "exit":
			fmt.Println("再见")
			break mainfor
		case "once":
			app.Mid_RunOnce()
		case "run":
			app.Mid_Worker()
		case "run-d":
			GoWorker("ppid", "run", "定时爬虫")
		case "http":
			app.Mid_Http()
		case "http-d":
			GoWorker("hpid", "http", "HTTP")
		default:
			fmt.Println("未找到相关命令，请输入 help 查看命令列表")
		}
	}
}
func Back() {
	arg := os.Args[1]
	if arg == "run" {
		app.Mid_Worker()
	} else if arg == "http" {
		app.Mid_Http()
	} else if arg == "once" {
		app.Mid_RunOnce()
	}
}
func Welcome() {
	fmt.Print("\n \n")
	fmt.Println("  ||----------------------------------------------- ")
	fmt.Println("  ||		|  url      	爬取指定地址            ")
	fmt.Println("  ||	党	|  once     	运行一次爬取程序         ")
	fmt.Println("  ||	建	|  run  [-d]	启动定爬取进程	        ")
	fmt.Println("  ||	爬	|  http [-d]	启动HTTP服务进程	        ")
	fmt.Println("  ||	虫	|  							           ")
	fmt.Println("  ||		|  exit     	退出")
	fmt.Println("  ||----------------------------------------------- ")
	fmt.Println("  ||               ")
	fmt.Print("\n \n")
	fmt.Println("请输入指令：")
}
func GoWorker(key string, command string, msg string) {
	pid := deamon.RunDae(key, command)
	if pid == "" {
		fmt.Println("后台" + msg + "进程：已启动")
	} else {
		fmt.Print("是否关闭此后台进程 (y/n): ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, " ", "", -1)
		text = string([]byte(text)[:len(text)-1])
		if text == "y" || text == "Y" {
			if deamon.KillPid(pid, key) {
				fmt.Println("后台" + msg + "进程：已关闭")
			}
		}
	}
}
