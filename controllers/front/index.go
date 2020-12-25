package front

import "github.com/astaxie/beego"

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Index() {
	c.Data["Website"] = "boxue.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "front/index.html"
}