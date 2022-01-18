package proxy

import (
	"testing"

	"github.com/boruns/go-design-pattern/proxy/static"
	"github.com/stretchr/testify/require"
)

func TestStaticProxy(t *testing.T) {
	userProxy := static.NewUserProxy(&static.User{})

	err := userProxy.Login("test", "123456")
	require.Nil(t, err)
}
