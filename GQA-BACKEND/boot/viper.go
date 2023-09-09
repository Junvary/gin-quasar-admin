package boot

import (
	"errors"
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
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			panic(fmt.Errorf("Config file not found：%s \n", err.Error()))
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
