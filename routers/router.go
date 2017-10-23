package routers

import (
	"assetStatistics/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/assets",
			beego.NSRouter("/", &controllers.AssetController{}, "post:Post"),
			beego.NSRouter("/", &controllers.AssetController{}, "put:Update"),
			//beego.NSRouter("/:name", &controllers.AssetController{}, "delete:Delete"),
			beego.NSRouter("/d/:name", &controllers.AssetController{}, "*:Delete"),
			beego.NSRouter("/all", &controllers.AssetController{}, "get:GetAll"),
			beego.NSRouter("/:key", &controllers.AssetController{}, "get:GetByNameOrMacOrIp"),
		),
	)
	beego.AddNamespace(ns)
}
