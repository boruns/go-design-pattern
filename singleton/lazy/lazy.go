// Package lazy 懒汉式单例模式（双重检测）
package lazy

import (
	"sync"

	"github.com/boruns/go-design-pattern/singleton"
)

var (
	// lazySingleton 懒汉式实例
	lazySingleton *singleton.Singleton

	// 锁
	once = &sync.Once{}
)

// GetInstance 获取懒汉式单例
func GetInstance() *singleton.Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &singleton.Singleton{}
		})
	}
	return lazySingleton
}
