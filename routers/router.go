package routers

import (
	"github.com/astaxie/beego"
	"github.com/godfather1103/packageRepo/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/uploadFile", &controllers.UploadController{})
	beego.Router("/downloadFile", &controllers.DownloadController{})
}
