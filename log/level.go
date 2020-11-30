package log

import "fmt"

// Level is a log level.
type Level uint8

const (
	// LevelDebug means a debug log level.
	LevelDebug = Level(iota)
	// LevelInfo means a info log level.
	LevelInfo
	// LevelWarn means a warn log level.
	LevelWarn
	// LevelError means a error log level.
	LevelError
	// LevelFatal means a fatal log level.
	LevelFatal
)

// DefaultLevel is a default logger level.
var DefaultLevel = LevelDebug

// mapLevelToString is map for level strings.
var mapLevelToString = map[Level]string{
	LevelDebug: "Debug",
	LevelInfo:  "Info",
	LevelWarn:  "Warn",
	LevelError: "Error",
	LevelFatal: "Fatal",
}

// canLog checks level can be logging.
func (level Level) canLog(loglevel Level) bool {
	return level <= loglevel
}

// DecorateOn is a flag for decorating messages.
var DecorateOn = true

// mapLevelToColorTag is map for color tag.
var mapLevelToColorTag = map[Level]string{
	LevelDebug: "\033[34m", // ANSI Blue
	LevelInfo:  "\033[32m", // ANSI Green
	LevelWarn:  "\033[33m", // ANSI Yellow
	LevelError: "\033[35m", // ANSI Magenta
	LevelFatal: "\033[31m", // ANSI Red
}

// decorate decorates message with colors.
func (level Level) decorate(msg string) string {
	if DecorateOn {
		tag := mapLevelToColorTag[level]
		return fmt.Sprintf("%s%s\033[0m", tag, msg)
	}
	return msg
}
