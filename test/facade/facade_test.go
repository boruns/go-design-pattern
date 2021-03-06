package facade

import (
	"testing"

	"github.com/boruns/go-design-pattern/facade"
	"github.com/stretchr/testify/assert"
)

func TestUserService_Login(t *testing.T) {
	service := facade.UserService{}
	user, err := service.Login(13001010101, 1234)
	assert.NoError(t, err)
	assert.Equal(t, &facade.User{Name: "test login"}, user)
}

func TestUserService_LoginOrRegister(t *testing.T) {
	service := facade.UserService{}
	user, err := service.LoginOrRegister(13001010101, 1234)
	assert.NoError(t, err)
	assert.Equal(t, &facade.User{Name: "test login"}, user)
}

func TestUserService_Register(t *testing.T) {
	service := facade.UserService{}
	user, err := service.Register(13001010101, 1234)
	assert.NoError(t, err)
	assert.Equal(t, &facade.User{Name: "test register"}, user)
}
