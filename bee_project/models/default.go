package models

type Books struct {
	Id     int    `orm:"column(id);auto"`
	Name   string `orm:"column(name);size(255)"`
	Author string `orm:"column(author);size(255)"`
	Price  int    `orm:"column(price)"`
}

type Person struct {
	Id             int      `orm:"column(id);auto"`
	Name           string   `orm:"column(name);size(255)"`
	Mobile         string   `orm:"column(mobile);size(255)"`
	Email          string   `orm:"column(email);size(255)"`
	BooksPurchased []*Books `orm:"reverse(many)"`
}
