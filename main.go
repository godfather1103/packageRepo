package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/godfather1103/packageRepo/models"
	_ "github.com/godfather1103/packageRepo/routers"
)

func init() {
	models.RegisterDB()
	orm.RunSyncdb("default", false, false)
}

func main() {
	beego.SetStaticPath("/myRepo", beego.AppConfig.String("uploadDir"))
	beego.Run()
}
