package db

import "github.com/mgutz/logxi"

type PgxLogger struct {
	log logxi.Logger
}

func (l *PgxLogger) Debug(msg string, ctx ...interface{}) {
	l.log.Debug(msg, ctx...)
}

func (l *PgxLogger) Info(msg string, ctx ...interface{}) {
	l.log.Info(msg, ctx...)
}

func (l *PgxLogger) Warn(msg string, ctx ...interface{}) {
	l.log.Warn(msg, ctx...)
}

func (l *PgxLogger) Error(msg string, ctx ...interface{}) {
	l.log.Error(msg, ctx...)
}
