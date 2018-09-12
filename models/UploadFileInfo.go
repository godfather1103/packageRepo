package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/godfather1103/packageRepo/util"
)

type UploadFileInfo struct {
	Id         int64
	GroupId    string `orm:size(1000)`
	ArtifactId string `orm:size(200)`
	Version    string `orm:size(100)`
	FileExt    string `orm:size(20)`
	FileName   string `orm:size(400)`
	FileMD5    string `orm:size(40)`
}

/**
 *注册modal
 */
func init() {
	orm.RegisterModel(new(UploadFileInfo))
}

func AddUploadFileInfo(vo *UploadFileInfo) *UploadFileInfo {
	o := orm.NewOrm()
	item, error := FindUploadFileInfoByGA(vo.GroupId, vo.ArtifactId)
	if error == nil && item != nil {
		vo.Id = item.Id
		_, error := o.Update(vo)
		if error == nil {
			artifactId, error := AddPathInfo(vo.GroupId, vo.ArtifactId)
			if error == nil {
				_, error = AddVersionInfo(vo.Version, artifactId, vo.Id)
				if error == nil {
					return vo
				} else {
					beego.Error("更新版本数据失败，error=" + error.Error())
					return nil
				}
			} else {
				beego.Error("更新路径数据失败，error=" + error.Error())
				return nil
			}
		} else {
			beego.Error("更新文件数据失败，error=" + error.Error())
			return nil
		}
	}
	id, error := o.Insert(vo)
	if error == nil {
		vo.Id = id
		artifactId, error := AddPathInfo(vo.GroupId, vo.ArtifactId)
		if error == nil {
			_, error = AddVersionInfo(vo.Version, artifactId, vo.Id)
			if error == nil {
				return vo
			} else {
				beego.Error("更新版本数据失败，error=" + error.Error())
				return nil
			}
		} else {
			beego.Error("更新路径数据失败，error=" + error.Error())
			return nil
		}
	} else {
		beego.Error("新建文件数据失败，error=" + error.Error())
		return nil
	}
}

func FindUploadFileInfoByGA(groupId string, artifactId string) (*UploadFileInfo, error) {
	o := orm.NewOrm()
	uploadFileInfos := make([]*UploadFileInfo, 0)
	_, err := o.QueryTable("upload_file_info").Filter("groupId", groupId).Filter("artifactId", artifactId).Limit(1).All(&uploadFileInfos)
	if len(uploadFileInfos) > 0 {
		return uploadFileInfos[0], err
	} else {
		return nil, err
	}
}

func CheckUploadFileInfo(fileInfo *UploadFileInfo) (bool, string) {
	if util.CheckStrIsEmpty(fileInfo.GroupId) {
		return false, "GroupId字段不能为空！"
	} else if util.CheckStrIsEmpty(fileInfo.ArtifactId) {
		return false, "ArtifactId字段不能为空！"
	} else if util.CheckStrIsEmpty(fileInfo.Version) {
		return false, "Version字段不能为空！"
	} else if util.CheckStrIsEmpty(fileInfo.FileExt) {
		return false, "无法自动获取文件后缀名，请手动通过FileExt传递！"
	} else {
		return true, ""
	}

}
