package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/godfather1103/packageRepo/models"
	"github.com/godfather1103/packageRepo/util"
	"io"
	"log"
	"os"
	"strings"
)

// UploadController operations for Upload
type UploadController struct {
	beego.Controller
}

func (c *UploadController) Post() {
	formFile, header, error := c.GetFile("file")
	if error != nil {
		log.Printf("文件获取失败：%s\n", error)
		resp := map[string]interface{}{
			"CODE": -1,
			"MSG":  "文件获取失败：" + error.Error(),
			"DATA": new(interface{}),
		}
		data, _ := json.Marshal(resp)
		c.Data["MSG"] = string(data)
		c.TplName = "toJson.tpl"
		return
	}
	var uploadFileInfo = new(models.UploadFileInfo)
	var groupId = c.GetString("groupId")
	uploadFileInfo.GroupId = groupId
	var artifactId = c.GetString("artifactId")
	uploadFileInfo.ArtifactId = artifactId
	var version = c.GetString("version")
	uploadFileInfo.Version = version
	var fileExt = c.GetString("fileExt")
	if util.CheckStrIsEmpty(fileExt) {
		fileExt = util.GetFileExt(header.Filename)
	}
	uploadFileInfo.FileExt = fileExt
	flag, info := models.CheckUploadFileInfo(uploadFileInfo)
	if !flag {
		log.Println(info)
		resp := map[string]interface{}{
			"CODE": -1,
			"MSG":  info,
			"DATA": new(interface{}),
		}
		data, _ := json.Marshal(resp)
		c.Data["MSG"] = string(data)
		c.TplName = "toJson.tpl"
		return
	}

	var descRootDir = util.GetOrDefault("uploadDir", "/repo")
	var descDir = descRootDir + "/" + strings.Replace(groupId, ".", "/", len(strings.Split(groupId, ".")))
	descDir += "/" + artifactId + "/" + version + "/"
	error = util.PathMkdir(descDir)
	if error != nil {
		log.Printf("创建路径失败：%s\n", error)
		resp := map[string]interface{}{
			"CODE": -1,
			"MSG":  "创建路径失败：" + error.Error(),
			"DATA": new(interface{}),
		}
		data, _ := json.Marshal(resp)
		c.Data["MSG"] = string(data)
		c.TplName = "toJson.tpl"
		return
	}
	uploadFileInfo.FileName = artifactId + "-" + version + "." + fileExt
	defer formFile.Close()
	destFile, error := os.Create(descDir + uploadFileInfo.FileName)
	if error != nil {
		log.Printf("创建文件失败：%s\n", error)
		resp := map[string]interface{}{
			"CODE": -1,
			"MSG":  "创建文件失败：" + error.Error(),
			"DATA": new(interface{}),
		}
		data, _ := json.Marshal(resp)
		c.Data["MSG"] = string(data)
		c.TplName = "toJson.tpl"
		return
	}
	defer destFile.Close()
	// 读取表单文件，写入保存文件
	_, error = io.Copy(destFile, formFile)
	if error != nil {
		log.Printf("文件写入失败：%s\n", error)
		resp := map[string]interface{}{
			"CODE": -1,
			"MSG":  "文件写入失败：" + error.Error(),
			"DATA": new(interface{}),
		}
		data, _ := json.Marshal(resp)
		c.Data["MSG"] = string(data)
		c.TplName = "toJson.tpl"
		return
	}
	md5File, _ := os.Open(descDir + uploadFileInfo.FileName)
	md5h := md5.New()
	defer md5File.Close()
	io.Copy(md5h, md5File)
	uploadFileInfo.FileMD5 = hex.EncodeToString(md5h.Sum([]byte("")))
	models.AddUploadFileInfo(uploadFileInfo)
	resp := map[string]interface{}{
		"CODE": 200,
		"MSG":  "上传成功！",
		"DATA": uploadFileInfo,
	}
	data, _ := json.Marshal(resp)
	c.Data["MSG"] = string(data)
	c.TplName = "toJson.tpl"
}
