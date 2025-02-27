package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/godfather1103/packageRepo/models"
	"github.com/godfather1103/utils"
	"io"
	"log"
	"os"
	"strings"
)

// UploadController operations for Upload
type UploadController struct {
	beego.Controller
}

func (c *UploadController) Upload() {
	c.TplName = "upload.tpl"
}

func (c *UploadController) UploadFile() {
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
	if utils.CheckStrIsEmpty(fileExt) {
		fileExt = utils.GetFileExt(header.Filename)
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

	var descRootDir = beego.AppConfig.DefaultString("uploadDir", "/repo")
	var descDir = descRootDir + "/" + strings.Replace(groupId, ".", "/", len(strings.Split(groupId, ".")))
	descDir += "/" + artifactId + "/" + version + "/"
	error = utils.PathMkdir(descDir)
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
	uploadFileInfo, error = models.AddUploadFileInfo(uploadFileInfo)
	if error != nil {
		log.Printf("文件写入失败：%s\n", error)
		resp := map[string]interface{}{
			"CODE": -1,
			"MSG":  error.Error(),
			"DATA": new(interface{}),
		}
		data, _ := json.Marshal(resp)
		c.Data["MSG"] = string(data)
		c.TplName = "toJson.tpl"
		return
	}
	var fileInfoMap = map[string]interface{}{}
	fileInfo, _ := json.Marshal(uploadFileInfo)
	json.Unmarshal(fileInfo, &fileInfoMap)
	fileDownloadUrl, _ := models.GetFileDownloadUrl(uploadFileInfo, "1")
	fileInfoMap["STREAMURL"] = fileDownloadUrl
	fileDownloadUrl, _ = models.GetFileDownloadUrl(uploadFileInfo, "0")
	fileInfoMap["WEBURL"] = fileDownloadUrl
	resp := map[string]interface{}{
		"CODE": 200,
		"MSG":  "上传成功！",
		"DATA": fileInfoMap,
	}
	data, _ := json.Marshal(resp)
	c.Data["MSG"] = string(data)
	c.TplName = "toJson.tpl"
}
