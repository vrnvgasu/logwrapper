package logwrapper

import (
	"context"

	formatter "github.com/fabienm/go-logrus-formatters"
	"github.com/sirupsen/logrus"
)

var Logger *StandardLogger

type Payload struct {
	id   string
	op   string
	pack string
}

func NewPayload() *Payload {
	return &Payload{}
}

type StandardLogger struct {
	*logrus.Logger
}

func NewLogger(level logrus.Level, hooks []logrus.Hook) *StandardLogger {
	baseLogger := logrus.New()
	standardLogger := &StandardLogger{baseLogger}

	gelfFmt := formatter.NewGelf("service-name")
	standardLogger.SetFormatter(gelfFmt)
	standardLogger.SetLevel(level)

	for _, hook := range hooks {
		standardLogger.AddHook(hook)
	}

	return standardLogger
}

func (p *Payload) Package(v string) *Payload {
	(*p).pack = v
	return p
}

func (p *Payload) Op(v string) *Payload {
	(*p).op = v
	return p
}

func (p *Payload) CtxID(ctx context.Context, key string) *Payload {
	(*p).id = ctx.Value(key).(string)
	return p
}

func (p *Payload) toMap() map[string]interface{} {
	res := map[string]interface{}{}
	res["id"] = p.id
	res["op"] = p.op
	res["pack"] = p.pack
	return res
}

func (l *StandardLogger) Payload(payload *Payload) *logrus.Entry {
	return l.WithFields(payload.toMap())
}
