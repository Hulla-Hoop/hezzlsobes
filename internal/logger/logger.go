package logger

import (
	formatter "github.com/fabienm/go-logrus-formatters"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	L *logrus.Logger
}

func New() *Logger {
	logger := logrus.New()
	rr := formatter.NewGelf("service-second")
	logger.SetFormatter(rr)
	logger.SetLevel(logrus.DebugLevel)
	return &Logger{
		L: logger,
	}
}

func (l *Logger) ErrorGH(reqId string, message error) {
	err := errors.WithStack(message)
	l.L.WithField("GetHash", reqId).Errorf("%+v", err)
}

func (l *Logger) ErrorCH(reqId string, message error) {
	err := errors.WithStack(message)
	l.L.WithField("CreateHash", reqId).Errorf("%+v", err)
}
