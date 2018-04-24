package model

import (
	"alexbrasser/app/database"
	"github.com/gobuffalo/uuid"
	"time"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
	"github.com/dgrijalva/jwt-go"
)

type Client struct {
	ID          uint          `json:"id,omitempty" gorm:"primary_key"`
	UUID        string        `json:"uuid,omitempty" gorm:"unique_index"`
	Name        string        `json:"name,omitempty"`
	Type        string        `json:"type,omitempty"`
	Secret      string        `json:"secret,omitempty" gorm:"unique_index"`
	Permissions []*Permission `json:"permissions,omitempty" gorm:"many2many:clients_permissions"`
	CreatedAt   time.Time     `json:"created_at,omitempty"`
	UpdatedAt   time.Time     `json:"updated_at,omitempty"`
	DeletedAt   *time.Time    `json:"deleted_at,omitempty"`
}

type Clients []Client

func (this *Client) BeforeCreate(scope *gorm.Scope) (error) {
	scope.SetColumn("uuid", this.GenerateUUID())
	scope.SetColumn("secret", random.String(128))
	return nil
}

func (this *Client) GenerateUUID() (string) {
	u, _ := uuid.NewV4()
	id := u.String()
	return id
}

func (this *Client) CreateToken() (error, *string) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = "api.zeus.dev"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["iat"] = time.Now().Unix()
	claims["client"] = &this.UUID

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err, nil
	}
	return nil, &t
}

func (this *Client) Get(id string) (error) {
	client := database.OpenMariaDB()
	defer client.Close()

	if err := client.First(&this, id); err != nil {
		return err.Error
	}

	return nil
}

func (this *Client) GetByUUID(uuid string) (error){
	db := database.OpenMariaDB()
	defer db.Close()

	if err := db.Where(&Client{UUID: uuid}).Find(&this); err !=nil{
		return err.Error
	}
	return nil
}

func (this *Clients) Get() (error) {
	client := database.OpenMariaDB()
	defer client.Close()

	if err := client.Find(&this); err != nil {
		return err.Error
	}

	return nil
}

func (this *Client) Update(data Client) (error) {
	client := database.OpenMariaDB()
	defer client.Close()

	this.Name = data.Name
	this.Type = data.Type

	if err := client.Model(&this).Updates(&this); err != nil {
		return err.Error
	}
	return nil
}

func (this *Client) Create(data Client) (error) {
	client := database.OpenMariaDB()
	defer client.Close()

	this.Name = data.Name
	this.Type = data.Type

	if err := client.Create(&this); err.Error != nil {
		return err.Error
	}
	return nil
}

func (this *Client) Delete() (error) {
	client := database.OpenMariaDB()
	defer client.Close()

	if err := client.Delete(&this); err.Error != nil {
		return err.Error
	}
	return nil
}

func (this *Client) GetBySecret(secret string) (error) {
	db := database.OpenMariaDB()
	defer db.Close()

	if err := db.Where(&Client{Secret: secret}).Find(&this); err != nil {
		return err.Error
	}
	return nil
}

func (this *Client) AddPermission(permission Permission)(error){
	client := database.OpenMariaDB()
	defer client.Close()

	if err := client.Find(&this, this.ID).Association("Permissions").Append(permission); err != nil {
		return err.Error
	}
	return nil
}

func (this *Client) DeletePermission(permission Permission) (error){
	client := database.OpenMariaDB()
	defer client.Close()

	if err := client.Find(&this, this.ID).Association("Permissions").Delete(permission); err != nil {
		return err.Error
	}
	return nil
}