package modelsOfEntty

import (
	"gorm.io/gorm"
	"time"
)

// User represents a user.
type Userm struct {
	gorm.Model	`json:"model"`
	ID           int       `gorm:"primaryKey" json:"Id"`
	Name         string    `gorm:"column:Name" json:"Name"`
	LastName     string    `gorm:"column:Surename" json:"Surename"`
	Date_Creation time.Time `gorm:"column:Date_creation" json:"Date_creation"`
	Date_Update   time.Time `gorm:"column:Date_update" json:"Date_update"`
	Login        string    `gorm:"column:Login" json:"Login"`
	Password     string    `gorm:"column:" json:"Enc_password"`
	Telephone    string    `gorm:"column:Telephone" json:"Telephone"`
	Email        string    `gorm:"column:Email" json:"Email"`
}

type Users []Userm

func (u *Userm) NotFound() bool {
	return u.Model.ID == 0
}

func (u Userm) TableName() string {
	// custom table name, this is default
	return "Users"
}

// login unaforized users
/*func Login(email, password string) map[string]interface{} {

	user := &User{}
	err := db.ConnectGormDB().Table("users").Where("email = ?", email).First(user).Error
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

 */


