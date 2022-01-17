package option

import (
	"fmt"

	"github.com/boruns/go-design-pattern/builder"
)

type ResourcePoolConfigOption func(config *builder.ResourcePoolConfig)

func SetMaxTotal(total int) ResourcePoolConfigOption {
	return func(config *builder.ResourcePoolConfig) {
		config.MaxTotal = total
	}
}

func SetMaxIdle(maxIdle int) ResourcePoolConfigOption {
	return func(config *builder.ResourcePoolConfig) {
		config.MaxIdle = maxIdle
	}
}

func SetMinIdle(minIdle int) ResourcePoolConfigOption {
	return func(config *builder.ResourcePoolConfig) {
		config.MinIdle = minIdle
	}
}

func NewResourcePoolConfig(name string, options ...ResourcePoolConfigOption) (*builder.ResourcePoolConfig, error) {
	config := &builder.ResourcePoolConfig{
		Name: name,
	}
	for _, f := range options {
		f(config)
	}
	// 设置默认值
	if config.MinIdle == 0 {
		config.MinIdle = builder.DefaultMinIdle
	}

	if config.MaxIdle == 0 {
		config.MaxIdle = builder.DefaultMaxIdle
	}

	if config.MaxTotal == 0 {
		config.MaxTotal = builder.DefaultMaxTotal
	}

	if config.MaxTotal < config.MaxIdle {
		return nil, fmt.Errorf("max total(%d) cannot < max idle(%d)", config.MaxTotal, config.MaxIdle)
	}

	if config.MinIdle > config.MaxIdle {
		return nil, fmt.Errorf("max idle(%d) cannot < min idle(%d)", config.MaxIdle, config.MinIdle)
	}
	return config, nil
}
