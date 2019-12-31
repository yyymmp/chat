package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

//初始化reds
var pool *redis.Pool

func initPool(address string, maxIdle int, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", address)
		},
		TestOnBorrow:    nil,
		MaxIdle:         maxIdle,     //最大空闲连接数
		MaxActive:       maxActive,   //0 表示没限制最大连接数量
		IdleTimeout:     idleTimeout, //连接多长时间后关闭连接 重新放回池
		Wait:            false,
		MaxConnLifetime: 0, //连接使用最大时长
	}
}
