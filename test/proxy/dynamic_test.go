package proxy

import (
	"testing"

	"github.com/boruns/go-design-pattern/proxy/dynamic"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDynamicProxy(t *testing.T) {
	want := `package static

type UserProxy struct {
	child *User
}

func NewUserProxy(child *User) *UserProxy {
	return &UserProxy{child: child}
}

func (p *UserProxy) Login(username, password string) (r0 error) {
	// before 这里可能会有一些统计的逻辑
	start := time.Now()

	r0 = p.child.Login(username, password)

	// after 这里可能也有一些监控统计的逻辑
	log.Printf("user login cost time: %s", time.Since(start))

	return r0
}
`
	got, err := dynamic.Generate("../../proxy/static/proxy.go")
	require.Nil(t, err)
	assert.Equal(t, want, got)
}
