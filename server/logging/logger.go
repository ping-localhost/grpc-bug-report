package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/grpclog"
)

// How many callers do we need to skip for proper caller
const callDepth = 3

// Taken from grpc/grpclog and exposed
const (
	// Indicates Info severity.
	InfoLog int = iota
	// Indicates Warning severity.
	WarningLog
	// Indicates Error severity.
	ErrorLog
	// Indicates Fatal severity.
	FatalLog
)

// A special logger to replace the default grpclog one
type grpcLogger struct {
	core          zapcore.Core
	sugaredLogger *zap.SugaredLogger
}

func (m grpcLogger) Info(args ...interface{}) {
	if !m.V(InfoLog) {
		return
	}

	m.sugaredLogger.Info(args)
}

func (m grpcLogger) Infoln(args ...interface{}) {
	if !m.V(InfoLog) {
		return
	}

	m.sugaredLogger.Info(args)
}

func (m grpcLogger) Infof(format string, args ...interface{}) {
	if !m.V(InfoLog) {
		return
	}

	m.sugaredLogger.Infof(format, args...)
}

func (m grpcLogger) Warning(args ...interface{}) {
	m.sugaredLogger.Warn(args)
}

func (m grpcLogger) Warningln(args ...interface{}) {
	m.sugaredLogger.Warn(args)
}

func (m grpcLogger) Warningf(format string, args ...interface{}) {
	m.sugaredLogger.Warnf(format, args...)
}

func (m grpcLogger) Error(args ...interface{}) {
	m.sugaredLogger.Error(args)
}

func (m grpcLogger) Errorln(args ...interface{}) {
	m.sugaredLogger.Error(args)
}

func (m grpcLogger) Errorf(format string, args ...interface{}) {
	m.sugaredLogger.Errorf(format, args...)
}

func (m grpcLogger) Fatal(args ...interface{}) {
	// Don't actually call fatal, as that will cause zap to shutdown the application
	m.sugaredLogger.Error(args)
}

func (m grpcLogger) Fatalln(args ...interface{}) {
	// Don't actually call fatal, as that will cause zap to shutdown the application
	m.sugaredLogger.Error(args)
}

func (m grpcLogger) Fatalf(format string, args ...interface{}) {
	// Don't actually call fatal, as that will cause zap to shutdown the application
	m.sugaredLogger.Errorf(format, args...)
}

func (m grpcLogger) V(l int) bool {
	switch l {
	// GRPC outputs a lot of info's, so only show them when in debug mode
	case InfoLog:
		return m.core.Enabled(zapcore.DebugLevel)
	case WarningLog:
		return m.core.Enabled(zapcore.WarnLevel)
	case ErrorLog:
		return m.core.Enabled(zapcore.ErrorLevel)
	case FatalLog:
		return m.core.Enabled(zapcore.FatalLevel)
	}

	return m.core.Enabled(zapcore.DebugLevel)
}

func NewLogger() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic("could not create development logger: " + err.Error())
	}

	return logger
}

func NewGRPCLogger(logger *zap.Logger) grpclog.LoggerV2 {
	return grpcLogger{core: logger.Core(), sugaredLogger: logger.WithOptions(zap.AddCallerSkip(callDepth)).Sugar()}
}
