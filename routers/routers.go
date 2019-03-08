package routers

import (
	"github.com/astaxie/beego"

	"controllers"
)

func init() {
	beego.Router("/",&controllers.MainController{})
	beego.Router("/assets",&controllers.Assetcontroller{},"get:DisAll")
	beego.Router("/assets",&controllers.Assetcontroller{},"post:Post")
	beego.Router("/d/:id",&controllers.Assetcontroller{},"*:DeletebyId")
	beego.Router("/:key",&controllers.Assetcontroller{},"get:GetByAny")
}
