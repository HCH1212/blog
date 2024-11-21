package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Config struct {
	System System `yaml:"system"`
	Mysql  Mysql  `yaml:"mysql"`
	Logger Logger `yaml:"logger"`
}

// InitConf 读取yaml文件的配置
func InitConf() *Config {
	const ConfigFile = "config.yaml"
	c := &Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalln(fmt.Errorf("yaml Unmarshal err:%v", err))
	}
	log.Println("config yamlFile load init success")
	return c
}
