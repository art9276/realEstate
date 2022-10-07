package main

import (
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/swag/example/celler/httputil"
	_ "github.com/swaggo/swag/example/celler/model"
	"realEstate/internal/db"
	"realEstate/internal/routes"
	_ "realEstate/internal/routes"
	logg "realEstate/pkg/log"
)

// @title         Real Estate Service
// @version       1.0
// @description   Service to byu and sell houses and flats  Gin framework.
// @contact.name  Alexander
// @contact.email sntshkmr60@gmail.com
// @license.name  Apache 2.0
// @host          localhost:8080

func main() {
	logg.GetLogger()
	db.InitDB()
	//viper.New()
	//conf.LoadConfig()
	routes.StartGin()
	//db.InitRedis()
	logg.Info("Server is running on port:8080")

}
