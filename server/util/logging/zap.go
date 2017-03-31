package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapConfig(env string) zap.Config {
	if env == "production" {
		return zap.NewProductionConfig()
	}
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return config
}
