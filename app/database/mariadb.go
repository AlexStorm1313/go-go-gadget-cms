package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func OpenMariaDB() (*gorm.DB) {
	db, err := gorm.Open("mysql", "alexbrasser:alexbrasser@/alexbrasser?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err)
	}
	return db
}
