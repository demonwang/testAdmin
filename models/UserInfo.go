package models

import "github.com/astaxie/beego/orm"

//分组表
type UserInfo struct {
	Id     int64
	Name   string `orm:"size(100)" form:"Name"  valid:"Reqiured"`
	Title  string `orm:"size(100)" form:"Title"  valid:"Required"`
	Status int    `orm:"default(2)" form:"Status" valid:"Range(1,2)"`
	Sort   int    `orm:"default(1)" form:"Sort" valid:"Numeric"`
}

func (this *UserInfo) TableName() string {
	return "user_info"
}

func init() {
	orm.RegisterModel(new(UserInfo))
}
