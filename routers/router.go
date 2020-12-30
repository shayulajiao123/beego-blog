package routers

import (
	"blog/controllers"
	"blog/controllers/backend"
	"blog/controllers/front"
	"github.com/astaxie/beego"

)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/index/index", &front.IndexController{},"get:Index")//test
	beego.Router("/index/bloglist", &front.IndexController{},"post:BlogList")//test

	beego.Router("/index/addblogpage", &backend.IndexControler{},"get:AddBlogPage")//test
	beego.Router("/index/addblog", &backend.IndexControler{},"post:AddBlog")//test


}
