package log

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/michelemendel/goutils/time"
)

const (
	defaultLogLevel = zapcore.DebugLevel
	stacktraceLevel = zapcore.ErrorLevel
)

func InitWithConsole(loglevel string) *zap.SugaredLogger {
	lglvl := getLoglevel(strings.ToLower(loglevel))
	consoleConfig := setConfigs(zap.NewProductionEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(consoleConfig)
	core := zapcore.NewTee(zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), lglvl))
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(stacktraceLevel))
	return logger.Sugar()
}

func initWithFile(filename string) *zap.SugaredLogger {
	defaultLogLevel := zapcore.DebugLevel

	consoleConfig := setConfigs(zap.NewProductionEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(consoleConfig)

	fileConfig := setConfigs(zap.NewProductionEncoderConfig())
	fileEncoder := zapcore.NewJSONEncoder(fileConfig)

	logFile := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    1, // megabytes
		MaxBackups: 2,
		MaxAge:     1, // days
	}

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(logFile), defaultLogLevel),
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return logger.Sugar()
}

func setConfigs(cfg zapcore.EncoderConfig) zapcore.EncoderConfig {
	cfg.EncodeLevel = encodeLevel()
	cfg.EncodeTime = zapcore.TimeEncoderOfLayout(time.TimestampFormatLayout)
	cfg.CallerKey = "c"
	cfg.LevelKey = "l"
	cfg.MessageKey = "m"
	// cfg.StacktraceKey = stacktrace // Use "" to disable stacktrace
	cfg.TimeKey = "t"
	return cfg
}

func getLoglevel(loglevel string) zapcore.Level {
	switch loglevel {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return defaultLogLevel
	}
}

// Zap and GCP: https://github.com/uber-go/zap/issues/1095
func encodeLevel() zapcore.LevelEncoder {
	return func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		switch l {
		case zapcore.DebugLevel:
			enc.AppendString("DEBUG")
		case zapcore.InfoLevel:
			enc.AppendString("INFO")
		case zapcore.WarnLevel:
			enc.AppendString("WARNING")
		case zapcore.ErrorLevel:
			enc.AppendString("ERROR")
		case zapcore.DPanicLevel:
			enc.AppendString("CRITICAL")
		case zapcore.PanicLevel:
			enc.AppendString("ALERT")
		case zapcore.FatalLevel:
			enc.AppendString("EMERGENCY")
		}
	}
}

// Set log level

func SetDebugLevel() *zap.SugaredLogger {
	return InitWithConsole("DEBUG")
}

func SetInfoLevel() *zap.SugaredLogger {
	return InitWithConsole("INFO")
}

func SetWarnLevel() *zap.SugaredLogger {
	return InitWithConsole("WARN")
}

func SetErrorLevel() *zap.SugaredLogger {
	return InitWithConsole("ERROR")
}

func SetDPanicLevel() *zap.SugaredLogger {
	return InitWithConsole("DPANIC")
}

func SetPanicLevel() *zap.SugaredLogger {
	return InitWithConsole("PANIC")
}

func SetFatalLevel() *zap.SugaredLogger {
	return InitWithConsole("FATAL")
}
