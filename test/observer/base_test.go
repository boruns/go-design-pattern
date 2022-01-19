package observer

import (
	"testing"

	"github.com/boruns/go-design-pattern/observer/base"
)

func TestBaseObserver(t *testing.T) {
	sub := &base.Subject{}
	sub.Register(&base.Observer1{})
	sub.Register(&base.Observer2{})
	sub.Notify("hi")
}
