package models

import "github.com/astaxie/beego/orm"

type PathInfo struct {
	Id           int64
	PathName     string `orm:size(1000)`
	PathType     int64
	ParentPathId int64
}

func init() {
	orm.RegisterModel(new(PathInfo))
}
