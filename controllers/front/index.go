package front

import (
	"blog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"time"

	//"time"
)

const title = "鲨鱼辣椒的博客"
const timeFormat = "2006-01-01 15:03:09"

var Category = [3]string{"a", "b", "c"}

type IndexController struct {
	beego.Controller
}

//首页
func (this *IndexController) Index() {
	this.Data["title"] = title
	this.TplName = "front/index.html"
}

//文章列表
func (this *IndexController) BlogList() {
	log := logs.NewLogger()
	log.SetLogger(logs.AdapterConsole)

	//获取分页参数
	pagesize, _ := this.GetInt("pagesize") //每页条数
	if pagesize == 0 {
		pagesize = 10
	}
	pagenum, _ := this.GetInt("pagenum") //页码
	if pagenum == 0 {
		pagenum = 3
	}
	offset := pagesize * (pagenum - 1)
	limit := pagesize
	param := make(map[string]interface{})
	param["offset"] = offset
	param["limit"] = limit
	//获取首条文章
	bloglist := new(models.BlogList)
	total, data, err := bloglist.GetBlogList(param)
	//res := make([]*models.BlogList, 0)
	result := make([]map[string]interface{}, 0)
	if total != 0 && err == nil {
		for _, v := range data {
			t := make(map[string]interface{})
			Createtime := time.Unix(int64(v.Createtime), 0).Format(timeFormat)
			t["id"] = v.Id
			t["Introduction"] = v.Introduction
			t["Category_id"] = Category[v.Category_id]
			t["Createtime"] = Createtime
			t["Content"] = v.Content
			t["Tag_id"] = v.Tag_id
			t["Title"] = v.Title

			result = append(result, t)
		}
	}

	this.Data["json"] = result
	this.ServeJSON()
}
