package front

import (
	"blog/models"
	"github.com/astaxie/beego"
)

const title  = "鲨鱼辣椒的博客"

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Test() {
	c.Data["Website"] = "boxue.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "front/index.html"
}

//首页
func (c *IndexController) Index() {
	c.Data["title"] = title
	c.TplName = "front/index.html"
}


//首页展示的文章
func (c *IndexController) FirstBlog() {
	//获取首条文章
	bloglist := new(models.BlogList)
	data := bloglist.GetBlogList()

	c.Data["json"] = data
	c.ServeJSON()
}