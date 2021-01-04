package backend

import (
	"blog/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"time"
)

type IndexControler struct {
	beego.Controller
}

//新增文章功能页面
func (this *IndexControler) AddBlogPage() {
	this.Data["title"] = beego.AppConfig.String("title")
	this.TplName = "backend/add.html"
}

//新增文章
func (this *IndexControler) AddBlog() {
	data := this.GetString("content")
	input := []byte(data)
	unsafe := blackfriday.MarkdownCommon(input)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	fmt.Println(string(html))

	blog := new(models.BlogList)
	blog.Category_id = 1
	blog.Title = "test"
	blog.Tag_id = 1
	blog.Introduction ="asdf"
	//content,_:= json.Marshal(html)
	blog.Content = string(html)
	blog.Createtime = time.Now().Unix()
	num, err := blog.AddBlog()
	fmt.Println(num)
	fmt.Println(err)
	this.Data["json"] = data
	this.ServeJSON()
}
