package routers

import (
	"blog/controllers/front"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/front/index", &front.IndexController{},"*:Index")
}
