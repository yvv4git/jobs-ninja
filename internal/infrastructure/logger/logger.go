package logger

import (
	"fmt"
	"log/slog"
	"os"
)

func SetupDefaultLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
}

func SetupLoggerWithLevel(level slog.Level) *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))
}

type LogLevel string

const (
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelWarn  LogLevel = "warn"
	LogLevelError LogLevel = "error"
)

func ParseLogLevel(levelStr LogLevel) slog.Level {
	switch levelStr {
	case LogLevelDebug:
		return slog.LevelDebug
	case LogLevelInfo:
		return slog.LevelInfo
	case LogLevelWarn:
		return slog.LevelWarn
	case LogLevelError:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func (l *LogLevel) UnmarshalText(text []byte) error {
	level := LogLevel(text)
	switch level {
	case LogLevelDebug, LogLevelInfo, LogLevelWarn, LogLevelError:
		*l = level
		return nil
	default:
		return fmt.Errorf("invalid log level: %s", text)
	}
}
