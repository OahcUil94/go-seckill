package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	redisPool *redis.Pool
)

func initRedis() (err error) {
	redisPool = &redis.Pool{
		MaxIdle: secKillConf.redisConf.redisMaxIdle,
		MaxActive: secKillConf.redisConf.redisMaxActive,
		IdleTimeout: time.Duration(secKillConf.redisConf.redisIdleTimeout) * time.Second, //
		Dial: func () (redis.Conn, error) {
			return redis.Dial("tcp", secKillConf.redisConf.redisAddr)
		},
	}

	conn := redisPool.Get()
	defer conn.Close()

	if _, err := conn.Do("ping"); err != nil {
		logs.Error("ping redis failed, err: %v", err)
		return
	}

	return
}

func initEtcd() (err error) {
	return
}

func initSec() (err error) {
	if err := initRedis(); err != nil {
		logs.Error("init redis failed, err:%v", err)
	}

	if err := initEtcd(); err != nil {
		logs.Error("init etcd failed, err:%v", err)
	}

	logs.Info("init sec success")

	return
}
