package configParcer

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Username     string
	Password     string
	DataBaseName string
	Url          string
}

var (
	defaults = map[string]interface{}{
		"username":     "admin",
		"password":     "admin",
		"databasename": "transactions",
	}
	configName = "config"
	configPath = "./configs"
	C          Config
)

func init() {
	for k, v := range defaults {
		viper.SetDefault(k, v)
	}
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("could not read config file: %v", err)
	}

	if err := viper.Unmarshal(&C); err != nil {
		log.Fatalf("could not decode config file into struct: %v", err)
	}
	C.Url = fmt.Sprintf("http://%v:%v@127.0.0.1:5984/", C.Username, C.Password)
}
