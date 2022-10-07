package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"realEstate/internal/db"
	"realEstate/internal/models"
)

func Logout(c *gin.Context) {
	//TODO add cookie and clear whem
c.Redirect(http.StatusSeeOther,"/")
}

// Login godoc
// @Summary Logging user
// @Description Logging user
// @ID 3
// @Accept *gin.Context
// @Param some_id path int true "Some ID"
// @Param some_id body web.Pet true "Some ID"
// @Success 200 {string} string "Login and Password correct"
// @Success 303 {object} http.Redirect("/")
// @Failure 400 {object} web.APIError "User not found"
// @Router /auth/login [get]
func Login(c *gin.Context) {
	//TODO add cookie
	Login := c.Query("Login")
	Enc_password:=c.Query("Enc_password")
	var user models.User
	row := db.InitDB().QueryRow(`SELECT "Login", "Enc_password"
 	 FROM public."Users" where "Login"=$1 
 	AND "Enc_password"=$2`, Login,Enc_password)
	err2 := row.Scan(&user.Login, &user.Enc_password)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User not found",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Login and Password correct",
	})
	c.Redirect(http.StatusSeeOther,"/")
}

// GetAllUser godoc
// @Summary Add a new pet to the store
// @Description get string by ID
// @ID 4
// @Accept json
// @Produce json
// @Param some_id path int true "Some ID"
// @Param some_id body web.Pet true "Some ID"
// @Success 200 {string} string "Welcome to site"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /users [get]
func GetAllUser(c *gin.Context) {
	// если добавлять поле date_creation,date_update
	rows, err := db.InitDB().Query(`SELECT "Name", "Surename", 
       "Login", "Enc_password", "Telephone", "Email" ,"Date_creation" FROM public."Users"`)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Query not Scanned",
		})
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		err2 := rows.Scan(&user.Name, &user.Surename,
			&user.Login, &user.Enc_password, &user.Telephone, &user.Email, &user.Date_creation)
		// Exit if we get an error
		if err2 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Users not scanned",
			})
		}
		users = append(users, user)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": users,
	})

}

// CreateUser godoc
// @Summary Add a new pet to the store
// @Description get string by ID
// @ID 5
// @Accept json
// @Produce json
// @Param some_id path int true "Some ID"
// @Param some_id body web.Pet true "Some ID"
// @Success 200 {string} string "Welcome to site"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /auth/register [post]
func CreateUser(c *gin.Context) {
	u := new(models.User)
	if err := c.BindJSON(u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "JSON non scanned",
		})
	}
	sqlStatement := `INSERT INTO public."Users"(
			"Name", "Surename", "Login", 
			"Enc_password", "Telephone", "Email","Date_creation")
			VALUES ($1, $2, $3, $4, $5, $6,$7) `
	res, err := db.InitDB().Query(sqlStatement, u.Name,
		u.Surename, u.Login,
		u.Enc_password, u.Telephone, u.Email, u.Date_creation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to saved user",
		})
	} else {
		fmt.Println(res)
		c.JSON(http.StatusCreated, gin.H{
			"User created": u,
		})
	}
}


// GetUsers godoc
// @Summary Add a new pet to the store
// @Description get string by ID
// @ID 6
// @Accept json
// @Produce json
// @Param some_id path int true "Some ID"
// @Param some_id body web.Pet true "Some ID"
// @Success 200 {string} string "Welcome to site"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /users/:id [get]
func GetUser(c *gin.Context) {
	Email := c.Query("Email")
	var user models.User
	row := db.InitDB().QueryRow(`SELECT "Name", "Surename","Login", "Enc_password",
 	"Telephone", "Email", "Date_creation" FROM public."Users" where "Email"=$1`, Email)
	err := row.Scan(&user.Name, &user.Surename, &user.Login, &user.Enc_password, &user.Telephone,
		&user.Email, &user.Date_creation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Scan not complited",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"User for email": user,
	})

}

// UpdateUsers godoc
// @Summary Add a new pet to the store
// @Description get string by ID
// @ID 7
// @Accept json
// @Produce json
// @Param some_id path int true "Some ID"
// @Param some_id body web.Pet true "Some ID"
// @Success 200 {string} string "Welcome to site"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /users/:id [put]
func UpdateUser(c *gin.Context) {
	//id := c.Param("id")
	u := new(models.User)
	if err := c.BindJSON(u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "JSON non scanned",
		})
	}
	sqlStatement := `UPDATE public."Users" SET
			"Name"=$1, "Surename"=$2, "Login"=$3, 
			"Enc_password"=$4, "Telephone"=$5, "Email"=$6,"Date_creation"=$7
			Where "Email"=$8`
	res, err := db.InitDB().Query(sqlStatement, u.Name,
		u.Surename, u.Login,
		u.Enc_password, u.Telephone, u.Email, u.Date_creation, u.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to updated user",
		})
	} else {
		fmt.Println(res)

		c.JSON(http.StatusOK, gin.H{
			"User updated": u,
		})
	}
}

// // DeleteUsers godoc
// @Summary Add a new pet to the store
// @Description get string by ID
// @ID 8
// @Accept json
// @Produce json
// @Param some_id path int true "Some ID"
// @Param some_id body web.Pet true "Some ID"
// @Success 200 {string} string "Welcome to site"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /users [delete]
func DeleteUser(c *gin.Context) {
	Email := c.Query("Email")
	sqlStatement := `Delete from public."Users" where "Email"=$1`
	res, err := db.InitDB().Query(sqlStatement, Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User not found",
		})
	} else {
		fmt.Println(res)

		c.JSON(http.StatusOK, "User deleted")
	}
}

// ShowIndex godoc
// @Schemes
// @Description Welcome page
// @Tags MainPage
// @Accept json
// @Produce json
// @Success 200 {string} Welcome to site
// @Router / [get]
func ShowIndexPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to site",
	})
	return
}

// NotFound godoc
// @Summary Add a new pet to the store
// @Description get string by ID
// @ID 9
// @Accept json
// @Produce json
// @Param some_id path int true "Some ID"
// @Param some_id body web.Pet true "Some ID"
// @Success 200 {string} string "Welcome to site"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router  [get]
func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "request not found",
	})
	return
}

// UploadFiles godoc
// @Summary Add a new pet to the store
// @Description get string by ID
// @ID 10
// @Accept json
// @Produce json
// @Param some_id path int true "Some ID"
// @Param some_id body web.Pet true "Some ID"
// @Success 200 {string} string "Welcome to site"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /upload [post]

func UploadFiles(c *gin.Context) {
	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "get form error",
		})
		return
	}
	files := form.File["files"]

	for _, file := range files {
		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "upload file err"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Uploaded successfully"})
	return
}
