package main

import (
	"github.com/astaxie/beego"
	_"assetStatistics/routers"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/plugins/auth"
//    "assetStatistics/controllers"
)

func init() {
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql","wumii:assetwm@tcp(127.0.0.1:3306)/asset?charset=utf8")
}

func main() {
	beego.InsertFilter("/v1/assets/d/*", beego.BeforeRouter,auth.Basic("wumii","147852369?><"))
//	beego.InsertFilter("/v1/assets/d/*", beego.BeforeRouter,controllers.FilterUser)
	beego.Run()
}
