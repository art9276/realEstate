package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"realEstate/internal/db"
	"realEstate/internal/models"
)

func Logout(c *gin.Context) {

}

// Login godoc
// @Summary Add a new pet to the store
// @Description get string by ID
// @ID 3
// @Accept json
// @Produce json
// @Param some_id path int true "Some ID"
// @Param some_id body web.Pet true "Some ID"
// @Success 200 {string} string "Welcome to site"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /auth/login [get]
func Login(c *gin.Context) {

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

	rows, err := db.InitDB().Query(`SELECT "Id_user", "Name", "Surename", "Date_creation", 
       "Login", "Enc_password", "Telephone", "Email" FROM public."Users"`)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Query not Scanned",
		})
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		err2 := rows.Scan(&user.Id_user, &user.Name, &user.Surename, &user.Date_creation,
			&user.Login, &user.Enc_password, &user.Telephone, &user.Email)
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

	c.JSON(http.StatusOK, gin.H{
		"message": "123",
	})
	return
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
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"message": id,
	})
	return
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

	c.JSON(http.StatusOK, gin.H{
		"message": "user updated",
	})
	return
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
// @Router /users/:id [delete]
func DeleteUser(c *gin.Context) {
	//id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted",
	})
	return
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
