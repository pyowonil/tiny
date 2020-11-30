package log

import "testing"

func TestNewLogger(t *testing.T) {
	_ = NewLogger("test", LevelDebug)
}

func TestLogger_GetName(t *testing.T) {
	logger := NewLogger("test", LevelDebug)
	if logger.GetName() != "test" {
		t.Fail()
	}
}

func TestLogger_GetLevel(t *testing.T) {
	logger := NewLogger("test", LevelDebug)
	if logger.GetLevel() != LevelDebug {
		t.Fail()
	}
}

func TestLogger_GetLevelString(t *testing.T) {
	logger := NewLogger("test", LevelDebug)
	if logger.GetLevelString() != "Debug" {
		t.Fail()
	}
}

func TestLogger_SetLevel(t *testing.T) {
	logger := NewLogger("test", LevelDebug)
	logger.SetLevel(LevelError)
	if logger.level != LevelError {
		t.Fail()
	}
}

func TestLogger_Debugf(t *testing.T) {
	logger := NewLogger("test", LevelDebug)
	logger.Debugf("hello world")
}

func TestLogger_Infof(t *testing.T) {
	logger := NewLogger("test", LevelDebug)
	logger.Infof("hello world")
}

func TestLogger_Warnf(t *testing.T) {
	logger := NewLogger("test", LevelDebug)
	logger.Warnf("hello world")
}

func TestLogger_Errorf(t *testing.T) {
	logger := NewLogger("test", LevelDebug)
	logger.Errorf("hello world")
}

func TestLogger_Fatalf(t *testing.T) {
	logger := NewLogger("test", LevelDebug)
	logger.Fatalf("hello world")
}

func TestLogger_Logf(t *testing.T) {
	debugLogger := NewLogger("debug", LevelDebug)
	fatalLogger := NewLogger("fatal", LevelFatal)
	debugLogger.Logf(LevelError, "printed")
	fatalLogger.Logf(LevelError, "not printed")
}
