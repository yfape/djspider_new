/*
 * @Author: yfape
 * @Date: 2021-07-09 16:23:49
 * @LastEditTime: 2021-07-28 16:31:52
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider\app\middle.go
 */
package app

import (
	"djspider/config"
	"djspider/db"
	"djspider/http"
	"log"
	"math"
	"sync"
	"time"

	"github.com/robfig/cron"

	"djspider/myredis"
)

var res_list []db.Article
var index = 1

var tw bool

func Mid_SpiderUrl() {
	url := "http://www.cac.gov.cn/wlaq/A0905index_1.htm"
	spi := db.GetSpiByUrl(url)
	var wg sync.WaitGroup
	wg.Add(1)
	go Fin_Spider(spi, 1, &wg)
	wg.Wait()
	defer func() {
		res_list = []db.Article{}
	}()
}

func Mid_RunOnce() {
	if tw {
		Fin_Tip()
	}
	spis := db.GetSpis()
	log.Println("---------- START:", index, " -----------")
	log.Println("开始爬取")
	if sl := len(spis); sl >= 6 {
		sarg := int(math.Floor(float64(sl/3) + 0.5))
		var wg sync.WaitGroup
		wg.Add(3)
		go Fin_Spider(spis[0:sarg], 1, &wg)
		go Fin_Spider(spis[sarg:sarg*2], 1, &wg)
		go Fin_Spider(spis[sarg*2:], 1, &wg)
		wg.Wait()
	} else {
		Fin_Spider(spis, 2, &sync.WaitGroup{})
	}

	if len(res_list) <= 0 {
		log.Println("无新增文章")
		log.Print("------------- END --------------\n\n")
		return
	}
	log.Printf("新增文章数量：%v \n", len(res_list))
	log.Println("开始推送")
	users := db.GetUsers()
	if ul := len(users); ul >= 6 {
		sarg := int(math.Floor(float64(ul/3) + 0.5))
		var wg sync.WaitGroup
		wg.Add(3)
		go Fin_Email(users[0:sarg], 1, &wg)
		go Fin_Email(users[sarg:sarg*2], 1, &wg)
		go Fin_Email(users[sarg*2:], 1, &wg)
		wg.Wait()
	} else {
		Fin_Email(users, 2, &sync.WaitGroup{})
	}
	log.Print("------------- END --------------\n\n")

	defer func() {
		res_list = []db.Article{}
	}()
}

func Mid_Worker() {
	c := cron.New()
	c.AddFunc(config.System.Period, func() {
		Mid_RunOnce()
		index++
	})
	log.Println("定时爬取进程：启动")
	tw = true
	Mid_RunOnce()
	c.Start()
	defer func() {
		c.Stop()
		tw = false
		log.Println("定时爬取进程：停止")
	}()
	select {}
}

func Mid_Http() {
	http.Create()
}

func Fin_Spider(spis []db.Spi, mode int, wg *sync.WaitGroup) {
	var spider Spider
	for i := 0; i < len(spis); i++ {
		chan1 := make(chan []db.Article)
		spider.Set(chan1, spis[i])
		go spider.Run()
		for list := range chan1 {
			res_list = append(res_list, list...)
		}
		if err := recover(); err != nil {
			log.Printf("%s\n", err)
		}
	}
	if mode == 1 {
		wg.Add(-1)
	}
}

func Fin_Email(users []db.User, mode int, wg *sync.WaitGroup) {
	var email Email
	for i := 0; i < len(users); i++ {
		email.Set(users[i], &res_list)
		email.Send()
	}
	if mode == 1 {
		wg.Add(-1)
	}
}

func Fin_Tip() {
	rd := myredis.Connect()
	defer (*rd).Close()
	_, err := (*rd).Do("Set", "lastspider", time.Now().Unix())
	if err != nil {
		log.Println(err)
	}
	(*rd).Do("Expire", "lastspider", 300)
	_, err = (*rd).Do("DEL", "centre")
	if err != nil {
		log.Println(err)
	}
}
