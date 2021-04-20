package loggerfx

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Provide(New)

func New() (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return logger, nil
}
