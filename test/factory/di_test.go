package factory

import (
	"testing"

	"github.com/boruns/go-design-pattern/factory/di"
	"github.com/stretchr/testify/assert"
)

func TestDi(t *testing.T) {
	container := di.NewContainer()
	if err := container.Provide(NewA); err != nil {
		t.Errorf("container provide err: %+v", err)
	}

	if err := container.Provide(NewB); err != nil {
		t.Errorf("container provide err: %+v", err)
	}

	if err := container.Provide(NewC); err != nil {
		t.Errorf("container provide err: %+v", err)
	}

	if err := container.Invoke(func(a *a) {
		assert.EqualValues(t, a.b.c.Num, 1, "验证失败")
	}); err != nil {
		t.Errorf("container provide err: %+v", err)
	}
}

// A 依赖关系 A -> B -> C
type a struct {
	b *b
}

// NewA NewA
func NewA(b *b) *a {
	return &a{
		b: b,
	}
}

// B B
type b struct {
	c *c
}

// NewB NewB
func NewB(c *c) *b {
	return &b{c: c}
}

// C C
type c struct {
	Num int
}

// NewC NewC
func NewC() *c {
	return &c{
		Num: 1,
	}
}
