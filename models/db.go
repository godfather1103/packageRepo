package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

const (
	_DB_Driver = "mysql"
)

func RegisterDB() {
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_schema := beego.AppConfig.String("db_schema")
	db_user := beego.AppConfig.String("db_user")
	db_passwd := beego.AppConfig.String("db_passwd")
	jdbcUrl := db_user + ":" + db_passwd + "@tcp(" + db_host + ":" + db_port + ")/" + db_schema + "?charset=utf8"
	beego.Info(fmt.Sprintf("connect to mysql server %v successfully !", db_host))
	orm.RegisterDriver(_DB_Driver, orm.DRMySQL)
	orm.RegisterDataBase("default", _DB_Driver, jdbcUrl, 30)

}
