package modelsOfEntty

import (
	"golang.org/x/crypto/bcrypt"
	"realEstate/internal/db"
	m "realEstate/pkg/json"

	"time"
)

// User represents a user.
type User struct {
	ID           int       `gorm:"column:Id_user" json:"Id_user"`
	Name         string    `gorm:"column:Name" json:"Name"`
	LastName     string    `gorm:"column:Surename" json:"Surename"`
	DateCreation time.Time `gorm:"column:Date_creation" json:"Date_creation"`
	DateUpdate   time.Time `gorm:"column:Date_update" json:"Date_update"`
	Login        string    `gorm:"column:Login" json:"Login"`
	Password     string    `gorm:"column:" json:"Enc_password"`
	Telephone    string    `gorm:"column:Telephone" json:"Telephone"`
	Email        string    `gorm:"column:Email" json:"Email"`
}

var hashedPassword string

// GetID returns the user ID.

func (u *User) User() string {
	return "Users"
}

func (u *User) CreateUser() error {
	return db.GetGormDB().Create(u).Error
}

func Login(email, password string) map[string]interface{} {

	user := &User{}
	err := db.GetGormDB().Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		return m.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return m.Message(false, "Invalid login credentials. Please try again")
	}
	//Worked! Logged In

	resp := m.Message(true, "Logged In")
	resp["user"] = user
	return resp
}

func GetUser(u uint) *User {

	us := &User{}
	db.GetGormDB().Table("accounts").Where("id = ?", u).First(us)
	if us.Email == "" { //User not found!
		return nil
	}

	us.Password = hashedPassword
	return us
}
