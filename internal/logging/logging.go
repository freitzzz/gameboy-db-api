package logging

var loggers = []Logger{}
var includeDebugLogs = true

type Logger interface {
	Info(fmt string, s ...any)
	Warning(fmt string, s ...any)
	Error(fmt string, s ...any)
	Fatal(fmt string, s ...any)
	Debug(fmt string, s ...any)
}

func AddLogger(logger Logger) {
	loggers = append(loggers, logger)
}

func DisableDebugLogs() {
	includeDebugLogs = false
}

func Info(fmt string, s ...any) {
	for _, l := range loggers {
		callLogger(l.Info, fmt, s...)
	}
}

func Warning(fmt string, s ...any) {
	for _, l := range loggers {
		callLogger(l.Warning, fmt, s...)
	}
}

func Error(fmt string, s ...any) {
	for _, l := range loggers {
		callLogger(l.Error, fmt, s...)
	}
}

func Fatal(fmt string, s ...any) {
	for _, l := range loggers {
		callLogger(l.Fatal, fmt, s...)
	}
}

func Debug(fmt string, s ...any) {
	if !includeDebugLogs {
		return
	}

	for _, l := range loggers {
		callLogger(l.Debug, fmt, s...)
	}
}

func callLogger(cb func(string, ...any), fmt string, s ...any) {
	cb(fmt, s...)
}
