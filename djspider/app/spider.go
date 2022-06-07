/*
 * @Author: yfape
 * @Date: 2021-07-13 09:18:11
 * @LastEditTime: 2021-07-18 00:53:24
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider\app\action\action_spider.go
 */
package app

import (
	"djspider/db"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/golang-module/carbon"
)

type Rule struct {
	Main     string
	Sub      string
	Title    string
	Href     string
	Posttime string
}

type Spider struct {
	Output      chan []db.Article
	Spi         db.Spi
	Rule        Rule
	NewArticles []db.Article
}

func (spider *Spider) Set(mainchan chan []db.Article, spi db.Spi) {
	spider.Output = mainchan
	spider.Spi = spi
	err := json.Unmarshal([]byte(spider.Spi.Strategy), &(spider.Rule))
	if err != nil {
		panic(err)
	}
}

func (spider *Spider) Run() {

	switch spider.Spi.Mode {
	case 1:
		spider.Mode_1()
	}

	defer close(spider.Output)
	defer func() {
		if len(spider.NewArticles) > 0 {
			spider.Output <- spider.NewArticles
		}
		spider.NewArticles = []db.Article{}
		spider.Spi = db.Spi{}
		spider.Rule = Rule{}
	}()
}

func (spider *Spider) Mode_1() {
	c := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"), colly.MaxDepth(1))
	c.OnHTML(spider.Rule.Main, func(e *colly.HTMLElement) {
		e.ForEach(spider.Rule.Sub, func(i int, item *colly.HTMLElement) {
			var ctx db.Article
			ctx.Title = item.ChildText(spider.Rule.Title)
			ctx.Url = item.ChildAttr(spider.Rule.Href, "href")
			ctx.SpiId = spider.Spi.SpiId
			ctx.Site = spider.Spi.Col0Name
			if spider.Rule.Posttime != "" {
				ctx.Posttime = ToTS(item.ChildText(spider.Rule.Posttime))
			}
			ctx.CheckUrl(spider.Spi.Url)
			spider.NewArticles = append(spider.NewArticles, ctx)
		})

		spider.NewArticles = db.PickNewArticle(spider.NewArticles)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println(e)
	})
	err1 := c.Visit(spider.Spi.Url)
	if err1 != nil {
		fmt.Println(err1)
	}
}

func ToTS(str string) int {
	str = strings.Trim(str, " ")
	str = strings.Trim(str, "[")
	str = strings.Trim(str, "]")
	str = strings.Trim(str, "(")
	str = strings.Trim(str, ")")
	str = strings.Trim(str, " ")
	return int(carbon.Parse(str).ToTimestamp())
}
