package debug

import (
	"os"

	"github.com/go-zoox/core-utils/fmt"
)

// EnvKey is the key of environment vars for DEBUG
var EnvKey = "DEBUG"

// IsDebugMode check is now in debug mode
func IsDebugMode() bool {
	return os.Getenv(EnvKey) != ""
}

// Debug prints debug message if it is in debug mode
func Debug(args ...interface{}) {
	Info(args...)
}

// Info prints debug message if it is in debug mode
func Info(args ...interface{}) {
	if !IsDebugMode() {
		return
	}

	argsx := append([]interface{}{"[debug]"}, args...)

	fmt.PrintJSON(argsx...)
}
