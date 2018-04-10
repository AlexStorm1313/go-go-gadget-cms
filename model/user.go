package model

import (
	"github.com/jinzhu/gorm"
	"alexbrasser/app/database"
	"log"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Email    string  `json:"email" gorm:"unique_index"`
	Password string  `json:"password"`
	Roles    []*Role `json:"roles" gorm:"many2many:users_roles"`
}

func (this *User)BeforeCreate() (error){
	if err := this.HashPassword(); err != nil{
		return err
	}
	return nil
}

func (this *User)HashPassword() (error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(this.Password), 14)
	if err != nil {
		return err
	}
	this.Password = string(bytes[:])
	return nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUser(user *User, id string) (error) {
	db := database.OpenMariaDB()
	defer db.Close()

	if err := db.Preload("Roles").First(&user, id); err != nil {
		return err.Error
	}

	return nil

}

func GetUsers(users *[]User) (error) {
	db := database.OpenMariaDB()
	defer db.Close()

	if err := db.Preload("Roles").Find(&users); err.Error != nil {
		return err.Error
	}
	log.Println("bier")

	return nil
}

func CreateUser(user *User) (error) {
	db := database.OpenMariaDB()
	defer db.Close()

	if err := db.Create(&user); err.Error != nil {
		return err.Error
	}
	return nil
}

func GetUserByEmailAndPassword(user *User) (error) {
	db := database.OpenMariaDB()
	defer db.Close()

	if err := db.Where(&user).First(&user); err.Error != nil {
		return err.Error
	}
	return nil
}
