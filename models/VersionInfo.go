package models

import "github.com/astaxie/beego/orm"

type VersionInfo struct {
	Id               int64
	Version          string `orm:size(100)`
	ArtifactId       int64
	UploadFileInfoId int64
}

func init() {
	orm.RegisterModel(new(VersionInfo))
}
