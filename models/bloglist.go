package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type BlogList struct {
	Id           int
	Title        string
	Introduction string
	Category_id  int
	Createtime   int
	Tag_id       int
}

type BlogListRes struct {
	Code int
	Msg  string
	Data BlogList
}

func init() {
	orm.RegisterModel(new(BlogList))
}

func (this *BlogList) GetBlogList() BlogList {
	o := orm.NewOrm()
	var list BlogList
	//err := o.Read(&list)
	err := o.QueryTable("blog_list").OrderBy("-id").One(&list)
	fmt.Println(err)
	if err == orm.ErrNoRows {
		return list
	} else {
		return list
	}
}
