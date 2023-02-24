package models

import (
	"gorm.io/gorm"
	"time"
)

// USE THIS WHEN YOU WANT TO USE BEEGO ORM

//type Person struct {
//	Id     int    `orm:"column(id);auto"`
//	Name   string `orm:"column(name);size(255)"`
//	Mobile string `orm:"column(mobile);size(255)"`
//	Email  string `orm:"column(email);size(255)"`
//}
//
//type Books struct {
//	Id     int     `orm:"column(id);auto"`
//	Name   string  `orm:"column(name);size(255)"`
//	Author string  `orm:"column(author);size(255)"`
//	Price  int     `orm:"column(price)"`
//	Person *Person `orm:"rel(fk)"`
//}

// USE THIS WHEN YOU WANT TO USE GORM

type Person struct {
	gorm.Model
	UserName string `gorm:"size:255;not null"`
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
}

type Book struct {
	gorm.Model
	Title       string `gorm:"size:255;not null"`
	Author      string `gorm:"size:255;not null"`
	Price       int    `gorm:"not null"`
	Quantity    int    `gorm:"not null"`
	Publication int    `gorm:"not null"`
	Genre       string `gorm:"size:255;not null"`
	Pages       int    `gorm:"not null"`
}

type BookManagement struct {
	gorm.Model
	IssuedDate   *time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	ReturnedDate *time.Time `gorm:"not null"`
	PersonID     uint       `gorm:"not null"`
	BookID       uint       `gorm:"not null"`
	Person       Person     `gorm:"foreignKey:PersonID"`
	Books        Book       `gorm:"foreignKey:BookID"`
}
