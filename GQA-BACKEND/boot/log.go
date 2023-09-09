package boot

import (
	elPath "github.com/Junvary/erleng/path"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"gopkg.in/natefinch/lumberjack.v2"
	"log/slog"
	"time"
)

func Log() *slog.Logger {
	logConfig := global.GqaConfig.Log
	if err := elPath.CreatePath(logConfig.Path); err != nil {
		panic(err)
	}
	lj := &lumberjack.Logger{
		Filename:   logConfig.Path + "/" + logConfig.Filename + "-" + time.Now().Format("20060102150405") + ".log",
		LocalTime:  true,
		MaxSize:    logConfig.MaxSize,
		MaxBackups: logConfig.MaxBackups,
		MaxAge:     logConfig.MaxAge,
	}
	defer lj.Close()
	sl := slog.NewJSONHandler(lj, nil)
	logger := slog.New(sl)
	return logger
}
