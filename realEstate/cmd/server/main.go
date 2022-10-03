package main

import (
	"github.com/spf13/viper"
	_ "github.com/swaggo/swag/example/celler/httputil"
	_ "github.com/swaggo/swag/example/celler/model"
	conf "realEstate/internal/config"
	"realEstate/internal/db"
	"realEstate/internal/routes"
	logg "realEstate/pkg/log"
)

func main() {
	logg.GetLogger()
	//db.InitDB()
	db.InitGorm()
	db.GetGormDB()
	viper.New()
	conf.LoadConfig()
	routes.StartGin()
	//db.InitRedis()
	logg.Info("Server is running on port:8080")
}
