package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_ "github.com/swaggo/swag/example/celler/httputil"
	_ "github.com/swaggo/swag/example/celler/model"
	conf "realEstate/internal/config"
	"realEstate/internal/db"
	"realEstate/internal/routes"
	logg "realEstate/pkg/log"
)

func init() {
	logg.GetLogger()
	//db.InitDB()
	db.InitGorm()
	db.GetGormDB()
	viper.New()
	conf.LoadConfig()
	routes.StartGin()
}

func main() {
	port := viper.Get("server.port")
	fmt.Println(port)
	logg.WithFields(logrus.Fields{
		"animal": "walrus",
		"number": 8,
	})
	logg.Info("232")
	defer db.GetGormDB().Close()
}
