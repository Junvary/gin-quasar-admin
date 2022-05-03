package boot

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	v := viper.New()
	//配置文件位置
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	//读取配置文件
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("没有发现配置文件：%s \n", err.Error()))
		} else {
			panic(fmt.Errorf("读取配置文件失败：%s \n", err))
		}
	}
	//监视配置文件变化
	v.WatchConfig()
	//配置文件变化时重新绑定
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件发生变化：", e.Name)
		BindConfig(v)
	})
	//绑定配置文件
	BindConfig(v)
	return v
}

func BindConfig(v *viper.Viper) {
	//绑定配置文件中的所有配置项
	if err := v.Unmarshal(&global.GqaConfig); err != nil {
		panic(fmt.Errorf("绑定配置文件失败：%s \n", err))
	} else {
		fmt.Println("绑定配置文件成功！")
	}
}
