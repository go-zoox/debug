package debug

// EnvKey is the key of environment vars for DEBUG
var EnvKey = "LOG_LEVEL"

var global = New(EnvKey)

// IsDebugMode check is now in debug mode
func IsDebugMode() bool {
	return global.IsDebugMode()
}

// Debug prints debug message if it is in debug mode
func Debug(args ...interface{}) {
	global.Info(args...)
}

// Info prints debug message if it is in debug mode
func Info(args ...interface{}) {
	argsx := append([]interface{}{"[debug]"}, args...)
	global.Info(argsx...)
}
