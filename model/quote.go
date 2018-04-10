package model

import (
	"github.com/jinzhu/gorm"
	"alexbrasser/app/database"
)

type Quote struct {
	gorm.Model
	Name string `json:"name"`
	Text string `json:"text"`
	Icon string `json:"icon"`
	Link string `json:"link"`
}

func GetQuote(quote *Quote, id string) (error) {
	db := database.OpenMariaDB()
	defer db.Close()

	if err := db.First(&quote, id); err != nil {
		return err.Error
	}

	return nil

}

func GetQuotes(quotes *[]Quote) (error) {
	db := database.OpenMariaDB()
	defer db.Close()

	if err := db.Find(&quotes); err.Error != nil {
		return err.Error
	}

	return nil
}

func CreateQuote(quote *Quote) (error) {
	db := database.OpenMariaDB()
	defer db.Close()

	if err := db.Create(&quote); err.Error != nil {
		return err.Error
	}

	return nil
}
