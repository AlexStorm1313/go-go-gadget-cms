package model

import (
	"alexbrasser/app/cache"
	"alexbrasser/app/database"
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type Action struct {
	ID          uint          `json:"id,omitempty" gorm:"primary_key"`
	UUID        string        `json:"uuid,omitempty" gorm:"unique_index"`
	Method      string        `json:"method,omitempty"`
	Path        string        `json:"path,omitempty"`
	Name        string        `json:"name,omitempty"`
	Permissions []*Permission `json:"permissions" gorm:"many2many:actions_permissions"`
	CreatedAt   time.Time     `json:"created_at,omitempty"`
	UpdatedAt   time.Time     `json:"updated_at,omitempty"`
	DeletedAt   *time.Time    `json:"deleted_at,omitempty"`
}

type Actions []Action

func (this *Action) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("uuid", this.GenerateUUID())
	return nil
}

func (this *Action) AfterCreate(scope *gorm.Scope) error {
	if err := this.SaveInCache(); err != nil {
		return err
	}

	return nil
}

func (this *Action) AfterUpdate(scope *gorm.Scope) error {
	if err := this.SaveInCache(); err != nil {
		return err
	}
	return nil
}

func (this *Action) AfterDelete(scope *gorm.Scope) error {
	if err := this.DeleteFromCache(); err != nil {
		return err
	}
	return nil
}

func (this *Action) GenerateUUID() string {
	u, _ := uuid.NewV4()
	id := u.String()
	return id
}

func (this *Action) MarshalBinary() ([]byte, error) {
	return json.Marshal(this)
}

func (this *Action) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &this); err != nil {
		return err
	}

	return nil
}

func (this *Action) Get(id string) error {
	database := database.OpenMariaDB()
	defer database.Close()

	if err := database.First(&this, id); err != nil {
		return err.Error
	}

	return nil
}

func (this *Action) GetByUUID(uuid string) error {
	database := database.OpenMariaDB()
	defer database.Close()

	if err := database.Where(&Action{UUID: uuid}).Find(&this); err != nil {
		return err.Error
	}
	return nil
}

func (this *Actions) Get() error {
	database := database.OpenMariaDB()
	defer database.Close()

	if err := database.Find(&this); err != nil {
		return err.Error
	}

	return nil
}

func (this *Action) Update(data Action) error {
	database := database.OpenMariaDB()
	defer database.Close()

	this.Method = data.Method
	this.Path = data.Path
	this.Name = data.Name

	if err := database.Model(&this).Updates(&this); err != nil {
		return err.Error
	}

	return nil
}

func (this *Action) Create(data Action) error {
	database := database.OpenMariaDB()
	defer database.Close()

	this.Method = data.Method
	this.Path = data.Path
	this.Name = data.Name

	if err := database.Create(&this); err.Error != nil {
		return err.Error
	}
	return nil
}

func (this *Action) Delete() error {
	database := database.OpenMariaDB()
	defer database.Close()

	if err := database.Delete(&this); err.Error != nil {
		return err.Error
	}

	if err := this.DeleteFromCache(); err != nil {
		return err
	}

	return nil
}

func (this *Action) SaveInCache() error {
	cache := cache.OpenRedis()
	defer cache.Close()

	if err := cache.Set(this.UUID, this, 0).Err(); err != nil {
		return err
	}

	return nil
}

func (this *Action) GetFromCache() error {
	cache := cache.OpenRedis()
	defer cache.Close()

	data, err := cache.Get(this.UUID).Result()
	if err != nil {
		return err
	}
	this.UnmarshalBinary([]byte(data))
	return nil
}

func (this *Action) DeleteFromCache() error {
	cache := cache.OpenRedis()
	defer cache.Close()

	if err := cache.Del(this.UUID).Err(); err != nil {
		return err
	}

	return nil
}

func (this *Action) AddPermission(permission Permission) error {
	database := database.OpenMariaDB()
	defer database.Close()

	if err := permission.GetByUUID(permission.UUID); err != nil {
		return err
	}

	if err := database.Preload("Permissions").Find(&this, this.ID).Association("Permissions").Append(&permission); err != nil {
		return err.Error
	}
	action := &Action{}

	if err := database.Preload("Permissions").Find(action, this.ID); err.Error != nil {
		return err.Error
	}

	if err := action.SaveInCache(); err != nil {
		return err
	}

	return nil
}

func (this *Action) DeletePermission(permission Permission) error {
	database := database.OpenMariaDB()
	defer database.Close()

	if err := permission.GetByUUID(permission.UUID); err != nil {
		return err
	}

	if err := database.Model(&this).Association("Permissions").Delete(&permission); err != nil {
		return err.Error
	}

	if err := this.GetPermissions(); err != nil {
		return err
	}

	// if err := database.Preload("Permissions").Find(&this, this.ID).Association("Permissions").Delete(&permission); err != nil {
	// 	return err.Error
	// }
	// action := &Action{}

	// if err := database.Preload("Permissions").Find(action, this.ID); err.Error != nil {
	// 	return err.Error
	// }

	if err := this.SaveInCache(); err != nil {
		return err
	}

	return nil
}

func (this *Action) GetPermissions() error {
	database := database.OpenMariaDB()
	defer database.Close()

	if err := database.Preload("Permissions").First(&this); err.Error != nil {
		return err.Error
	}

	// if err := this.SaveInCache(); err != nil {
	// 	return err
	// }

	return nil
}
