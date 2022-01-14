package singleton

import (
	"testing"

	"github.com/boruns/go-design/singleton/hungry"
	"github.com/stretchr/testify/assert"
)

func TestGetHungryInstance(t *testing.T) {
	assert.Equal(t, hungry.GetInstance(), hungry.GetInstance())
}

func BenchmarkGetHungryInstanceParallel(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			if hungry.GetInstance() != hungry.GetInstance() {
				b.Error("test fail")
			}
		}
	})
}
