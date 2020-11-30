package cmd

import "github.com/pyowonil/tiny/log"

// logger is a logger.
var logger *log.Logger

// init initializes this package.
func init() {
	logger = log.DefaultManager.GetLogger("cmd")
	logger.SetLevel(log.LevelFatal)
}
