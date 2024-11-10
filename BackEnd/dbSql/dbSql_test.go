package dbsql

import (
	"log"

	"github.com/lyx1213812138/BilibiliCleanPlan/data"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("../")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	config := viper.AllSettings()
	data.SetMyVmid(config["user_id"].(float64))
	InitSql()
}
