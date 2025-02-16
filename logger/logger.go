package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumber "gopkg.in/natefinch/lumberjack.v2"
)

var _ log.Logger = (*ZapLogger)(nil)

type ZapLogger struct {
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
}

func InitLog() *ZapLogger {
	var coreArr []zapcore.Core
	var logger *zap.Logger
	// 获取编码器
	// NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoderConfig := zap.NewProductionEncoderConfig()
	// 指定时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 按级别显示不同颜色，不需要的话取值zapcore.CapitalLevelEncoder就可以了
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//显示完整文件路径
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder

	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	// 日志级别
	// error级别
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})
	// info和debug级别,debug级别是最低的
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev < zap.ErrorLevel && lev >= zap.DebugLevel
	})

	_, currentFile, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(currentFile)
	currentDir = filepath.Dir(currentDir)

	fmt.Println(currentDir)

	// info文件writeSyncer
	infoFileWriteSyncer := zapcore.AddSync(&lumber.Logger{
		Filename:   currentDir + "/logs/info.log", // 日志文件存放目录，
		MaxSize:    200,                           // 文件大小限制,单位MB
		MaxBackups: 30,                            // 最大保留日志文件数量
		MaxAge:     7,                             // 日志文件保留天数
		Compress:   false,
	})
	infoFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer, zapcore.AddSync(os.Stdout)), lowPriority)

	// error文件writeSyncer
	errorFileWriteSyncer := zapcore.AddSync(&lumber.Logger{
		Filename:   currentDir + "/logs/error.log", // 日志文件存放目录，
		MaxSize:    200,                            // 文件大小限制,单位MB
		MaxBackups: 30,                             // 最大保留日志文件数量
		MaxAge:     7,                              // 日志文件保留天数
		Compress:   false,
	})
	errorFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errorFileWriteSyncer, zapcore.AddSync(os.Stdout)), highPriority)

	coreArr = append(coreArr, infoFileCore)
	coreArr = append(coreArr, errorFileCore)

	// zap.AddCaller()为显示文件名和行号，可省略
	logger = zap.New(zapcore.NewTee(coreArr...), zap.AddCaller(), zap.AddCallerSkip(1))
	sugar := logger.Sugar()
	logger.Info("init logger success")

	return &ZapLogger{Logger: logger, Sugar: sugar}

}

func (zapl *ZapLogger) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		zapl.Logger.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}

	var data []zap.Field
	for i := 0; i < len(keyvals); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), keyvals[i+1]))
	}

	switch level {
	case log.LevelDebug:
		zapl.Logger.Debug("", data...)
	case log.LevelInfo:
		zapl.Logger.Info("", data...)
	case log.LevelWarn:
		zapl.Logger.Warn("", data...)
	case log.LevelError:
		zapl.Logger.Error("", data...)
	case log.LevelFatal:
		zapl.Logger.Fatal("", data...)
	}
	return nil
}

func (zapl *ZapLogger) Infof(s string, v ...interface{}) {
	zapl.Sugar.Infof(s, v...)
}

func (zapl *ZapLogger) Infow(s string, v ...interface{}) {
	zapl.Sugar.Infow(s, v...)
}

func (zapl *ZapLogger) Info(v ...interface{}) {
	zapl.Sugar.Info(v...)
}

func (zapl *ZapLogger) Debugf(s string, v ...interface{}) {
	zapl.Sugar.Debugf(s, v...)
}

func (zapl *ZapLogger) Debugw(s string, v ...interface{}) {
	zapl.Sugar.Debugw(s, v...)
}

func (zapl *ZapLogger) Debug(v ...interface{}) {
	zapl.Sugar.Debug(v...)
}

func (zapl *ZapLogger) Errorf(s string, v ...interface{}) {
	zapl.Sugar.Errorf(s, v...)
}

func (zapl *ZapLogger) InitErrf(s string, v ...interface{}) {
	s = "[init err] " + s
	zapl.Sugar.Errorf(s, v...)
}
func (zapl *ZapLogger) InitInfof(s string, v ...interface{}) {
	s = "[init success] " + s
	zapl.Sugar.Infof(s, v...)
}

func (zapl *ZapLogger) ErrorfInBiz(s string, v ...interface{}) {
	s = "[biz] " + s
	zapl.Sugar.Errorf(s, v...)
}
func (zapl *ZapLogger) ErrorfInData(s string, v ...interface{}) {
	s = "[data] " + s
	zapl.Sugar.Errorf(s, v...)
}
func (zapl *ZapLogger) WarnfInData(s string, v ...interface{}) {
	s = "[data] " + s
	zapl.Sugar.Warnf(s, v...)
}

func (zapl *ZapLogger) ErrorfInDataKafka(s string, v ...interface{}) {
	s = "[data kafka] " + s
	zapl.Sugar.Errorf(s, v...)
}
func (zapl *ZapLogger) InfofInDataKafka(s string, v ...interface{}) {
	s = "[data kafka] " + s
	zapl.Sugar.Infof(s, v...)
}
func (zapl *ZapLogger) ErrorInDataKafka(s string) {
	s = "[data kafka] " + s
	zapl.Sugar.Error(s)
}
func (zapl *ZapLogger) InfoInDataKafka(s string) {
	s = "[data kafka] " + s
	zapl.Sugar.Info(s)
}

func (zapl *ZapLogger) ErrorfInService(s string, v ...interface{}) {
	s = "[service] " + s
	zapl.Sugar.Errorf(s, v...)
}
func (zapl *ZapLogger) InfofInBiz(s string, v ...interface{}) {
	s = "[biz] " + s
	zapl.Sugar.Infof(s, v...)
}
func (zapl *ZapLogger) InfofInData(s string, v ...interface{}) {
	s = "[data] " + s
	zapl.Sugar.Infof(s, v...)
}
func (zapl *ZapLogger) InfoInData(s string) {
	s = "[data] " + s
	zapl.Sugar.Info(s)
}

func (zapl *ZapLogger) InfofInService(s string, v ...interface{}) {
	s = "[service] " + s
	zapl.Sugar.Infof(s, v...)
}

func (zapl *ZapLogger) LogDurationInData(s string) {
	s = "[Data Duration] | " + s
	zapl.Sugar.Info(s)
}

func (zapl *ZapLogger) LogDurationInBiz(s string, v ...interface{}) {
	s = "[Biz Duration] | " + s
	zapl.Sugar.Infof(s, v...)
}
func (zapl *ZapLogger) LogDurationInService(s string, v ...interface{}) {
	s = "[Service Duration] | " + s
	zapl.Sugar.Infof(s, v...)
}

func (zapl *ZapLogger) Errorw(s string, v ...interface{}) {
	zapl.Sugar.Errorw(s, v...)
}

func (zapl *ZapLogger) Error(v ...interface{}) {
	zapl.Sugar.Error(v...)
}

func (zapl *ZapLogger) Fatalf(s string, v ...interface{}) {
	zapl.Sugar.Fatalf(s, v...)
}

func (zapl *ZapLogger) Fatalw(s string, v ...interface{}) {
	zapl.Sugar.Fatalw(s, v...)
}

func (zapl *ZapLogger) Fatal(v ...interface{}) {
	zapl.Sugar.Error(v...)
}
func (zapl *ZapLogger) Sync() {
	_ = zapl.Logger.Sync()
}
