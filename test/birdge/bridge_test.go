package brigde

import (
	"testing"

	"github.com/boruns/go-design-pattern/bridge"
	"github.com/stretchr/testify/assert"
)

func TestBridge(t *testing.T) {
	sender := bridge.NewEmailMsgSender([]string{"test@test.test"})
	notice := bridge.NewErrorNotification(sender)
	err := notice.Notify("test")
	assert.Nil(t, err)
}
