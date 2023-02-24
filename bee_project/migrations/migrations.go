package main

import (
	"bee_project/initializers"
	"bee_project/models"
	"fmt"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	fmt.Println("Migrating database...")
	initializers.DB.AutoMigrate(&models.Book{}, &models.Person{}, &models.BookManagement{})
}
