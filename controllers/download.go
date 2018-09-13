package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/godfather1103/packageRepo/models"
	"github.com/godfather1103/packageRepo/util"
	"io/ioutil"
	"log"
	"strings"
)

// DownloadController operations for Download
type DownloadController struct {
	beego.Controller
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
			var descRootDir = util.GetOrDefault("uploadDir", "/repo")
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
