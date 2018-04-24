package model

import (
	"alexbrasser/app/database"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          uint          `json:"id,omitempty" gorm:"primary_key"`
	UUID        string        `json:"uuid,omitempty" gorm:"unique_index"`
	Email       string        `json:"email,omitempty" gorm:"unique_index"`
	Password    string        `json:"password,omitempty"`
	Permissions []*Permission `json:"permissions,omitempty" gorm:"many2many:users_permissions"`
	CreatedAt   time.Time     `json:"created_at,omitempty"`
	UpdatedAt   time.Time     `json:"updated_at,omitempty"`
	DeletedAt   *time.Time    `json:"deleted_at,omitempty"`
}

type Users []User

func (this *User) BeforeCreate(scope *gorm.Scope) error {
	if err := this.HashPassword(); err != nil {
		return err
	}

	scope.SetColumn("uuid", this.GenerateUUID())
	return nil
}

func (this *User) GenerateUUID() string {
	u, _ := uuid.NewV4()
	id := u.String()
	return id
}

func (this *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(this.Password), 10)
	if err != nil {
		return err
	}
	this.Password = string(bytes[:])
	return nil
}

func (this *User) CheckPasswordHash(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(this.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}

func (this *User) CreateToken() (error, *string) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = "api.zeus.dev"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["iat"] = time.Now().Unix()
	claims["user"] = &this.UUID

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err, nil
	}
	return nil, &t

}

func (this *User) Get(id string) error {
	db := database.OpenMariaDB()
	defer db.Close()

	if err := db.First(&this, id); err != nil {
		return err.Error
	}

	return nil
}

func (this *User) GetByUUID(uuid string) error {
	db := database.OpenMariaDB()
	defer db.Close()

	if err := db.Where(&User{UUID: uuid}).Find(&this); err != nil {
		return err.Error
	}
	return nil
}

func (this *Users) Get() error {
	client := database.OpenMariaDB()
	defer client.Close()

	if err := client.Find(&this); err.Error != nil {
		return err.Error
	}

	return nil
}

func (this *User) Update(data User) error {
	client := database.OpenMariaDB()
	defer client.Close()

	this.Email = data.Email
	this.Password = data.Password

	if err := client.Model(&this).Updates(&this); err != nil {
		return err.Error
	}
	return nil
}

func (this *User) Create(data User) error {
	client := database.OpenMariaDB()
	defer client.Close()

	this.Email = data.Email
	this.Password = data.Password

	if err := client.Create(&this); err.Error != nil {
		return err.Error
	}
	return nil
}

func (this *User) Delete() error {
	client := database.OpenMariaDB()
	defer client.Close()

	if err := client.Delete(&this); err.Error != nil {
		return err.Error
	}
	return nil
}

func (this *User) GetByEmail(email string) error {
	db := database.OpenMariaDB()
	defer db.Close()
	if err := db.Preload("Permissions").Where(&User{Email: email}).First(&this); err != nil {
		return err.Error
	}
	return nil
}

func (this *User) AddPermission(permission Permission) error {
	client := database.OpenMariaDB()
	defer client.Close()

	if err := client.Find(&this, this.ID).Association("Permissions").Append(permission); err != nil {
		return err.Error
	}
	return nil
}

func (this *User) DeletePermission(permission Permission) error {
	client := database.OpenMariaDB()
	defer client.Close()

	if err := client.Find(&this, this.ID).Association("Permissions").Delete(permission); err != nil {
		return err.Error
	}
	return nil
}
