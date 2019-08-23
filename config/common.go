package config

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// ProConfig 全局配置
var ProConfig Config

func init() {
	var err error

	v := viper.New()
	v.AddConfigPath("./")
	v.SetConfigName("conf")
	v.SetConfigType("yaml")

	//尝试进行配置读取
	if err = v.ReadInConfig(); err != nil {
		log.Println(err)
	}

	if err = v.Unmarshal(&ProConfig); err != nil {
		log.Println(err)
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置发生变化", e.String())
	})

}
