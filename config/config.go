package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Conf struct {
	Host     string `yaml:"host"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	Port     string `yaml:"port"`
}

func (c *Conf) GetConf() *Conf {
	//讀取config/connect.yaml檔案
	yamlFile, err := ioutil.ReadFile("config/config.yaml")

	//若出現錯誤，列印錯誤訊息
	if err != nil {
		fmt.Println(err.Error())
	}

	//將讀取的字串轉換成結構體conf
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
