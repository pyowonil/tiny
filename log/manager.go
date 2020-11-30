package log

import "sync"

// Manager is a manager for loggers.
type Manager struct {
	store map[string]*Logger
	mutex *sync.RWMutex
}

// NewManager creates a new Manager.
func NewManager() *Manager {
	manager := &Manager{
		store: map[string]*Logger{},
		mutex: &sync.RWMutex{},
	}
	return manager
}

// GetLogger gets a logger.
func (manager *Manager) GetLogger(name string) *Logger {
	logger, ok := manager.store[name]
	if !ok {
		logger = NewLogger(name, DefaultLevel)
		logger.mutex = manager.mutex
		manager.store[name] = logger
	}
	return logger
}

// SetLevel sets level for all loggers in store.
func (manager *Manager) SetLevel(level Level) {
	for _, logger := range manager.store {
		logger.SetLevel(level)
	}
}
