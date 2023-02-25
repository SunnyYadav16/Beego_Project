package controllers

import (
	"bee_project/initializers"
	"bee_project/models"
	"bee_project/utils"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type PersonController struct {
	beego.Controller
}

func (managePerson *PersonController) CreatePerson() {
	if managePerson.Ctx.Input.IsPost() {
		var person models.Person

		userName := managePerson.GetString("username")
		userEmail := managePerson.GetString("email")
		userPassword := managePerson.GetString("password")

		userFound := initializers.DB.Where("email = ?", userEmail).First(&person)
		if userFound != nil {
			managePerson.Data["emailExist"] = "Email already exists"
			managePerson.TplName = "register.html"
			managePerson.Redirect("/register", 302)
		}

		hashPassword, pwdError := utils.HashPassword(userPassword)
		if pwdError != nil {
			fmt.Println(pwdError)
		}

		var user = models.Person{UserName: userName, Email: userEmail, Password: hashPassword}

		result := initializers.DB.Create(&user)
		if result.Error != nil {
			managePerson.Data["result"] = result.Error
		} else {
			managePerson.Data["result"] = "User created successfully"
		}

		managePerson.Redirect("/showBooks", 302)

	} else {
		managePerson.TplName = "register.html"
	}
}

func (managePerson *PersonController) LoginPerson() {
	if managePerson.Ctx.Input.IsPost() {
		var person models.Person

		userEmail := managePerson.GetString("email")
		userPassword := managePerson.GetString("password")
		fmt.Println(userEmail, userPassword)

		userFound := initializers.DB.Where("email = ?", userEmail).First(&person)
		if userFound == nil {
			//managePerson.Data["emailExist"] = "User with this email does not exist"
			//managePerson.TplName = "login.html"
			managePerson.Redirect("/login", 302)
		}

	} else {
		managePerson.TplName = "login.html"
	}
}
