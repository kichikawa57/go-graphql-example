package logger

import (
	"cloud.google.com/go/logging"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	l *zap.Logger
	s *zap.SugaredLogger

	Debug     func(args ...interface{})
	Info      func(args ...interface{})
	Warn      func(args ...interface{})
	Error     func(args ...interface{})
	Debugf    func(tmpl string, args ...interface{})
	Infof     func(tmpl string, args ...interface{})
	Warnf     func(tmpl string, args ...interface{})
	Errorf    func(tmpl string, args ...interface{})
	InfoQuick func(msg string, fields ...zap.Field)
	WarnQuick func(msg string, fields ...zap.Field)
)

func SetUp() error {

	conf := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "console",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:      "severity",
			NameKey:       "logger",
			StacktraceKey: "stack_trace",
			TimeKey:       "time",
			MessageKey:    "message",
			CallerKey:     "caller",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeTime:    zapcore.RFC3339NanoTimeEncoder,
			EncodeLevel:   levelEncode,
			EncodeCaller:  zapcore.ShortCallerEncoder,
		},
	}
	var err error
	l, err = conf.Build()
	if err != nil {
		return err
	}
	s = l.Sugar()

	Debug = s.Debug
	Info = s.Info
	Warn = s.Warn
	Error = s.Error
	Debugf = s.Debugf
	Infof = s.Infof
	Warnf = s.Warnf
	Errorf = s.Errorf
	InfoQuick = l.Info
	WarnQuick = l.Warn

	return nil
}

func levelEncode(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch l {
	case zapcore.DebugLevel:
		enc.AppendString(logging.Debug.String())
	case zapcore.InfoLevel:
		enc.AppendString(logging.Info.String())
	case zapcore.WarnLevel:
		enc.AppendString(logging.Warning.String())
	case zapcore.ErrorLevel:
		enc.AppendString(logging.Error.String())
	default:
		enc.AppendString(logging.Critical.String())
	}
}
