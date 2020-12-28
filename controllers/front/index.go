package front

import (
	"blog/models"
	"github.com/astaxie/beego"
	//"reflect"
	"time"
)

const title = "鲨鱼辣椒的博客"
const timeFormat = "2006-01-01 15:03:09"

var Category = [3]string{"a","b","c"}


type IndexController struct {
	beego.Controller
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
	res := make(map[string]interface{})

	res["id"] = data.Id
	res["Introduction"] = data.Introduction
	t := time.Unix(int64(data.Createtime), 0).Format(timeFormat)
	res["Createtime"] = t
	res["Category_id"] = Category[data.Category_id]
	res["tag_id"] = data.Tag_id
	c.Data["json"] = res
	c.ServeJSON()
}

//文章列表
func (c *IndexController)BlogList(){


}
