package front

import (
	"blog/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"time"

	//"time"
)

var Category = [3]string{"杂谈", "技术", ""}

type IndexController struct {
	beego.Controller
}

//首页
func (this *IndexController) Index() {
	this.Data["title"] = beego.AppConfig.String("title")
	this.TplName = "front/index.html"
}

//文章列表
func (this *IndexController) BlogList() {
	log := logs.NewLogger()
	log.SetLogger(logs.AdapterConsole)

	//获取分页参数
	pagesize, _ := this.GetInt("pagesize") //每页条数
	if pagesize == 0 {
		pagesize = 5
	}
	pagenum, _ := this.GetInt("pagenum") //页码
	if pagenum == 0 {
		pagenum = 1
	}
	offset := pagesize * (pagenum - 1)
	limit := pagesize
	param := make(map[string]interface{})
	param["offset"] = offset
	param["limit"] = limit
	//获取首条文章
	bloglist := new(models.BlogList)
	total, data, err := bloglist.GetBlogList(param)
	fmt.Println(total)
	fmt.Println(err)
	//res := make([]*models.BlogList, 0)
	result := make([]map[string]interface{}, 0)
	if total != 0 && err == nil {
		for _, v := range data {
			t := make(map[string]interface{})
			Createtime := time.Unix(int64(v.Createtime), 0).Format(beego.AppConfig.String("timeFormat"))
			fmt.Println(v)
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

	this.Data["json"] = models.BlogListRes{Code: 0, Msg: "成功", Data: result}
	this.ServeJSON()
}

//文章详情
func (this *IndexController) GetBlog() {
	id,_ := this.GetInt("blog_id")

	blog := new(models.BlogList)
	blog.Id = id
	err := blog.GetBlog()
 var data []byte

   json.Unmarshal(data,&blog)
	fmt.Println(data)
	fmt.Println(err)

	this.Data["json"] = models.BlogListRes{Code: 0, Msg: "成功", Data: blog}
	this.ServeJSON()
}

//文章详情
func (this *IndexController) Details() {
	this.Data["blog_id"],_ = this.GetInt("blog_id")
	this.TplName = "front/details.html"
}

