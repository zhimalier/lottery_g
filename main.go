package main

import (
	_ "lottery_g/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetLogger("file", `{"filename":"d://lottery.log"}`)
	beego.Run()
}
