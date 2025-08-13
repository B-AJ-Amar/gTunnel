package logger

import (
	"os"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

var (
	Logger *logrus.Logger
)

type LogLevel string

const (
	LevelTrace LogLevel = "trace"
	LevelDebug LogLevel = "debug"
	LevelInfo  LogLevel = "info"
	LevelWarn  LogLevel = "warn"
	LevelError LogLevel = "error"
	LevelFatal LogLevel = "fatal"
	LevelPanic LogLevel = "panic"
)

func Init(level LogLevel, enableColors bool) {
	Logger = logrus.New()

	Logger.SetOutput(os.Stdout)

	switch level {
	case LevelTrace:
		Logger.SetLevel(logrus.TraceLevel)
	case LevelDebug:
		Logger.SetLevel(logrus.DebugLevel)
	case LevelInfo:
		Logger.SetLevel(logrus.InfoLevel)
	case LevelWarn:
		Logger.SetLevel(logrus.WarnLevel)
	case LevelError:
		Logger.SetLevel(logrus.ErrorLevel)
	case LevelFatal:
		Logger.SetLevel(logrus.FatalLevel)
	case LevelPanic:
		Logger.SetLevel(logrus.PanicLevel)
	default:
		Logger.SetLevel(logrus.InfoLevel)
	}

	if enableColors {
		Logger.SetFormatter(&logrus.TextFormatter{
			ForceColors:     true,
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			PadLevelText:    true,
		})
	} else {
		Logger.SetFormatter(&logrus.TextFormatter{
			DisableColors:   true,
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		})
	}
}

func GetLogger() *logrus.Logger {
	if Logger == nil {
		Init(LevelInfo, true) // Default initialization
	}
	return Logger
}

func Trace(args ...interface{}) {
	GetLogger().Trace(args...)
}

func Tracef(format string, args ...interface{}) {
	GetLogger().Tracef(format, args...)
}

func Debug(args ...interface{}) {
	GetLogger().Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	GetLogger().Debugf(format, args...)
}

func Info(args ...interface{}) {
	GetLogger().Info(args...)
}

func Infof(format string, args ...interface{}) {
	GetLogger().Infof(format, args...)
}

func Warn(args ...interface{}) {
	GetLogger().Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	GetLogger().Warnf(format, args...)
}

func Error(args ...interface{}) {
	GetLogger().Error(args...)
}

func Errorf(format string, args ...interface{}) {
	GetLogger().Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	GetLogger().Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	GetLogger().Fatalf(format, args...)
}

func Panic(args ...interface{}) {
	GetLogger().Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	GetLogger().Panicf(format, args...)
}

func WithField(key string, value interface{}) *logrus.Entry {
	return GetLogger().WithField(key, value)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return GetLogger().WithFields(fields)
}

// Critical displays a serious warning message in bold red color
func Critical(args ...interface{}) {
	red := color.New(color.FgRed, color.Bold)
	red.Print("⚠️  WARNING: ")
	red.Print(args...)
	red.Print("\n")
	// GetLogger().Warn(args...)
}

// Criticalf displays a serious warning message with formatting in bold red color
func Criticalf(format string, args ...interface{}) {
	red := color.New(color.FgRed, color.Bold)
	red.Printf("⚠️  WARNING: ")
	red.Printf(format, args...)
	red.Print("\n")
	// GetLogger().Warnf(format, args...)
}
