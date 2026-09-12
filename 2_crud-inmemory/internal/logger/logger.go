package logger

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"
)

type Logger struct {
	*slog.Logger

	levelVar     *slog.LevelVar
	multiHandler []slog.Handler

	file *os.File
	dir  *string
}

// Save config in struct.
// Set levelVar for logger:
// If config == nil levelVar = INFO
func NewBlank(config *Config) *Logger {
	l := &Logger{
		levelVar: new(slog.LevelVar),
	}

	if config == nil {
		l.levelVar.Set(slog.LevelInfo)
		return l
	}

	l.dir = &config.Dir

	if err := l.SetLevel(config.Level); err != nil {
		fmt.Println(err)
	}
	return l
}

func (l *Logger) WithCTextHandlerConsoleOut() *Logger {
	opt := &slog.HandlerOptions{
		Level: l.levelVar,
	}

	myHandler := newCustomHandler(os.Stdout, opt)

	l.multiHandler = append(l.multiHandler, myHandler)
	return l
}

func (l *Logger) WithCTextHandlerFileOut() *Logger {
	opt := &slog.HandlerOptions{
		Level: l.levelVar,
	}

	fileName := time.Now().UTC().Format("2006-01-02T15_04_05Z07_00")
	path := fmt.Sprintf("./%s/", *l.dir)
	err := os.MkdirAll(path, 0755)
	if err != nil {
		panic("Trouble with create folders")
	}

	name := fmt.Sprintf("./%s/%s.log", *l.dir, fileName)
	l.file, err = os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("Trouble with open file for logs")
	}

	myHandler := newCustomHandler(l.file, opt)

	l.multiHandler = append(l.multiHandler, myHandler)
	return l
}

func (l *Logger) Construct() {
	l.Logger = slog.New(slog.NewMultiHandler(l.multiHandler...))
}

func (l *Logger) SetLevel(level string) error {
	level = strings.ToLower(level)
	switch level {
	case "debug":
		l.levelVar.Set(slog.LevelDebug)
	case "info":
		l.levelVar.Set(slog.LevelInfo)
	case "warn":
		l.levelVar.Set(slog.LevelWarn)
	case "error":
		l.levelVar.Set(slog.LevelError)
	default:
		l.levelVar.Set(slog.LevelDebug)
		return fmt.Errorf("Invalid logLevel value: %s, level changed on DEBUG", level)
	}
	return nil
}

func (l *Logger) CloseFile() {
	if l.file != nil {
		l.file.Sync()
		l.file.Close()
	}
}
