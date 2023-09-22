package debug

import (
	"os"

	"github.com/go-zoox/core-utils/fmt"
)

// Debugger is a interface for debug
type Debugger interface {
	IsDebugMode(compare ...func(envValue string) bool) bool
	//
	Debug(args ...interface{})
	Info(args ...interface{})
	//
	// SetEnvKey(envKey string)
	SetLogger(func(args ...interface{}) error)
}

type debugger struct {
	envKey string
	logger func(args ...interface{}) error
}

// New returns a new Debugger
func New(envKey string, logger ...func(args ...interface{}) error) Debugger {
	loggerX := fmt.PrintJSON
	if len(logger) > 0 {
		loggerX = logger[0]
	}

	return &debugger{
		envKey: envKey,
		logger: loggerX,
	}
}

// IsDebugMode check is now in debug mode
func (d *debugger) IsDebugMode(compare ...func(envValue string) bool) bool {
	if len(compare) > 0 && compare[0] != nil {
		return compare[0](os.Getenv(d.envKey))
	}

	return os.Getenv(d.envKey) != ""
}

// Debug prints debug message if it is in debug mode
func (*debugger) Debug(args ...interface{}) {
	Info(args...)
}

// Info prints debug message if it is in debug mode
func (d *debugger) Info(args ...interface{}) {
	if !IsDebugMode() {
		return
	}

	argsx := append([]interface{}{"[debug]"}, args...)

	d.logger(argsx...)
}

// // SetEnvKey is the key of environment vars for DEBUG
// func (d *debugger) SetEnvKey(envKey string) {
// 	d.envKey = envKey
// }

// SetLogger sets the logger
func (d *debugger) SetLogger(logger func(args ...interface{}) error) {
	d.logger = logger
}
