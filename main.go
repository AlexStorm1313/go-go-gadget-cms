package main

import (
	"alexbrasser/app/server"
	"alexbrasser/app/database"
	"alexbrasser/model"
)

func main() {
	db := database.OpenMariaDB()

	db.AutoMigrate(&model.User{}, &model.Role{}, &model.Quote{}, &model.Client{})
	db.Close()
	server.Run()
}