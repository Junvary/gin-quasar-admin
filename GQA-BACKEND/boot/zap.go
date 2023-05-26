package boot

import (
	elPath "github.com/Junvary/erleng/path"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
)

func Zap() (logger *zap.Logger) {
	zapConfig := global.GqaConfig.Zap
	if err := elPath.CreatePath(zapConfig.Path); err != nil {
		panic(err)
	}
	writeSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   zapConfig.Path + "/" + zapConfig.Filename + "-" + time.Now().Format("20060102150405") + ".log",
		MaxSize:    zapConfig.MaxSize,
		MaxBackups: zapConfig.MaxBackups,
		MaxAge:     zapConfig.MaxAge,
	})
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = CustomTimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	var l = new(zapcore.Level)
	err := l.UnmarshalText([]byte(zapConfig.Level))
	if err != nil {
		panic(err)
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)
	logger = zap.New(core, zap.AddCaller())
	return logger
}

func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.GqaConfig.Zap.Prefix + "-2006/01/02 15:04:05.000"))
}
