package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
	"github.com/pkg/errors"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type ENV_CONFIG struct {
	Runmode string
}

var config ENV_CONFIG  //环境变量
var dataJson map[string]interface{} //配置文件Json数据

func init() {
	if err := envconfig.Process("config", &config); err != nil {
		log.Println(err)
		panic(errors.New("no CONFIG_RUNMODE exported."))
	}
	
	//加载配置文件
	configPath := "config/dev.json"
	switch config.Runmode {
	case "test":
		configPath = "config/test.json"
	case "prod":
		configPath = "config/prod.json"
	}
	log.Println("configPath is: ", configPath)
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Println("no config file found.")
		return
	}
	//fmt.Println("data is: ", string(data))
	
	if err := json.Unmarshal(data, &dataJson); err != nil{
		fmt.Println(err)
		panic(errors.New(fmt.Sprintf("config file: %s is not json format", configPath)))
	}
	
	log.Println("config init finished.")
}

/**
获取runmode
 */
func GetRunMode() (string) {
	if config.Runmode == "" {
		return "prod"
	}
	return config.Runmode
}

/**
获取配置文件Json数据
 */
func GetConfigJson() (map[string]interface{}) {
	return dataJson
}

