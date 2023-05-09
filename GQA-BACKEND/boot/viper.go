package boot

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("Config file not found：%s \n", err.Error()))
		} else {
			panic(fmt.Errorf("Config file not found：%s \n", err))
		}
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config changed：", e.Name)
		BindConfig(v)
	})
	BindConfig(v)
	return v
}

func BindConfig(v *viper.Viper) {
	if err := v.Unmarshal(&global.GqaConfig); err != nil {
		panic(fmt.Errorf("Failed to bind config file：%s \n", err))
	} else {
		fmt.Println("Bind config file succeeded！")
	}
}
