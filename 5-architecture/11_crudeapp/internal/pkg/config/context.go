package config

import (
	"context"
)

var configKey = &struct{}{}

func WithContext(ctx context.Context, cfg *Config) context.Context {
	return context.WithValue(ctx, configKey, cfg)
}

func FromContext(ctx context.Context) *Config {
	return ctx.Value(configKey).(*Config)
}
