package routers

import (
	"bee_project/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	//beego.Router("/", &controllers.BookController{}, "get:GetAllBooks")
	beego.Router("/createBooks", &controllers.BookController{}, "post:CreateBooks")
}
