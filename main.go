package main

import (
	"alexbrasser/app/server"
	"alexbrasser/app/database"
	"alexbrasser/model"
	modelQuote "alexbrasser/packages/quacky-quotes/model"
)

func main() {
	db := database.OpenMariaDB()

	db.DropTable(&model.Action{})
	db.AutoMigrate(&model.User{}, &model.Permission{}, modelQuote.Quote{}, &model.Client{}, &model.Action{})
	db.Close()
	server.Run()
}
