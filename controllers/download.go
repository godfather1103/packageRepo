package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/godfather1103/packageRepo/models"
	"io/ioutil"
	"log"
	"strings"
)

// DownloadController operations for Download
type DownloadController struct {
	beego.Controller
}

func (c *DownloadController) GetPathFileList() {
	var pathFileId, err = c.GetInt64("PathFileId", int64(0))
	idType, err := c.GetInt64("IdType", int64(0))
	var returnType = c.GetString("returnType", "html")
	var resp map[string]interface{}
	var groupIdPaths = make([]*models.PathInfo, 0)
	var versionPaths = make([]*models.VersionInfo, 0)
	var fileInfos = make([]map[string]interface{}, 0)
	var error error
	if err == nil {
		if idType == 0 {
			// GroupId
			groupIdPaths, error = models.FindPathInfos(pathFileId)
			if error != nil {
				log.Printf("查询路径失败：%s\n", error)
				resp = map[string]interface{}{
					"CODE": -1,
					"MSG":  "查询路径失败：" + error.Error(),
					"DATA": new(interface{}),
				}
				data, _ := json.Marshal(resp)
				c.Data["MSG"] = string(data)
				c.TplName = "toJson.tpl"
				return
			}
		} else if idType == 1 {
			// ArtifactId
			versionPaths, error = models.FindVersonInfos(pathFileId)
			if error != nil {
				log.Printf("查询版本失败：%s\n", error)
				resp = map[string]interface{}{
					"CODE": -1,
					"MSG":  "查询版本失败：" + error.Error(),
					"DATA": new(interface{}),
				}
				data, _ := json.Marshal(resp)
				c.Data["MSG"] = string(data)
				c.TplName = "toJson.tpl"
				return
			}
		} else if idType == 2 {
			// Version
			fileInfos, error = models.FindUploadFileInfoByVersion(pathFileId)
			if error != nil {
				log.Printf("查询文件失败：%s\n", error)
				resp = map[string]interface{}{
					"CODE": -1,
					"MSG":  "查询文件失败：" + error.Error(),
					"DATA": new(interface{}),
				}
				data, _ := json.Marshal(resp)
				c.Data["MSG"] = string(data)
				c.TplName = "toJson.tpl"
				return
			}
		} else {
			log.Printf("传入的数据异常")
			resp = map[string]interface{}{
				"CODE": -1,
				"MSG":  "传入的数据异常",
				"DATA": new(interface{}),
			}
			data, _ := json.Marshal(resp)
			c.Data["MSG"] = string(data)
			c.TplName = "toJson.tpl"
			return
		}
		var DATA = make(map[string]interface{}, 4)
		DATA["GROUPIDPATHS"] = groupIdPaths
		DATA["VERSIONPATHS"] = versionPaths
		DATA["FILEINFOS"] = fileInfos
		DATA["TITLE"] = models.GetWebTitle(pathFileId, idType)
		var breadCrumb = models.GetBreadcrumb(pathFileId, idType)
		DATA["BREADCRUMB"] = breadCrumb
		if len(breadCrumb) > 1 {
			DATA["PARENTPATH"] = breadCrumb[len(breadCrumb)-2]
		} else {
			DATA["PARENTPATH"] = &models.PathAndVersion{
				PATHFILEID: 0,
				IDTYPE:     0,
				PATHNAME:   "",
			}
		}
		resp = map[string]interface{}{
			"CODE": 200,
			"MSG":  "数据获取成功！",
			"DATA": DATA,
		}
	} else {
		resp = map[string]interface{}{
			"CODE": -1,
			"MSG":  "未能获得相关文件：" + err.Error(),
			"DATA": new(interface{}),
		}
	}
	data, _ := json.Marshal(resp)
	c.Data["MSG"] = string(data)
	if returnType == "html" {
		c.Data["GROUPIDPATHS"] = groupIdPaths
		c.Data["VERSIONPATHS"] = versionPaths
		c.Data["FILEINFOS"] = fileInfos
		c.Data["TITLE"] = models.GetWebTitle(pathFileId, idType)
		var breadCrumb = models.GetBreadcrumb(pathFileId, idType)
		c.Data["BREADCRUMB"] = breadCrumb
		if len(breadCrumb) > 1 {
			c.Data["PARENTPATH"] = breadCrumb[len(breadCrumb)-2]
		} else {
			c.Data["PARENTPATH"] = &models.PathAndVersion{
				PATHFILEID: 0,
				IDTYPE:     0,
				PATHNAME:   "",
			}
		}
		c.TplName = "fileList.tpl"
	} else {
		c.TplName = "toJson.tpl"
	}
}

