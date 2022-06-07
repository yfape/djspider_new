/*
 * @Author: your name
 * @Date: 2021-07-27 10:09:15
 * @LastEditTime: 2021-07-27 16:50:26
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider\app\http\article.go
 */
package http

import (
	"djspider/db"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ArticleFormat struct {
	Article_id  int
	Title       string
	Site        *string
	Url         string
	Post_time   string
	Create_time string
}

func GetArticles(c *gin.Context) {
	var article_list []ArticleFormat

	//获取页长
	row, err1 := strconv.Atoi(c.Param("row"))
	if err1 != nil {
		c.JSON(401, gin.H{"msg": "参数错误"})
		return
	}

	//获取页码
	page, err2 := strconv.Atoi(c.Param("page"))
	if err2 != nil {
		c.JSON(401, gin.H{"msg": "参数错误"})
		return
	}
	//获取搜索语句
	search := c.Param("search")
	search = strings.Trim(search, "/")
	ss := ""
	if search == "" {
		ss = "where true"
	} else {
		ss = `where title like '%` + search + `%'`
	}

	//连接Mysql
	Db = db.MysqlConnect()
	defer Db.Close()

	//获取最大页码
	var a1 int
	err := Db.Get(&a1, "select count(*) from dj_article "+ss)
	if err != nil {
		log.Println(err)
		c.JSON(522, gin.H{"msg": "获取失败，内部错误"})
		return
	}
	max := int(math.Ceil(float64(a1) / float64(row)))
	//页码不能溢出
	if page > max && max > 0 {
		page = max
	} else if max == 0 {
		page = 1
	}

	sc := `
		select 
		a.article_id as article_id,
		c.name as site,
		a.title as title,
		a.url as url,
		FROM_UNIXTIME(a.post_time, '%Y-%m-%d') as post_time,
		FROM_UNIXTIME(a.create_time, '%Y-%m-%d') as create_time 
		from dj_article a
		join dj_spi b on a.spi_id=b.spi_id
		left join dj_spi_colname c on b.col0_id=c.col_id
		` + ss + `
		order by a.Article_id desc
		limit ` + strconv.Itoa((page-1)*row) + `,` + strconv.Itoa(row) + `
	`

	err = Db.Select(&article_list, sc)
	if err != nil {
		log.Println(err)
		c.JSON(522, gin.H{"msg": "获取失败，内部错误"})
		return
	}

	if err != nil {
		log.Println(err)
		c.JSON(522, gin.H{"msg": "获取失败，内部错误"})
	} else {
		c.JSON(200, gin.H{"msg": "获取成功", "data": article_list, "max": max, "page": page})
	}

}
