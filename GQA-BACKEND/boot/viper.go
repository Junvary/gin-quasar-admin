package boot

import (
	"fmt"
	"gin-quasar-admin/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Viper() *viper.Viper {
	v := viper.New()

	v.SetConfigName("config.yaml")
	v.SetConfigType("yaml")
	v.AddConfigPath("config")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("启动失败，没有发现配置文件：%s \n", err.Error()))
		} else {
			panic(fmt.Errorf("启动失败，读取配置文件失败：%s \n", err))
		}
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件改变：", e.Name)
		global.GqaLog.Info("配置文件改变！")
		if err := v.Unmarshal(&global.GqaConfig); err != nil {
			fmt.Println("绑定配置文件失败：", err.Error())
			global.GqaLog.Error("更改配置文件失败：", zap.Any("err", err))
		}else {
			fmt.Println("重新配置完成！")
			global.GqaLog.Info("重新配置完成！")
		}
	})
	if err := v.Unmarshal(&global.GqaConfig); err != nil {
		panic(fmt.Errorf("绑定配置文件失败：%s \n", err))
	}
	return v
}
