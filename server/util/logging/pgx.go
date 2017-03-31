package logging

import "go.uber.org/zap"

func NewPGXLog(z *zap.SugaredLogger) *PGXLogWrapper {
	return &PGXLogWrapper{z: z}
}

type PGXLogWrapper struct {
	z *zap.SugaredLogger
}

func (p *PGXLogWrapper) Debug(msg string, ctx ...interface{}) {
	p.z.Debugw(msg, ctx...)
}

func (p *PGXLogWrapper) Info(msg string, ctx ...interface{}) {
	p.z.Infow(msg, ctx...)
}

func (p *PGXLogWrapper) Warn(msg string, ctx ...interface{}) {
	p.z.Warnw(msg, ctx...)
}

func (p *PGXLogWrapper) Error(msg string, ctx ...interface{}) {
	p.z.Errorw(msg, ctx...)
}
