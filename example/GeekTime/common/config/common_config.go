package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var BasePath string

func InitConfigs() {
	BasePath = "conf/"
	initLog()
}

// 加载配置信息
func initConf(path string, config interface{}) error {
	content, err := ioutil.ReadFile(path)
	fmt.Printf("conten: %s", content)
	defer func() {
		if err != nil {
			log.Fatalf("%v", err.Error())
		}
	}()
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(content, config)
	fmt.Printf("err: %s", err)
	if err != nil {
		return err
	}
	//logs.Info("%+v", config)
	return err
}
