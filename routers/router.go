package routers

import (
	"github.com/godfather1103/packageRepo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
