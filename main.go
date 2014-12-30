package main

import (
	_ "core/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	models "core/models"
)

func main() {
	beego.Run()
}


func init() {
    orm.RegisterDriver("sqlite", orm.DR_Sqlite)
    orm.RegisterDataBase("default", "sqlite3", "database/orm_test.db")
    orm.RegisterModel(new(models.Article))
}
