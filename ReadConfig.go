package main

import (
	"GoReadConfig/config"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func InitConfig() {
	yamlFile, err := ioutil.ReadFile("./config/config.yml")
	if err != nil {
		fmt.Println(err.Error())
	}
	var _config *config.Config
	err = yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("config.app: %#v\n", _config.App)
	fmt.Printf("config.log: %#v\n", _config.Log)

}

func main() {
	InitConfig()
}
