package plugins

import (
	"os"

	joonix "github.com/joonix/log"
	"github.com/sirupsen/logrus"
)

var SysLogger *Logger

type Logger struct {
	*logrus.Logger
}

func NewLogger(env *Env) *Logger {
	l := &logrus.Logger{
		Out:          os.Stderr,
		Formatter:    new(logrus.TextFormatter),
		Hooks:        make(logrus.LevelHooks),
		Level:        logrus.DebugLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}

	if env.GetEnv("ENVIRONMENT") == "production" {
		l.SetFormatter(joonix.NewFormatter())
		l.SetReportCaller(true)
	}

	switch level := env.GetEnv("LOG_LEVEL"); level {
	case "trace":
		l.SetLevel(logrus.TraceLevel)
	case "debug":
		l.SetLevel(logrus.DebugLevel)
	case "info":
		l.SetLevel(logrus.InfoLevel)
	default:
		if env.GetEnv("ENVIRONMENT") == "production" {
			l.SetLevel(logrus.InfoLevel)
		} else {
			l.SetLevel(logrus.DebugLevel)
		}
	}
	l.SetOutput(os.Stdout)

	logger := &Logger{
		Logger: l,
	}

	SysLogger = logger
	return logger
}
