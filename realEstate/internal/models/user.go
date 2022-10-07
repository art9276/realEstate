package models

// User represents a user.
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

type Users struct {
	Users []User `json:"Users"`
}

func (u User) TableName() string {
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
