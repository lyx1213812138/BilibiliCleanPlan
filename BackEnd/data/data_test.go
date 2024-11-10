package data

import (
	"log"
	"testing"

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
	SetMyVmid(config["user_id"].(float64))
}

func TestGetUp(t *testing.T) {
	log.Printf("TestGetUp")
	u, err := GetUp()
	if err != nil {
		t.Error(err)
	}
	log.Printf("%#v", u)
}
