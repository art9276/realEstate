package db

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	logg "realEstate/pkg/log"
)

var gor *gorm.DB //база данных

func InitGorm() {
	GormConfig()
	host := viper.Get("database.host")
	port := viper.Get("database.port")
	usergorm := viper.Get("database.user")
	password := viper.Get("database.password")
	dbname := viper.Get("database.dbname")
	dsn :=fmt.Sprintf( "host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		host,usergorm, password,dbname,port)
	logg.Info(dsn)
	con, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logg.Warning("Database is not initialize")
	}
	logg.Info("Database is connected with GORM")
	//TODO сделать дефер закрытие соединения
	println(con)

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


