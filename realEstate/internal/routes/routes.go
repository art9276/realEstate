package routes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"realEstate/internal/db"
	logg "realEstate/pkg/log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/swag/example/celler/httputil"
	_ "github.com/swaggo/swag/example/celler/model"
	"net/http"
	_ "realEstate/cmd/server/docs"
	"realEstate/internal/modelsOfEntty"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

var router *gin.Engine

// @BasePath /api/v1

func StartGin() {
	router = gin.Default()
	InitializeRoutes()
	if err := router.Run("localhost:8080"); err != nil {
		logg.Fatal("Server not run!")
	}
	logg.Info("Server is running")
}

func InitializeRoutes() {
	router.GET("/", ShowIndexPage)
	router.GET("/users", GetAllUser)
	router.POST("/auth/register", CreateUser)
	router.GET("/auth/login", Login)
	router.GET("/auth/logout", Logout)
	router.GET("/users/:id", GetUser)
	router.PUT("/users/:id", UpdateUser)
	router.DELETE("/users/:id", DeleteUser)
	router.GET(" ", NotFound)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	logg.Info("Routes initialized")
}

func Logout(c *gin.Context) {

}

func Login(c *gin.Context) {

}

// GetStringByInt example
// @Summary Add a new pet to the store
// @Description get string by ID
// @ID get-string-by-int
// @Accept json
// @Produce json
// @Param some_id path int true "Some ID"
// @Param some_id body web.Pet true "Some ID"
// @Success 200 {string} string "Welcome to site"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /users [get]
func GetAllUser(c *gin.Context) {
	users := []modelsOfEntty.User{}
	us := db.GetGormDB().Find(&users)
	println(us)
	json.Marshal(us)
	c.JSON(http.StatusOK, gin.H{
		"message": us,
	})
}
func CreateUser(c *gin.Context) {
	users := []modelsOfEntty.User{}
	if err := c.BindJSON(&users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "JSON not unserialized",
		})
	}
	if err := db.GetGormDB().Save(users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "JSON not serialized",
		})
	}
	json.Marshal(users)
	c.JSON(http.StatusOK, gin.H{
		"message": users,
	})
}
func GetUser(c *gin.Context) {
	id := c.Param("id")
	user := &modelsOfEntty.User{}
	if db.GetGormDB().First(user, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	user := &modelsOfEntty.User{}
	if db.GetGormDB().First(user, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	updated := &modelsOfEntty.User{}
	if err := c.BindJSON(&updated); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "user not unmarshal!"})
		return
	}
	if err := db.GetGormDB().Save(user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "updated user not saved!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user updated",
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	user := &modelsOfEntty.User{}
	if db.GetGormDB().First(user, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	if err := db.GetGormDB().Delete(user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "user not deleted"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted",
	})
}

// PingExample godoc// @Summary ping example
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
}
func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "request not found",
	})

}
