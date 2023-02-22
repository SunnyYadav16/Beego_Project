package main

import (
	"bee_project/models"
	_ "bee_project/routers"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres",
		"host=localhost user=postgres password=Simform123 dbname=book_project port=5432 sslmode=disable")
	orm.RegisterModel(new(models.Books), new(models.Person))
	orm.RunSyncdb("default", false, true)
}

func main() {
	beego.Run()
}
