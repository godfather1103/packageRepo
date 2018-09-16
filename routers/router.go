package routers

import (
	"github.com/astaxie/beego"
	"github.com/godfather1103/packageRepo/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/upload", &controllers.UploadController{}, "*:Upload")
	beego.Router("/uploadFile", &controllers.UploadController{}, "*:UploadFile")
	beego.Router("/getDownloadUrl", &controllers.DownloadController{})
	beego.Router("/getFileStream", &controllers.DownloadController{}, "*:GetFileStream")
	beego.Router("/getFileList", &controllers.DownloadController{}, "*:GetPathFileList")
}
