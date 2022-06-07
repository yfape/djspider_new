/*
 * @Author: your name
 * @Date: 2021-07-15 10:15:50
 * @LastEditTime: 2021-08-02 10:29:10
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider\app\email.go
 */
package app

import (
	"djspider/config"
	"djspider/db"
	"log"
	"net/smtp"
	"strconv"
	"strings"

	"github.com/jordan-wright/email"
)

type Email struct {
	User     db.User
	Articles *[]db.Article
	Sendlist []byte
}

func (ema *Email) Set(user db.User, articles *[]db.Article) {
	ema.User = user
	ema.Articles = articles
	ema.PickSendlist()
}

func (ema *Email) Send() {

	if ema.User.Keywords == "" && ema.User.Spi_ids == "" {
		return
	}
	if len(ema.Sendlist) <= 0 {
		return
	}

	e := email.NewEmail()
	e.From = config.Email.From
	// 设置接收方的邮箱
	e.To = []string{ema.User.Email}
	//设置主题
	e.Subject = config.Email.Subject
	//设置文件发送的内容
	e.HTML = []byte(ema.Sendlist)
	//设置服务器相关的配置
	log.Println("发送：", ema.User.Email)
	// mjuborvxxfwwbjii
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", config.Email.Username, config.Email.Authcode, "smtp.qq.com"))
	if err != nil {
		panic(err)
	}

	ema.User = db.User{}
	ema.Sendlist = []byte{}
}

func (ema *Email) PickSendlist() {
	for i := 0; i < len(*(ema.Articles)); i++ {
		if strings.Contains(","+ema.User.Spi_ids+",", ","+strconv.Itoa((*ema.Articles)[i].SpiId)+",") {
			text := SetHtml((*ema.Articles)[i].Title, (*ema.Articles)[i].Url, (*ema.Articles)[i].Site)
			ema.Sendlist = append(ema.Sendlist, []byte(text)...)
		} else if CheckKeywords(ema.User.Keywords, (*ema.Articles)[i].Title) {
			text := SetHtml((*ema.Articles)[i].Title, (*ema.Articles)[i].Url, (*ema.Articles)[i].Site)
			ema.Sendlist = append(ema.Sendlist, []byte(text)...)
		}
	}
}

func CheckKeywords(keywords string, title string) bool {
	res := false
	if keywords == "" || keywords == "null" || keywords == "<nil>" {
		return false
	}

	karr := strings.Split(keywords, ",")
	for i := 0; i < len(karr); i++ {
		if strings.Contains(title, karr[i]) {
			res = true
			break
		}
	}
	return res
}

func SetHtml(title string, url string, site string) string {
	text := `<div><a target='_blank' style='text-decoration:none;color:#000' href="` + url + `"><div style='padding:5px 10px;background:#ECE6E6;border-radius:5px;margin-bottom:8px'>
		<div style='font-size:18px;'>` + title + `</div>
		<div style='text-align:right;color:#A19999'>` + site + `</div>
	</div></a></div>`
	return text
}
