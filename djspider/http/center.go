/*
 * @Author: your name
 * @Date: 2021-07-26 10:58:12
 * @LastEditTime: 2021-07-29 17:11:32
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider\app\http\center.go
 */
package http

import (
	"djspider/db"
	"djspider/myredis"
	"encoding/json"
	"log"

	"github.com/garyburd/redigo/redis"

	"github.com/gin-gonic/gin"
)

type Oa struct {
	Title string
	Url   string
	Site  string
	Time  string
}

type grap1 struct {
	Name  string
	Count int
}

var rd *redis.Conn

func Center(c *gin.Context) {
	rd = myredis.Connect()
	defer (*rd).Close()
	resout, bo := GetRedis1()
	if bo {
		c.String(201, resout)
		return
	}

	Db = db.MysqlConnect()
	defer Db.Close()
	var count int
	err := Db.Get(&count, "select count(*) as num from dj_article")
	if err != nil {
		log.Println(err)
		count = 0
	}
	var spidercount int
	err = Db.Get(&spidercount, "select count(*) as num from dj_spi")
	if err != nil {
		log.Println(err)
		spidercount = 0
	}
	var webcount int
	err = Db.Get(&webcount, "select count(col0_id) as num from dj_spi group by col0_id")
	if err != nil {
		log.Println(err)
		webcount = 0
	}
	var usercount int
	err = Db.Get(&usercount, "select count(*) as num from dj_user")
	if err != nil {
		log.Println(err)
		usercount = 0
	}
	var newcount int
	err = Db.Get(&newcount, "select count(*) as num from dj_article where date_format(from_unixtime(create_time),'%Y-%m-%d') = date_format(now(),'%Y-%m-%d')")
	if err != nil {
		log.Println(err)
		newcount = 0
	}

	spider, http := ProcessStatus()

	lastspider := 0

	if spider == 1 {
		b, err := redis.Bool((*rd).Do("EXISTS", "lastspider"))
		if err != nil {
			log.Println(err)
		} else if b {
			lastspider, err = redis.Int((*rd).Do("GET", "lastspider"))
			if err != nil {
				log.Println(err)
			}
		}
	}

	articles := GetNewArticles()

	sites := GetSites()

	resoutput := gin.H{
		"count":       count,
		"spidercount": spidercount,
		"webcount":    webcount,
		"usercount":   usercount,
		"newcount":    newcount,
		"spider":      spider,
		"http":        http,
		"lastspider":  lastspider,
		"articles":    articles,
		"sites":       sites,
	}
	SaveRedis(resoutput)
	c.JSON(200, resoutput)
}

func ProcessStatus() (int, int) {
	var spider int
	var http int
	rd := myredis.Connect()
	b, err := redis.Bool((*rd).Do("EXISTS", "ppid"))
	if err != nil {
		log.Println(err)
	} else if b {
		spider = 1
	} else {
		spider = 0
	}
	b, err = redis.Bool((*rd).Do("EXISTS", "hpid"))
	(*rd).Close()
	if err != nil {
		log.Println(err)
	} else if b {
		http = 1
	} else {
		http = 0
	}
	return spider, http
}

func GetNewArticles() []Oa {
	var articles []Oa
	err := Db.Select(&articles, `
	select a.title as title, a.url as url, c.name as site, FROM_UNIXTIME(a.create_time, '%Y/%m/%d %H:%i:%d') as time
	from dj_article a 
	join dj_spi b on a.spi_id=b.spi_id 
	join dj_spi_colname c on b.col0_id=c.col_id
	order by a.create_time desc limit 6
	`)
	if err != nil {
		log.Println(err)
	}
	return articles
}

func GetSites() []grap1 {
	sc := `
	select count(*) as count, c.name as name from dj_article a
	left join dj_spi b on a.spi_id=b.spi_id
	left join dj_spi_colname c on b.col0_id=c.col_id
	group by c.name;
	`
	var res []grap1
	err := Db.Select(&res, sc)
	if err != nil {
		log.Println(err)
	}
	return res
}

func GetRedis1() (string, bool) {
	res, err := redis.Bytes((*rd).Do("Get", "centre"))
	if err != nil {
		log.Println(err)
		return "", false
	}
	return string(res), true
}

func SaveRedis(resoutput gin.H) {
	res, err := json.Marshal(resoutput)
	if err != nil {
		log.Println(err)
	}

	_, err = (*rd).Do("Set", "centre", res)
	if err != nil {
		log.Println(err)
		log.Println("center-redis保存失败")
	} else {
		(*rd).Do("Expire", "centre", 200)
	}

}
