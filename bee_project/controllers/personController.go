package controllers

import beego "github.com/beego/beego/v2/server/web"

type PersonController struct {
	beego.Controller
}

func (managePerson *PersonController) CreatePerson() {
	if managePerson.Ctx.Input.IsPost() {
		//userName := manage.
	} else {
		managePerson.TplName = "register.html"
	}
}
