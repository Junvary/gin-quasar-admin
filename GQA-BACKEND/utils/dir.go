package utils

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"go.uber.org/zap"
	"os"
)

func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExist(v)
		if err != nil {
			return err
		}
		if !exist {
			global.GqaLog.Debug("创建目录：" + v)
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				global.GqaLog.Error("创建目录失败："+v, zap.Any(" error:", err))
				return err
			}
		}
	}
	return err
}
