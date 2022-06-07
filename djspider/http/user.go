/*
 * @Author: your name
 * @Date: 2021-07-30 16:47:26
 * @LastEditTime: 2021-08-03 17:21:43
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider\http\user.go
 */
package http

import (
	"djspider/db"
	"djspider/tool"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUserColsAndKeywords(c *gin.Context) {
	user_id, bo := c.Get("user_id")
	if !bo {
		c.JSON(421, gin.H{"msg": "参数错误"})
		return
	}
	Db = db.MysqlConnect()
	defer Db.Close()
	var user db.User
	err := Db.Get(&user, "select * from dj_user where user_id=?", user_id)
	if err != nil {
		c.JSON(421, gin.H{"msg": "不存在用户"})
		return
	}
	c.JSON(200, gin.H{"msg": "获取成功", "cols": user.Col_ids, "keywords": user.Keywords})
}

type I1 struct {
	Cols []string `json:"cols" binding:"required"`
}

func SaveUserCols(c *gin.Context) {
	user_id, bo := c.Get("user_id")
	if !bo || user_id == nil {
		c.JSON(421, gin.H{"msg": "认证参数错误"})
		return
	}
	var input I1
	if err := c.BindJSON(&input); err != nil {
		c.JSON(421, gin.H{"msg": "请求参数错误"})
		return
	}
	cols := strings.Join(input.Cols, ",")
	Db = db.MysqlConnect()
	defer Db.Close()
	//为空
	if len(input.Cols) == 0 {
		_, err := Db.Exec(`update dj_user set spi_ids=?,col_ids=? where user_id=?`, "", "", user_id)
		if err != nil {
			log.Println(err)
			c.JSON(521, gin.H{"msg": "内部错误"})
			return
		} else {
			c.JSON(200, gin.H{"msg": "修改成功"})
			return
		}
	}
	//非空
	var spi_ids []string
	err := Db.Select(&spi_ids, `
		select spi_id from dj_spi where 
		col0_id in(`+cols+`) or
		col1_id in(`+cols+`) or
		col2_id in(`+cols+`) or
		col3_id in(`+cols+`)
		order by spi_id asc
	`)
	if err != nil {
		log.Println(err)
		c.JSON(521, gin.H{"msg": "内部错误"})
		return
	}
	is := strings.Join(spi_ids, ",")
	_, err = Db.Exec(`update dj_user set spi_ids=?,col_ids=? where user_id=?`, is, cols, user_id)
	if err != nil {
		log.Println(err)
		c.JSON(521, gin.H{"msg": "内部错误"})
		return
	} else {
		c.JSON(200, gin.H{"msg": "修改成功"})
	}
}

type I2 struct {
	Keyword string
}

func AddKeyword(c *gin.Context) {
	user_id, bo := c.Get("user_id")
	if !bo || user_id == nil {
		c.JSON(421, gin.H{"msg": "认证参数错误"})
		return
	}
	var input I2
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(421, gin.H{"msg": "请求参数错误"})
		return
	}
	keyword := input.Keyword

	var user db.User
	Db = db.MysqlConnect()
	defer Db.Close()
	err = Db.Get(&user, "select * from dj_user where user_id=?", user_id)
	if err != nil {
		c.JSON(421, gin.H{"msg": "认证参数错误"})
		return
	}
	if len(keyword) > 20 {
		c.JSON(424, gin.H{"msg": "关键字过长"})
		return
	}
	if strings.Contains(","+user.Keywords+",", ","+keyword+",") {
		c.JSON(423, gin.H{"msg": "关键字已存在"})
		return
	}
	uks := strings.Split(user.Keywords, ",")
	uks = append(uks, keyword)
	im := strings.Join(uks, ",")

	_, err = Db.Exec("update dj_user set keywords=? where user_id=?", im, user_id)
	if err != nil {
		log.Println(err)
		c.JSON(521, gin.H{"msg": "内部错误"})
		return
	}
	c.JSON(200, gin.H{"msg": "新增成功", "keywords": uks})
}

