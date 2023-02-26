package main

import (
	"bee_project/initializers"
	_ "bee_project/routers"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

// USE THIS FOR CONNECTION AND MIGRATION OF DATABASE USING BEEGO ORM
//func init() {
//orm.RegisterDriver("postgres", orm.DRPostgres)
//orm.RegisterDataBase("default", "postgres",
//	"host=localhost user=postgres password=Simform123 dbname=book_project port=5432 sslmode=disable")
//orm.RegisterModel(new(models.Books), new(models.Person))
//orm.RunSyncdb("default", false, true)
//}

// USE THIS FOR CONNECTION AND MIGRATION OF DATABASE USING GORM
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.OfficeConnector()
	initializers.SessionConnect()
}

func main() {
	beego.SetStaticPath("/static", "static")
	beego.Run()
}
