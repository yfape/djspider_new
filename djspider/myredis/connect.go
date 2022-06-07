/*
 * @Author: your name
 * @Date: 2021-07-19 16:19:51
 * @LastEditTime: 2021-07-28 17:17:36
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider\redis\connect.go
 */
package myredis

import (
	"github.com/garyburd/redigo/redis"
)

func Connect() *redis.Conn {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	return &c
}
