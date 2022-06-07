/*
 * @Author: your name
 * @Date: 2021-07-13 09:18:11
 * @LastEditTime: 2021-07-16 13:32:05
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider\app\db\db_article.go
 */
package db

import (
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	ArticleId  int    `db:"article_id"`
	SpiId      int    `db:"spi_id"`
	Title      string `db:"title"`
	Url        string `db:"url"`
	Summary    string `db:"summary"`
	Img        string `db:"img"`
	Localimg   string `db:"localimg"`
	Posttime   int    `db:"post_time"`
	CreateTime int    `db:"create_time"`
	Site       string
}

func PickNewArticle(list []Article) []Article {
	var res []Article
	Db := MysqlConnect()
	for i := 0; i < len(list); i++ {
		a := list[i]
		if a.Title == "" || a.Url == "" || a.Title == "<nil>" || a.Title == "null" {
			continue
		}

		var checkarts []Article
		err := Db.Select(&checkarts, "select * from dj_article where title=? and spi_id=? order by create_time desc limit 20", a.Title, a.SpiId)
		if err != nil {
			panic(err)
		}
		if len(checkarts) <= 0 {
			res = append(res, a)
			_, err := Db.Exec("insert into dj_article(spi_id, title, url, post_time, create_time)values(?, ?, ?, ?, ?)", a.SpiId, a.Title, a.Url, a.Posttime, time.Now().Unix())
			if err != nil {
				panic(err)
			}
		}
		checkarts = []Article{}
	}
	defer MysqlClose(Db)
	return res
}

func (a *Article) CheckUrl(dom string) {
	if a.Url == "" || a.Url == "<nil>" || a.Url == "null" {
		return
	}
	if !strings.Contains(a.Url, "http") {
		if a.Url[1] == "."[0] || a.Url[0] == "/"[0] {
			tarr := strings.Split(dom, "/")
			tempdo := tarr[0] + "/" + tarr[1] + "/" + tarr[2]
			temp := strings.Trim(a.Url, ".")
			temp = strings.Trim(temp, ".")
			temp = strings.Trim(temp, "/")
			a.Url = tempdo + "/" + temp
		} else {
			tempdo := strings.Trim(dom, "/")
			temp := strings.Trim(a.Url, ".")
			temp = strings.Trim(temp, "/")
			a.Url = tempdo + "/" + temp
		}
	}
}