func (c *DownloadController) GetFileStream() {
	var fileId, err = c.GetInt64("FileId", int64(0))
	var resp map[string]interface{}
	if err == nil && fileId > 0 {
		uploadFileInfo, err := models.FindUploadFileInfoById(fileId)
		if err == nil {
			var groupId = uploadFileInfo.GroupId
			var artifactId = uploadFileInfo.ArtifactId
			var version = uploadFileInfo.Version
			var fileName = uploadFileInfo.FileName
			var descRootDir = beego.AppConfig.DefaultString("uploadDir", "/repo")
			var descDir = descRootDir + "/" + strings.Replace(groupId, ".", "/", len(strings.Split(groupId, ".")))
			descDir += "/" + strings.Replace(artifactId, ".", "/", len(strings.Split(artifactId, "."))) + "/" + version + "/"
			var fileLocalDir = descDir + fileName
			fileBytes, err := ioutil.ReadFile(fileLocalDir)
			if err == nil {
				c.Ctx.Output.Header("Content-Disposition", "filename="+uploadFileInfo.FileName)
				c.Ctx.Output.Body(fileBytes)
				return
			} else {
				resp = map[string]interface{}{
					"CODE": -1,
					"MSG":  "未能获得相关文件：" + err.Error(),
					"DATA": new(interface{}),
				}
			}
		} else {
			resp = map[string]interface{}{
				"CODE": -1,
				"MSG":  "未能获得相关文件：" + err.Error(),
				"DATA": new(interface{}),
			}
		}
	} else {
		resp = map[string]interface{}{
			"CODE": -1,
			"MSG":  "未能获得相关文件：" + err.Error(),
			"DATA": new(interface{}),
		}
	}
	data, _ := json.Marshal(resp)
	c.Data["MSG"] = string(data)
	c.TplName = "toJson.tpl"
}

func (c *DownloadController) Post() {
	var groupId = c.GetString("groupId")
	var artifactId = c.GetString("artifactId")
	var version = c.GetString("version")
	var useStreamUrl = c.GetString("useStreamUrl", "0")
	uploadFile, error := models.FindUploadFileInfoByGAV(groupId, artifactId, version)
	var resp map[string]interface{}
	if error == nil {
		url, error := models.GetFileDownloadUrl(uploadFile, useStreamUrl)
		if error != nil {
			log.Printf("未能获得相关文件：%s\n", error)
			resp = map[string]interface{}{
				"CODE": -1,
				"MSG":  "未能获得相关文件：" + error.Error(),
				"DATA": new(interface{}),
			}
		} else {
			resp = map[string]interface{}{
				"CODE": 200,
				"MSG":  "获取文件成功！",
				"DATA": url,
			}
		}
	} else {
		log.Printf("未能获得相关文件：%s\n", error)
		resp = map[string]interface{}{
			"CODE": -1,
			"MSG":  "未能获得相关文件：" + error.Error(),
			"DATA": new(interface{}),
		}
	}
	data, _ := json.Marshal(resp)
	c.Data["MSG"] = string(data)
	c.TplName = "toJson.tpl"
}

func (c *DownloadController) Get() {
	var groupId = c.GetString("groupId")
	var artifactId = c.GetString("artifactId")
	var version = c.GetString("version")
	var useStreamUrl = c.GetString("useStreamUrl", "0")
	uploadFile, error := models.FindUploadFileInfoByGAV(groupId, artifactId, version)
	var resp map[string]interface{}
	if error == nil {
		url, error := models.GetFileDownloadUrl(uploadFile, useStreamUrl)
		if error != nil {
			log.Printf("未能获得相关文件：%s\n", error)
			resp = map[string]interface{}{
				"CODE": -1,
				"MSG":  "未能获得相关文件：" + error.Error(),
				"DATA": new(interface{}),
			}
		} else {
			resp = map[string]interface{}{
				"CODE": 200,
				"MSG":  "获取文件成功！",
				"DATA": url,
			}
		}
	} else {
		log.Printf("未能获得相关文件：%s\n", error)
		resp = map[string]interface{}{
			"CODE": -1,
			"MSG":  "未能获得相关文件：" + error.Error(),
			"DATA": new(interface{}),
		}
	}
	data, _ := json.Marshal(resp)
	c.Data["MSG"] = string(data)
	c.TplName = "toJson.tpl"
}