func RemoveKeyword(c *gin.Context) {
	user_id, bo := c.Get("user_id")
	if !bo || user_id == nil {
		c.JSON(421, gin.H{"msg": "认证参数错误"})
		return
	}
	keyword := c.Param("keyword")
	var user db.User
	Db = db.MysqlConnect()
	defer Db.Close()
	err := Db.Get(&user, "select * from dj_user where user_id=?", user_id)
	if err != nil {
		c.JSON(421, gin.H{"msg": "认证参数错误"})
		return
	}
	if !strings.Contains(user.Keywords, keyword) {
		c.JSON(425, gin.H{"msg": "不存在关键字"})
		return
	}

	tk := strings.Replace(","+user.Keywords+"", ","+keyword, "", -1)
	tk = strings.Trim(tk, ",")
	_, err = Db.Exec("update dj_user set keywords=? where user_id=?", tk, user_id)
	if err != nil {
		log.Println(err)
		c.JSON(521, gin.H{"msg": "内部错误"})
		return
	}
	c.JSON(200, gin.H{"msg": "删除成功", "keywords": strings.Split(tk, ",")})
}

func UpdateUser(c *gin.Context) {
	user_id, bo := c.Get("user_id")
	if !bo || user_id == nil {
		c.JSON(421, gin.H{"msg": "认证参数错误"})
		return
	}
	var iu db.User
	err := c.BindJSON(&iu)
	if err != nil {
		c.JSON(421, gin.H{"msg": "请求参数错误", "data": err})
		return
	}

	Db = db.MysqlConnect()
	defer Db.Close()
	if iu.Pass == "" {
		_, err = Db.Exec("update dj_user set headimg=?,name=?,email=? where user_id=?", iu.Headimg, iu.Name, iu.Email, user_id)
	} else {
		npass := tool.MD5(iu.Pass)
		_, err = Db.Exec("update dj_user set pass=?,headimg=?,name=?,email=? where user_id=?", npass, iu.Headimg, iu.Name, iu.Email, user_id)
	}
	if err != nil {
		c.JSON(421, gin.H{"msg": "请求参数错误", "data": err})
	} else {
		c.JSON(200, gin.H{"msg": "修改成功"})
	}
}

func GetUsers(c *gin.Context) {
	defer Db.Close()
	var users []db.User
	err := Db.Select(&users, "select user_id,name,email,headimg,identify,create_time from dj_user order by user_id asc")
	if err != nil {
		log.Println(err)
		c.JSON(521, gin.H{"msg": "内部错误"})
		return
	}
	c.JSON(200, gin.H{"msg": "获取成功", "users": users})
}

func AddUser(c *gin.Context) {
	defer Db.Close()
	var user db.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(421, gin.H{"msg": "参数错误"})
		return
	}
	if user.Identify != 2 {
		user.Identify = 1
	}
	user.Pass = tool.MD5("123")
	_, err = Db.Exec("insert into dj_user (name,email,pass,headimg,identify,create_time) value(?,?,?,?,?,?)", user.Name, user.Email, user.Pass, user.Headimg, user.Identify, time.Now().Unix())
	if err != nil {
		c.JSON(421, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"msg": "新增成功"})
}

func SaveUser(c *gin.Context) {
	defer Db.Close()
	var user db.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(421, gin.H{"msg": "参数错误"})
		return
	}
	if user.Identify != 2 {
		user.Identify = 1
	}
	_, err = Db.Exec(`update dj_user set name=?,email=?,headimg=?,identify=? where user_id=?`, user.Name, user.Email, user.Headimg, user.Identify, user.UserId)
	if err != nil {
		c.JSON(421, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"msg": "保存成功"})
}

func ResetPass(c *gin.Context) {
	defer Db.Close()
	var user db.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(421, gin.H{"msg": "参数错误"})
		return
	}
	user.Pass = tool.MD5("123")
	_, err = Db.Exec(`update dj_user set pass=? where user_id=?`, user.Pass, user.UserId)
	if err != nil {
		c.JSON(421, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"msg": "重置密码成功"})
}

func DeleteUser(c *gin.Context) {
	defer Db.Close()
	UserId, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		c.JSON(421, gin.H{"msg": "参数错误"})
		return
	}
	var user db.User
	err = Db.Get(&user, "select * from dj_user where user_id=?", UserId)
	if err != nil {
		c.JSON(421, gin.H{"msg": "请求参数错误"})
		return
	}
	iu, bo := c.Get("user")
	if !bo {
		c.JSON(422, gin.H{"msg": "认证错误"})
		return
	}
	selfuser := iu.(db.User)
	if user.UserId == selfuser.UserId {
		c.JSON(522, gin.H{"msg": "不能删除本账号"})
		return
	}
	_, err = Db.Exec("delete from dj_user where user_id=?", UserId)
	if err != nil {
		c.JSON(525, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"msg": "删除成功"})
}
