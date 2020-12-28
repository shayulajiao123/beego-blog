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
}

type BlogListRes struct {
	Code int
	Msg  string
	Data BlogList
}



func init() {
	orm.RegisterModel(new(BlogList))
}

func (this *BlogList) GetBlogList() BlogListRes {
	list := BlogList{Id: 1}
	o := orm.NewOrm()
	err := o.Read(&list)
	fmt.Println(err)
	if err == orm.ErrNoRows {
		return BlogListRes{1, "id不能为空", BlogList{}}
	} else if err == orm.ErrMissPK {
		return BlogListRes{1, "id不能为空", BlogList{}}
	} else {
		return BlogListRes{1, "id不能为空", BlogList{
			list.Id,
			list.Title,
			list.Introduction,
			list.Category_id, list.Createtime}}
	}
}
