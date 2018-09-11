package models

import "github.com/godfather1103/packageRepo/util"

type UploadFileInfo struct {
	GroupId    string
	ArtifactId string
	Version    string
	FileExt    string
	FileName   string
	FileMD5    string
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

type RespInfo struct {
	CODE int         `json:"code"`
	MSG  string      `json:"message"`
	DATA interface{} `json:"data"`
}
