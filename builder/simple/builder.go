package simple

import (
	"fmt"

	"github.com/boruns/go-design-pattern/builder"
)

type ResourcePoolConfigBuilder struct {
	Name     string
	MaxTotal int
	MaxIdle  int
	MinIdle  int
}

// SetName SetName
func (b *ResourcePoolConfigBuilder) SetName(name string) error {
	if name == "" {
		return fmt.Errorf("name can not be empty")
	}
	b.Name = name
	return nil
}

// SetMinIdle SetMinIdle
func (b *ResourcePoolConfigBuilder) SetMinIdle(minIdle int) error {
	if minIdle < 0 {
		return fmt.Errorf("max tatal cannot < 0, input: %d", minIdle)
	}
	b.MinIdle = minIdle
	return nil
}

// SetMaxIdle SetMaxIdle
func (b *ResourcePoolConfigBuilder) SetMaxIdle(maxIdle int) error {
	if maxIdle < 0 {
		return fmt.Errorf("max tatal cannot < 0, input: %d", maxIdle)
	}
	b.MaxIdle = maxIdle
	return nil
}

// SetMaxTotal SetMaxTotal
func (b *ResourcePoolConfigBuilder) SetMaxTotal(maxTotal int) error {
	if maxTotal <= 0 {
		return fmt.Errorf("max tatal cannot <= 0, input: %d", maxTotal)
	}
	b.MaxTotal = maxTotal
	return nil
}

// Build Build
func (b *ResourcePoolConfigBuilder) Build() (*builder.ResourcePoolConfig, error) {
	if b.Name == "" {
		return nil, fmt.Errorf("name can not be empty")
	}

	// 设置默认值
	if b.MinIdle == 0 {
		b.MinIdle = builder.DefaultMinIdle
	}

	if b.MaxIdle == 0 {
		b.MaxIdle = builder.DefaultMaxIdle
	}

	if b.MaxTotal == 0 {
		b.MaxTotal = builder.DefaultMaxTotal
	}

	if b.MaxTotal < b.MaxIdle {
		return nil, fmt.Errorf("max total(%d) cannot < max idle(%d)", b.MaxTotal, b.MaxIdle)
	}

	if b.MinIdle > b.MaxIdle {
		return nil, fmt.Errorf("max idle(%d) cannot < min idle(%d)", b.MaxIdle, b.MinIdle)
	}

	return &builder.ResourcePoolConfig{
		Name:     b.Name,
		MaxTotal: b.MaxTotal,
		MaxIdle:  b.MaxIdle,
		MinIdle:  b.MinIdle,
	}, nil
}
