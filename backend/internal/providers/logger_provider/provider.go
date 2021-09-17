package logger_provider

import "go.uber.org/zap"

func NewLogger(level string) (*zap.Logger, error) {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	atom := zap.NewAtomicLevel()
	_ = atom.UnmarshalText([]byte(level))

	loggerConfig.Level = atom

	return loggerConfig.Build(zap.Hooks())
}
