package routers

import (
	"app/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/event", &controllers.EventController{})
}
