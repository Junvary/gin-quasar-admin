package boot

import (
	"fmt"
	"gin-quasar-admin/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

func Zap() (logger *zap.Logger) {
	zapCfg := global.GqaConfig.Zap
	_, pathErr := os.Stat(zapCfg.Filepath)
	if pathErr != nil {
		fmt.Printf("开始创建 %v 日志文件夹\n", zapCfg.Filepath)
		_ = os.Mkdir(zapCfg.Filepath, os.ModePerm)
	}
	writeSyncer := getLogWriter(zapCfg.Filepath, zapCfg.Filename, zapCfg.MaxSize, zapCfg.MaxBackups, zapCfg.MaxAge)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err := l.UnmarshalText([]byte(zapCfg.Level))
	if err != nil {
		return
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)

	logger = zap.New(core, zap.AddCaller())
	return logger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.EncodeTime = CustomTimeEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filepath string, filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filepath + "/" + filename + "-" + time.Now().Format("20060102150405" )+ ".log",
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.GqaConfig.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}
