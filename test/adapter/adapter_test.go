package adapter

import (
	"testing"

	"github.com/boruns/go-design-pattern/adapter"
	"github.com/stretchr/testify/assert"
)

func TestAliyunAdapter(t *testing.T) {
	ICreateServer := &adapter.AliyunClientAdapter{}
	err := ICreateServer.CreateServer(1, 1)
	assert.Nil(t, err)
}

func TestAwsAdapter(t *testing.T) {
	ICreateServer := &adapter.AwsClientAdapter{}
	err := ICreateServer.CreateServer(2, 2)
	assert.Nil(t, err)
}
