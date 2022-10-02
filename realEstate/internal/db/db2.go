package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	logg "realEstate/pkg/log"
)

var gor *gorm.DB //база данных

func InitGorm() {
	GormConfig()
	driver := viper.Get("database.driver")
	host := viper.Get("database.host")
	port := viper.Get("database.port")
	user := viper.Get("database.user")
	password := viper.Get("database.password")
	dbname := viper.Get("database.dbname")
	gormconn := fmt.Sprintf("%v://%v:%s@%v:%v/%v?sslmode=disable",
		driver, user, password, host, port, dbname)
	logg.Info(gormconn)
	con, err := gorm.Open("postgres", gormconn)
	if err != nil {
		logg.Warning("Database is not initialize")
	}
	logg.Info("Database is connected with GORM")
	gor = con
	//gor.Debug().AutoMigrate(&modelsOfEntty.User{}) //Миграция базы данных
}

func GormConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		logg.Error("Logs not read")
	}
}
func GetGormDB() *gorm.DB {
	return gor
}
