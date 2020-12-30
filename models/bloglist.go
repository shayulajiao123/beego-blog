package models

import (
	"github.com/astaxie/beego/orm"
)

type BlogList struct {
	Id           int
	Title        string
	Introduction string
	Category_id  int
	Createtime   int64
	Tag_id       int
	Content      string
	//Createtime_string string
}

type BlogListRes struct {
	Code int
	Msg  string
	Data interface{}
}

func init() {
	orm.RegisterModel(new(BlogList))
}

//获取一条
func (c *BlogList) GetBlogList(params map[string]interface{}) (int64, []*BlogList, error) {
	o := orm.NewOrm()
	var list []*BlogList
	num, err := o.QueryTable("blog_list").Limit(params["limit"], params["offset"]).OrderBy("-id").All(&list)
	return num, list, err

}

//保存数据
func (c *BlogList) AddBlog() (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(c)
	if err == nil {
		return id, err
	}
	return 0, err
}
