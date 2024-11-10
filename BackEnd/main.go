package main

import (
	"log"

	"github.com/lyx1213812138/BilibiliCleanPlan/data"
	dbsql "github.com/lyx1213812138/BilibiliCleanPlan/dbSql"
	"github.com/spf13/viper"
)

func main() {
	// Load configuration by viper
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetDefault("recommend.num", 20)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	config := viper.AllSettings()
	data.SetMyVmid(config["user_id"].(float64))

	dbsql.InitSql()
	// update data from api
	server()
}
