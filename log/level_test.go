package log

import "testing"

func TestLevel_canLog(t *testing.T) {
	testHelper := func(base, target Level, expect bool) {
		baseStr := mapLevelToString[base]
		targetStr := mapLevelToString[target]
		if ok := base.canLog(target); ok != expect {
			t.Errorf("Level%s.canLog(Level%s) - expect: %v, actual: %v",
				baseStr, targetStr, expect, ok)
		}
	}

	// LevelDebug
	testHelper(LevelDebug, LevelDebug, true)
	testHelper(LevelDebug, LevelInfo, true)
	testHelper(LevelDebug, LevelWarn, true)
	testHelper(LevelDebug, LevelError, true)
	testHelper(LevelDebug, LevelFatal, true)

	// LevelInfo
	testHelper(LevelInfo, LevelDebug, false)
	testHelper(LevelInfo, LevelInfo, true)
	testHelper(LevelInfo, LevelWarn, true)
	testHelper(LevelInfo, LevelError, true)
	testHelper(LevelInfo, LevelFatal, true)

	// LevelWarn
	testHelper(LevelWarn, LevelDebug, false)
	testHelper(LevelWarn, LevelInfo, false)
	testHelper(LevelWarn, LevelWarn, true)
	testHelper(LevelWarn, LevelError, true)
	testHelper(LevelWarn, LevelFatal, true)

	// LevelError
	testHelper(LevelError, LevelDebug, false)
	testHelper(LevelError, LevelInfo, false)
	testHelper(LevelError, LevelWarn, false)
	testHelper(LevelError, LevelError, true)
	testHelper(LevelError, LevelFatal, true)

	// LevelFatal
	testHelper(LevelFatal, LevelDebug, false)
	testHelper(LevelFatal, LevelInfo, false)
	testHelper(LevelFatal, LevelWarn, false)
	testHelper(LevelFatal, LevelError, false)
	testHelper(LevelFatal, LevelFatal, true)
}

func TestLevel_deocrate(t *testing.T) {
	DecorateOn = true
	decorated := LevelDebug.decorate("")
	if decorated != "\033[34m\033[0m" {
		t.Fail()
	}
	DecorateOn = false
	decorated = LevelDebug.decorate("")
	if decorated != "" {
		t.Fail()
	}
}
