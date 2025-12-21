package main

import (
	"log"

	dbsql "github.com/lyx1213812138/BilibiliCleanPlan/dbSql"
	"github.com/lyx1213812138/BilibiliCleanPlan/get"
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
	get.SetMyVmid(config["user_id"].(float64))

	// init sql
	if err := dbsql.InitSql(); err != nil {
		log.Fatalf("init sql fail: %v", err)
	}

	// update database
	err = get.GetAndStoreHistory()
	if err != nil {
		log.Fatalf("get and store history fail: %v", err)
	}
	go func() { // TODO: db lock?
		err := get.StoreAllVideo()
		if err != nil {
			log.Fatalf("get and store video fail: %v", err)
		} else {
			log.Println("store all video success")
		}
	}()
	server()
}
