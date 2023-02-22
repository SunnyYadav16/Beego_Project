package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type BookController struct {
	beego.Controller
}

type PersonController struct {
	beego.Controller
}

//func (c *MainController) Get() {
//	c.Data["Website"] = "beego.me"
//	c.Data["Email"] = "astaxie@gmail.com"
//	c.TplName = "index.tpl"
//}

//type Product struct {
//	Name  string
//	Price int
//}
//
//func (c *MainController) ShowData() {
//	c.Data["product"] = &Product{
//		Name:  "Macbook Pro",
//		Price: 2000,
//	}
//
//	mp := map[string]interface{}{
//		"Name": "Sunny",
//		"Age":  20,
//	}
//	c.Data["myMap"] = mp
//
//	languages := []string{"Go", "Python", "Java", "C++"}
//	c.Data["languages"] = languages
//
//	c.TplName = "hello.html"
//}

//func (book *BookController) GetAllBooks() {
//	name := book.GetString("book_name")
//	fmt.Println("name: ", name)
//	book.Data["name"] = name
//	book.TplName = "book_index.html"
//}

func (book *BookController) CreateBooks() {
	name := book.GetString("book_name")
	fmt.Println("((((((((((((")
	fmt.Println("name: ", name)
	book.Data["name"] = name
	book.TplName = "book_index.html"
}
