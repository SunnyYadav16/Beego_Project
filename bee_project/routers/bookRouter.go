package routers

import (
	"bee_project/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.BookController{})
	beego.Router("/showBooks", &controllers.BookController{}, "get:ShowAllBooks")
	beego.Router("/createBooks", &controllers.BookController{}, "get,post:CreateBooks")
	beego.Router("/updateBooks", &controllers.BookController{}, "get,post:UpdateBooks")
	beego.Router("/deleteBooks", &controllers.BookController{}, "post:DeleteBooks")
}
