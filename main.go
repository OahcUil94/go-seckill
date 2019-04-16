package main

import (
	"github.com/astaxie/beego"

	_ "github.com/OahcUil94/go-seckill/route"
)

func main() {
	if err := initConfig(); err != nil {
		panic(err)
	}
	beego.Run()
}

