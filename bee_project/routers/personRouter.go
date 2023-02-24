package routers

import (
	"bee_project/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/register", &controllers.PersonController{}, "get,post:CreatePerson")
}
