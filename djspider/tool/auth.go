/*
 * @Author: your name
 * @Date: 2021-07-23 14:53:04
 * @LastEditTime: 2021-07-23 17:27:59
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider\tool\auth.go
 */
package tool

import (
	"djspider/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtkey = []byte(config.Auth.Key)
var str string

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func AuthCreate(userid uint) string {
	expireTime := time.Now().Add(time.Hour * time.Duration(24*config.Auth.ExpireDay))
	claims := &Claims{
		UserId: userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    config.Auth.Issue,   // 签名颁发者
			Subject:   config.Auth.Subject, //签名主题
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		panic(err)
	}
	return tokenString
}

func AuthVertify(tokenString string) (bool, int) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	if err != nil || !token.Valid {
		return false, 0
	}
	return true, int(Claims.UserId)
}
