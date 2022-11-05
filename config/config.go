package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Init() {
	// 设置config文件名称，类型，路径
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
