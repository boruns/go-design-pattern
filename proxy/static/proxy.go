// Package static 静态代理设计模式
package static

import (
	"log"
	"time"
)

// IUser 定义需要实现的接口
type IUser interface {
	Login(username, password string) error
}

// User 用户
// @proxy IUser
type User struct{}

// Login 用户登录
func (u *User) Login(username, password string) error {
	// 不实现具体细节
	return nil
}

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
