package controllers

import (
	"bee_project/initializers"
	"bee_project/models"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
)

type BookController struct {
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

func (book *BookController) ShowAllBooks() {
	var books []models.Book

	initializers.DB.Find(&books)

	book.Data["result"] = books
	book.TplName = "book_index.html"
}

func (manageBook *BookController) CreateBooks() {
	if manageBook.Ctx.Input.IsPost() {
		bookTitle := manageBook.GetString("book_title")
		bookAuthor := manageBook.GetString("book_author")
		book := models.Book{Title: bookTitle, Author: bookAuthor}

		result := initializers.DB.Create(&book)
		if result.Error != nil {
			manageBook.Data["result"] = result.Error
		}

		manageBook.Redirect("/showBooks", 302)
	} else {
		manageBook.TplName = "create_books.html"
	}
}

func (manageBook *BookController) DeleteBooks() {
	if manageBook.Ctx.Input.IsPost() {
		bookId, _ := strconv.Atoi(manageBook.GetString("id"))

		var book models.Book

		initializers.DB.Where("id = ?", bookId).Delete(&book)
		manageBook.Redirect("/showBooks", 302)
	}
}

func (manageBook *BookController) UpdateBooks() {
	if manageBook.Ctx.Input.IsPost() {
		bookId, _ := strconv.Atoi(manageBook.GetString("id"))

		bookTitle := manageBook.GetString("book_title")
		bookAuthor := manageBook.GetString("book_author")

		result := initializers.DB.Model(&models.Book{}).Where("id = ?",
			bookId).Updates(map[string]interface{}{"title": bookTitle, "author": bookAuthor})

		if result.Error != nil {
			manageBook.Data["result"] = result.Error
		} else {
			manageBook.Redirect("/showBooks", 302)
		}
		manageBook.TplName = "update_books.html"
	} else {
		bookId, _ := strconv.Atoi(manageBook.GetString("id"))

		var book models.Book

		initializers.DB.Where("id = ?", bookId).First(&book)

		manageBook.Data["books"] = book
		manageBook.TplName = "update_books.html"
	}
}
