package config

import (
	"github.com/spf13/viper"
	logg "realEstate/pkg/log"
)

// load viper config
func LoadConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		logg.Error("Logs not read")
	}
	logg.Info("Configs loaded succesfully!")
}
