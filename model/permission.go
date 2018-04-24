package model

import (
	"alexbrasser/app/database"
	"time"
)

type Permission struct {
	ID        uint       `json:"id,omitempty" gorm:"primary_key"`
	UUID      string     `json:"uuid,omitempty" gorm:"unique_idex"`
	Name      string     `json:"name,omitempty"`
	Users     []*User    `json:"users,omitempty" gorm:"many2many:users_permissions"`
	Client    []*Client  `json:"clients,omitempty" gorm:"many2many:clients_permissions"`
	Action    []*Action  `json:"actions,omitempty" gorm:"many2many:actions_permissions"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type Permissions []Permission

func (this *Permission) Get(id string) error {
	client := database.OpenMariaDB()
	defer client.Close()

	if err := client.First(&this, id); err != nil {
		return err.Error
	}

	return nil
}

func (this *Permission) GetByUUID(uuid string) error {
	db := database.OpenMariaDB()
	defer db.Close()

	if err := db.Where(&Permission{UUID: uuid}).Find(&this); err != nil {
		return err.Error
	}
	return nil
}

func (this *Permissions) Get() error {
	client := database.OpenMariaDB()
	defer client.Close()

	if err := client.Find(&this); err.Error != nil {
		return err.Error
	}

	return nil
}

func (this *Permission) Update(data Permission) error {
	client := database.OpenMariaDB()
	defer client.Close()

	this.Name = data.Name

	if err := client.Model(&this).Updates(&this); err != nil {
		return err.Error
	}
	return nil
}

func (this *Permission) Create(data Permission) error {
	client := database.OpenMariaDB()
	defer client.Close()

	this.Name = data.Name

	if err := client.Create(&this); err.Error != nil {
		return err.Error
	}
	return nil
}

func (this *Permission) Delete() error {
	client := database.OpenMariaDB()
	defer client.Close()

	if err := client.Delete(&this); err.Error != nil {
		return err.Error
	}
	return nil
}
