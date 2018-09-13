package models

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/godfather1103/packageRepo/util"
	"strconv"
	"strings"
	"time"
)

type UploadFileInfo struct {
	Id              int64
	GroupId         string `orm:size(1000)`
	ArtifactId      string `orm:size(200)`
	Version         string `orm:size(100)`
	FileExt         string `orm:size(20)`
	FileName        string `orm:size(400)`
	FileMD5         string `orm:size(40)`
	LastVersionTime time.Time
}

/**
 *注册modal
 */
func init() {
	orm.RegisterModel(new(UploadFileInfo))
}

func FindUploadFileInfoById(Id int64) (*UploadFileInfo, error) {
	o := orm.NewOrm()
	uploadFileInfos := make([]*UploadFileInfo, 0)
	_, err := o.QueryTable("upload_file_info").Filter("id", Id).Limit(1).All(&uploadFileInfos)
	if err == nil && len(uploadFileInfos) > 0 {
		return uploadFileInfos[0], nil
	} else {
		if err != nil {
			return nil, err
		} else {
			return nil, errors.New("未查询到ID=" + strconv.FormatInt(Id, 10) + "的文件")
		}
	}
}

func AddUploadFileInfo(vo *UploadFileInfo) *UploadFileInfo {
	o := orm.NewOrm()
	item, error := FindUploadFileInfoByGAV(vo.GroupId, vo.ArtifactId, vo.Version)
	vo.LastVersionTime = time.Now()
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

func FindUploadFileInfoByGAV(groupId string, artifactId string, version string) (*UploadFileInfo, error) {
	o := orm.NewOrm()
	uploadFileInfos := make([]*UploadFileInfo, 0)
	_, err := o.QueryTable("upload_file_info").Filter("groupId", groupId).Filter("artifactId", artifactId).Filter("version", version).Limit(1).All(&uploadFileInfos)

	if err != nil {
		return nil, err
	} else {
		if len(uploadFileInfos) > 0 {
			return uploadFileInfos[0], nil
		} else {
			return nil, errors.New("未查询到相关的上传文件")
		}
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

func GetFileDownloadUrl(fileInfo *UploadFileInfo, useStreamUrl string) (string, error) {
	flag, msg := CheckUploadFileInfo(fileInfo)
	if flag {
		var url string
		if useStreamUrl != "1" {
			var groupId = fileInfo.GroupId
			var version = fileInfo.Version
			var fileName = fileInfo.FileName
			url = beego.AppConfig.String("webPrefixForUpload")
			url += "/" + strings.Replace(groupId, ".", "/", len(strings.Split(groupId, ".")))
			url += "/" + strings.Replace(version, ".", "/", len(strings.Split(version, ".")))
			url += "/" + fileName
		} else {
			url = beego.AppConfig.String("webStreamPrefixForUpload")
			url += "?FileId=" + strconv.FormatInt(fileInfo.Id, 10)
		}
		return url, nil
	} else {
		return "", errors.New(msg)
	}
}
