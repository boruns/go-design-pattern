package singleton

import (
	"testing"

	"github.com/boruns/go-design-pattern/singleton/lazy"
	"github.com/stretchr/testify/assert"
)

func TestGetLazyInstance(t *testing.T) {
	assert.Equal(t, lazy.GetInstance(), lazy.GetInstance())
}

func BenchmarkGetLazyInstance(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			if lazy.GetInstance() != lazy.GetInstance() {
				b.Error("test fail")
			}
		}
	})
}
