package main

import (
	"GoReadConfig/config"
	"fmt"
	"github.com/spf13/viper"
)

func InitConfigByViper() {

	viper.SetConfigType("yaml")
	viper.SetConfigFile("./config/config.yml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	var _config *config.Config
	err = viper.Unmarshal(&_config)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("config.app: %#v\n", _config.App)
	fmt.Printf("config.log: %#v\n", _config.Log)

}

func main() {
	InitConfigByViper()
}
