package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"

)

var db *sql.DB

func InitDB() {
	LoaddbConfig()
	driver := viper.Get("database.driver")
	host := viper.Get("database.host")
	port := viper.Get("database.port")
	user := viper.Get("database.user")
	password := viper.Get("database.password")
	dbname := viper.Get("database.dbname")
	// connection string
	psqlconn := fmt.Sprintf("%v://%v:%s@%v:%v/%v?sslmode=disable",
		driver, user, password, host, port, dbname)
	//connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
	//	postgres,user,password,host,port,dbname)
	// open database
	fmt.Println(psqlconn)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println("The database is not initialize")
	}
	if err = db.Ping(); err != nil {
		fmt.Println("The database is not connected")
	}
	fmt.Println("The database is connected")
}

func LoaddbConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		println("Logs not read")
	}
}
