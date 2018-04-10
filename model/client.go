package model

import (
	"github.com/jinzhu/gorm"
	"alexbrasser/app/database"
	"github.com/gobuffalo/uuid"
)

type Client struct {
	gorm.Model
	Name   string `json:"name"`
	Type   string `json:"type"`
	Secret string `json:"secret" gorm:"unique_index"`
}

func GetClient(client *Client, id string) (error) {
	db := database.OpenMariaDB()
	defer db.Close()

	if err := db.First(&client, id); err != nil {
		return err.Error
	}

	return nil

}

func CreateClient(client *Client) (error) {
	db := database.OpenMariaDB()
	defer db.Close()

	secret, err := uuid.NewV4();
	if err != nil {
		return err
	}

	client.Secret = secret.String()

	if err := db.Create(&client); err.Error != nil {
		return err.Error
	}

	return nil
}

func GetClientBySecret(client *Client) (error) {
	db := database.OpenMariaDB()
	defer db.Close()

	if err := db.Where(&client).First(&client); err.Error != nil {
		return err.Error
	}

	return nil
}
