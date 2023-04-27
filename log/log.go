package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/michelemendel/goutils/time"
)

const (
	defaultLogLevel = zapcore.DebugLevel
	stacktraceLevel = zapcore.ErrorLevel
)

var lg *zap.SugaredLogger

func init() {
	initWithConsole(defaultLogLevel)
}

func initWithConsole(loglevel zapcore.Level) {
	consoleConfig := setConfigs(zap.NewProductionEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(consoleConfig)
	core := zapcore.NewTee(zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), loglevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(stacktraceLevel))
	lg = logger.Sugar()
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

	lg = logger.Sugar()
	return lg
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

func SetDebugLevel() {
	initWithConsole(zapcore.DebugLevel)
}

func SetInfoLevel() {
	initWithConsole(zapcore.InfoLevel)
}

func SetWarnLevel() {
	initWithConsole(zapcore.WarnLevel)
}

func SetErrorLevel() {
	initWithConsole(zapcore.ErrorLevel)
}

func SetDPanicLevel() {
	initWithConsole(zapcore.DPanicLevel)
}

func SetPanicLevel() {
	initWithConsole(zapcore.PanicLevel)
}

func SetFatalLevel() {
	initWithConsole(zapcore.FatalLevel)
}

// Debug
func Debug(msg string) {
	lg.Debug(msg)
}
func Debugf(msg string, fields ...interface{}) {
	lg.Debugf(msg, fields...)
}
func Debugw(msg string, keysAndValues ...interface{}) {
	lg.Debugw(msg, keysAndValues...)
}

// Info
func Info(msg string) {
	lg.Info(msg)
}
func Infof(msg string, fields ...interface{}) {
	lg.Infof(msg, fields...)
}
func Infow(msg string, keysAndValues ...interface{}) {
	lg.Infow(msg, keysAndValues...)
}

// Warn
func Warn(msg string) {
	lg.Warn(msg)
}
func Warnf(msg string, fields ...interface{}) {
	lg.Warnf(msg, fields...)
}
func Warnw(msg string, keysAndValues ...interface{}) {
	lg.Infow(msg, keysAndValues...)
}

// Error
func Error(msg string) {
	lg.Error(msg)
}
func Errorf(msg string, fields ...interface{}) {
	lg.Errorf(msg, fields...)
}
func Errorw(msg string, keysAndValues ...interface{}) {
	lg.Errorw(msg, keysAndValues...)
}

// DPanic
func DPanic(msg string) {
	lg.DPanic(msg)
}
func DPanicf(msg string, fields ...interface{}) {
	lg.DPanicf(msg, fields...)
}
func DPanicw(msg string, keysAndValues ...interface{}) {
	lg.DPanicw(msg, keysAndValues...)
}

// Panic
func Panic(msg string) {
	lg.Panic(msg)
}
func Panicf(msg string, fields ...interface{}) {
	lg.Panicf(msg, fields...)
}
func Panicw(msg string, keysAndValues ...interface{}) {
	lg.Panicw(msg, keysAndValues...)
}

// Fatal
func Fatal(msg string) {
	lg.Fatal(msg)
}
func Fatalf(msg string, fields ...interface{}) {
	lg.Fatalf(msg, fields...)
}
func Fatalw(msg string, keysAndValues ...interface{}) {
	lg.Fatalw(msg, keysAndValues...)
}
