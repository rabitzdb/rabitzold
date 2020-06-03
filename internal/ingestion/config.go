package ingestion

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

//Read configuration from config file or environment variables
func init() {
	//Load from file /config/config.yml if environment is Development
	// production config is set on environment variables
	if os.Getenv("ENVIRONMENT") == "DEVELOPMENT" {
		viper.SetConfigName("config")   //File name
		viper.AddConfigPath("../../config") // Path where is config file
		viper.SetConfigType("yaml")         // Config type
		err := viper.ReadInConfig()         // Find and read the config file
		if err != nil {                     // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}
}

//Get config value from a key
func GetConfigValue(key string) string {
	var val string
	if os.Getenv("ENVIRONMENT") == "DEVELOPMENT" {
		val = viper.GetString(key)
	} else {
		val = os.Getenv(key)
	}
	return val
}
