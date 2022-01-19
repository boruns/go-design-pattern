package observer

import (
	"fmt"
	"testing"
	"time"

	"github.com/boruns/go-design-pattern/observer/eventbus"
)

func sub1(msg1, msg2 string) {
	time.Sleep(1 * time.Microsecond)
	fmt.Printf("sub1, %s %s\n", msg1, msg2)
}

func sub2(msg1, msg2 string) {
	fmt.Printf("sub2, %s %s\n", msg1, msg2)
}

func sub3(msg1, msg2 string) {
	fmt.Printf("sub3 %s %s\n", msg1, msg2)
}

func TestEventBusObserver(t *testing.T) {
	eventBus := eventbus.NewAsyncEventBus()
	eventBus.Subscribe("topic:1", sub1)
	eventBus.Subscribe("topic:1", sub3)
	eventBus.Subscribe("topic:2", sub2)
	eventBus.Publish("topic:1", "testA", "testB")
	eventBus.Publish("topic:2", "testA", "testB")
	time.Sleep(1 * time.Second)
}
