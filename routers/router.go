package routers

import (
	"blog/controllers"
	"blog/controllers/front"
	"github.com/astaxie/beego"

)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/index/index", &front.IndexController{},"get:Index")//test
	beego.Router("/index/first-blog", &front.IndexController{},"get:FirstBlog")//test

}
