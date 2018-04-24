package model

import (
	"alexbrasser/app/database"
	"time"
)

type Quote struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	UUID string `json:"uuid" gorm:"unique_index"`

	Name string `json:"name"`
	Text string `json:"text"`
	Icon string `json:"icon"`
	Link string `json:"link"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Quotes []Quote

func (this *Quote) Get(id string) (error) {
	db := database.OpenMariaDB()
	defer db.Close()

	if err := db.First(&this, &id); err != nil {
		return err.Error
	}

	return nil
}

func (this *Quotes) Get() (error) {
	db := database.OpenMariaDB()
	defer db.Close()

	if err := db.Find(&this); err.Error != nil {
		return err.Error
	}

	return nil
}

func (this *Quote) Save() (error) {
	db := database.OpenMariaDB()
	defer db.Close()

	if err := db.Save(&this); err != nil {
		return err.Error
	}
	return nil
}

func (this *Quote) Create() (error) {
	db := database.OpenMariaDB()
	defer db.Close()

	if err := db.Create(&this); err.Error != nil {
		return err.Error
	}
	return nil
}

func (this *Quote) Delete() (error) {
	db := database.OpenMariaDB()
	defer db.Close()

	if err := db.Delete(&this); err.Error != nil {
		return err.Error
	}
	return nil
}