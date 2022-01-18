package decorator

import (
	"testing"

	"github.com/boruns/go-design-pattern/decorator"
	"github.com/stretchr/testify/assert"
)

func TestDecorator(t *testing.T) {
	square := decorator.Square{}
	csq := decorator.NewColorSquare(square, "red")
	got := csq.Draw()
	assert.Equal(t, got, "this is a square, color is red")
}
