package middleware

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	//Id_user       int       `json:"Id_user"`
	Name          string `json:"Name"`
	Surename      string `json:"Surename"`
	Login         string `json:"Login"`
	Enc_password  string `json:"Enc_password"`
	Telephone     string `json:"Telephone"`
	Email         string `json:"Email"`
	Date_creation string `json:"Date_creation"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name, validation.Required, validation.Length(5, 30)),
		validation.Field(&u.Surename, validation.Required, validation.Length(5, 30)),
		validation.Field(&u.Login, validation.Required, validation.Length(5, 20)),
		validation.Field(&u.Enc_password, validation.Required, validation.Length(3, 30)),
		validation.Field(&u.Telephone, validation.Required, validation.Length(5, 30)),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Date_creation, validation.Required, validation.Length(5, 30)),
	)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
