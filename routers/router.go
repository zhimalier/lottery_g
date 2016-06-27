package routers

import (
	"lottery_g/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/lottery", &controllers.LotteryController{})
}
