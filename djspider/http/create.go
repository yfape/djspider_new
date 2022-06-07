/*
 * @Author: your name
 * @Date: 2021-07-21 10:42:14
 * @LastEditTime: 2021-08-03 17:20:47
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider\app\http\create.go
 */
package http

import (
	"djspider/db"
	"djspider/tool"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func Create() {
	r := gin.Default()
	r.Use(Cors())
	v1 := r.Group("/v1")
	{
		v1.POST("/login", Login)
		v1.POST("/loginuser", LoginUser)
		v1.GET("/center", Center)
		v1.POST("/switchspider", SwitchSpider)
		v1.POST("/switchhttp", SwitchHttp)
		v1.POST("/runonce", RunOnce)
		v1.GET("/articles/:row/:page/*search", GetArticles)
		v1.GET("/cols", GetCols)
		v1.GET("/user/colsandkeywords", MiddleWare(), GetUserColsAndKeywords)
		v1.POST("/user/cols", MiddleWare(), SaveUserCols)
		v1.POST("/user/keyword", MiddleWare(), AddKeyword)
		v1.DELETE("/user/keyword/:keyword", MiddleWare(), RemoveKeyword)
		v1.POST("/user", MiddleWare(), UpdateUser)
		v1.GET("/users", AdminMiddleWare(), GetUsers)
		v1.POST("/admin/user", AdminMiddleWare(), AddUser)
		v1.PATCH("/admin/user", AdminMiddleWare(), SaveUser)
		v1.PATCH("/admin/user/pass", AdminMiddleWare(), ResetPass)
		v1.DELETE("/admin/user/:userid", AdminMiddleWare(), DeleteUser)
	}
	r.Static("/static", "./static")
	r.StaticFile("/favicon.ico", "./static/favicon.ico")
	r.Run(":8000")
}

//全局跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                              // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//				允许跨域设置																										可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //	跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //	处理请求
	}
}

//中间件
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		bo, user_id := tool.AuthVertify(token)
		if bo {
			c.Set("user_id", user_id)
			c.Next()
		} else {
			c.JSON(424, gin.H{"msg": "认证失败"})
		}
	}
}

func AdminMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		bo, user_id := tool.AuthVertify(token)
		if bo {
			Db = db.MysqlConnect()
			var user db.User
			err := Db.Get(&user, "select * from dj_user where user_id=?", user_id)
			if err != nil {
				c.JSON(421, gin.H{"msg": "不存在用户"})
			} else if user.Identify != 2 {
				c.JSON(425, gin.H{"msg": "权限不足"})
			} else {
				c.Set("user", user)
				c.Next()
			}
		} else {
			c.JSON(424, gin.H{"msg": "认证失败"})
		}
	}
}
