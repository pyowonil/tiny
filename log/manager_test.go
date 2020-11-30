package log

import "testing"

func TestNewManager(t *testing.T) {
	_ = NewManager()
}

func TestManager_GetLogger(t *testing.T) {
	manager := NewManager()
	l1 := manager.GetLogger("test")
	l2 := manager.GetLogger("test")
	if l1 != l2 {
		t.Fail()
	}
}

func TestManager_SetLevel(t *testing.T) {
	manager := NewManager()
	l1 := manager.GetLogger("test1")
	l2 := manager.GetLogger("test2")
	manager.SetLevel(LevelFatal)
	if l1.GetLevel() != LevelFatal || l2.GetLevel() != LevelFatal {
		t.Fail()
	}
	manager.SetLevel(LevelDebug)
	if l1.GetLevel() != LevelDebug || l2.GetLevel() != LevelDebug {
		t.Fail()
	}
}
