package config

import (
	"log"

	"github.com/spf13/viper"
)

// ProConfig 全局配置
var ProConfig Config

func init() {
	var err error

	configReader := viper.New()
	configReader.AddConfigPath("./")
	configReader.SetConfigName("config")
	configReader.SetConfigType("yaml")

	//尝试进行配置读取
	if err = configReader.ReadInConfig(); err != nil {
		panic(err)
	}

	if err = configReader.Unmarshal(&ProConfig); err != nil {
		log.Println(err)
	}

}
