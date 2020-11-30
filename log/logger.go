package log

import (
	"fmt"
	"sync"
	"time"
)

// Logger is a tiny logger.
type Logger struct {
	name  string
	level Level
	mutex *sync.RWMutex
}

// NewLogger creates a new Logger.
func NewLogger(name string, level Level) *Logger {
	logger := &Logger{
		name:  name,
		level: level,
		mutex: &sync.RWMutex{},
	}
	return logger
}

// GetName gets logger's name.
func (logger *Logger) GetName() string {
	return logger.name
}

// GetLevel gets logger's level.
func (logger *Logger) GetLevel() Level {
	logger.mutex.RLock()
	defer logger.mutex.RUnlock()
	return logger.level
}

// GetLevelString gets logger's level formed string.
func (logger *Logger) GetLevelString() string {
	return mapLevelToString[logger.GetLevel()]
}

// SetLevel sets logger's level.
func (logger *Logger) SetLevel(level Level) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()
	logger.level = level
}

// Debugf logs debug level messages with custom format.
func (logger *Logger) Debugf(format string, args ...interface{}) {
	logger.Logf(LevelDebug, format, args...)
}

// Infof logs info level messages with custom format.
func (logger *Logger) Infof(format string, args ...interface{}) {
	logger.Logf(LevelInfo, format, args...)
}

// Warnf logs warn level messages with custom format.
func (logger *Logger) Warnf(format string, args ...interface{}) {
	logger.Logf(LevelWarn, format, args...)
}

// Errorf logs error level messages with custom format.
func (logger *Logger) Errorf(format string, args ...interface{}) {
	logger.Logf(LevelError, format, args...)
}

// Fatalf logs fatal level messages with custom format.
func (logger *Logger) Fatalf(format string, args ...interface{}) {
	logger.Logf(LevelFatal, format, args...)
}

// Logf logs messages with custom format.
func (logger *Logger) Logf(level Level, format string, args ...interface{}) {
	if logger.GetLevel().canLog(level) {
		stamp := logger.genStamp(level)
		msg := fmt.Sprintf(format, args...)
		envelope := fmt.Sprintf("%s %s\n", stamp, msg)
		decorated := level.decorate(envelope)
		logger.mutex.Lock()
		defer logger.mutex.Unlock()
		fmt.Print(decorated)
	}
}

// genStamp generates a log stamp.
func (logger *Logger) genStamp(level Level) string {
	levelstr := mapLevelToString[level]
	ts := time.Unix(0, time.Now().UnixNano())
	name := logger.GetName()
	return fmt.Sprintf("[%s] %v -%s-", levelstr, ts, name)
}
