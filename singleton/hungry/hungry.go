// Package hungry 饿汉式单例模式
package hungry

import (
	"github.com/boruns/go-design/singleton"
)

var hungrySingleton *singleton.Singleton

// init 加载类文件时就实例化类
func init() {
	hungrySingleton = &singleton.Singleton{}
}

// GetInstance 返回实例
func GetInstance() *singleton.Singleton {
	return hungrySingleton
}
