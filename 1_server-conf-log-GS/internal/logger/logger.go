package logger

import (
	"fmt"
	"log/slog"
	"os"
	"path"
	"strings"
	"time"
)

type Logger struct {
	*slog.Logger
	level *slog.LevelVar
	file  *os.File

	textHandlers []slog.Handler
	config       *Config
	// options
}

func (l *Logger) WithLevelFromConfig() *Logger {
	l.level = new(slog.LevelVar)
	startloggerLevel, err := getLevelFromConfig(l.config.Level)
	if err != nil {
		fmt.Println(err)
	}
	l.level.Set(startloggerLevel)
	return l
}

func (l *Logger) WithFileHandlerFromConfig() *Logger {
	foolPath := path.Join(os.Getenv("PROJECT_ROOT"), l.config.Dir)
	err := os.MkdirAll(foolPath, 0755)
	if err != nil {
		fmt.Println("Dir for logs not created")
		os.Exit(1)
	}

	fileName := fmt.Sprintf("%s/%s.txt", foolPath, time.Now().Format("2006-01-02_15-04-05.00"))

	l.file, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)

	opt := &slog.HandlerOptions{
		Level: l.level,
	}

	fileHandler := slog.NewTextHandler(l.file, opt)
	l.textHandlers = append(l.textHandlers, fileHandler)
	return l
}

func (l *Logger) WithSTDOutHandler() *Logger {
	opt := &slog.HandlerOptions{
		Level: l.level,
	}
	outputHandler := slog.NewTextHandler(os.Stdout, opt)
	l.textHandlers = append(l.textHandlers, outputHandler)
	return l
}

func NewLogger() *Logger {
	config, err := NewLoggerConfig()
	if err != nil {
		fmt.Println("Env data for logger config invalid")
		os.Exit(1)
	}

	return &Logger{
		config: config,
	}
}

func (l *Logger) Exec() {
	multihandler := slog.NewMultiHandler(l.textHandlers...)
	l.Logger = slog.New(multihandler)
}

func (l *Logger) CloseFile() {
	l.file.Sync()
	_ = l.file.Close()
}

func (l *Logger) SetLevel(level slog.Level) {
	l.level.Set(level)
}

func getLevelFromConfig(level string) (slog.Level, error) {
	level = strings.ToLower(level)
	switch level {
	case "debug":
		return slog.LevelDebug, nil
	case "info":
		return slog.LevelInfo, nil
	case "warn":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	default:
		return slog.LevelDebug, fmt.Errorf("Unexpected slog level: %s", level)
	}
}
