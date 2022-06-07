/*
 * @Author: your name
 * @Date: 2021-07-23 10:04:23
 * @LastEditTime: 2021-08-03 13:09:16
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider\app\http\login.go
 */
package http

import (
	"djspider/db"
	"djspider/tool"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Input struct {
	Email string
	Pass  string
}

type InputUser struct {
	Token string
}

type UserCommon struct {
	UserId   int
	Name     string
	Email    string
	Headimg  string
	Identify string
}

func Login(c *gin.Context) {

	var input Input
	if err := c.BindJSON(&input); err != nil {
		c.JSON(420, gin.H{"msg": "邮箱或密码不能为空"})
		return
	}

	Db := db.MysqlConnect()
	var user db.User
	err := Db.Get(&user, "select * from dj_user where email=?", input.Email)
	Db.Close()
	if err != nil {
		fmt.Println(err)
		c.JSON(421, gin.H{"msg": "邮箱或密码错误"})
		return
	}

	if user.Pass != tool.MD5(input.Pass) {
		c.JSON(422, gin.H{"msg": "邮箱或密码错误"})
		return
	}

	var uc UserCommon
	uc.UserId = user.UserId
	uc.Name = user.Name
	uc.Email = user.Email
	uc.Headimg = user.Headimg
	uc.Identify = user.IdentifyName()

	token := tool.AuthCreate(uint(user.UserId))

	c.JSON(200, gin.H{"msg": "登录成功", "token": token, "user": uc})
}

func LoginUser(c *gin.Context) {
	var iu InputUser
	var user db.User
	if err := c.BindJSON(&iu); err != nil {
		c.JSON(420, gin.H{"msg": "参数错误"})
		return
	}
	b, userid := tool.AuthVertify(iu.Token)
	if !b {
		c.JSON(401, gin.H{"msg": "认证不通过"})
	} else {
		Db := db.MysqlConnect()
		err := Db.Get(&user, "select * from dj_user where user_id=?", userid)
		Db.Close()
		if err != nil {
			c.JSON(520, gin.H{"msg": "服务异常，请稍后再试"})
			log.Println(err)
		}
		var uc UserCommon
		uc.UserId = user.UserId
		uc.Name = user.Name
		uc.Email = user.Email
		uc.Headimg = user.Headimg
		uc.Identify = user.IdentifyName()
		c.JSON(200, gin.H{"msg": "认证成功", "user": uc})
	}
}
