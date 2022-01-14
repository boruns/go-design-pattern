// Package di  依赖注入设计模式
package di

import (
	"errors"
	"fmt"
	"reflect"
)

// Container di 容器
type Container struct {
	// providers 假设一种类型只能有一个 provider 提供
	providers map[reflect.Type]prodiver

	results map[reflect.Type]reflect.Value
}

type prodiver struct {
	value reflect.Value

	params []reflect.Type
}

// NewContainer 实例化容器
func NewContainer() *Container {
	return &Container{
		providers: map[reflect.Type]prodiver{},
		results:   map[reflect.Type]reflect.Value{},
	}
}

// IsError 判断是否是 error 类型
func IsError(t reflect.Type) bool {
	if t.Kind() != reflect.Interface {
		return false
	}
	return t.Implements(reflect.TypeOf(reflect.TypeOf((*error)(nil)).Elem()))
}

// Provide 对象提供者，需要传入一个对象的工厂方法，后续会用于对象的创建
func (c *Container) Provide(constructor interface{}) error {
	v := reflect.ValueOf(constructor)

	// 仅支持函数 provider
	if v.Kind() != reflect.Func {
		return errors.New("constructor must be func")
	}

	vt := v.Type()

	// 获取参数
	params := make([]reflect.Type, vt.NumIn())
	for i := 0; i < vt.NumIn(); i++ {
		params[i] = vt.In(i)
	}

	// 获取返回值
	results := make([]reflect.Type, vt.NumOut())
	for i := 0; i < vt.NumOut(); i++ {
		results[i] = vt.Out(i)
	}

	prodiver := prodiver{
		value:  v,
		params: params,
	}

	for _, result := range results {
		if IsError(result) {
			continue
		}
		if _, ok := c.providers[result]; ok {
			return fmt.Errorf("%s had a provider", result)
		}
		c.providers[result] = prodiver
	}
	return nil
}

// Invoke 函数执行入口
func (c *Container) Invoke(function interface{}) error {
	v := reflect.ValueOf(function)

	// 仅支持函数 provider
	if v.Kind() != reflect.Func {
		return fmt.Errorf("constructor must be a func")
	}

	vt := v.Type()

	// 获取参数
	var err error
	params := make([]reflect.Value, vt.NumIn())
	for i := 0; i < vt.NumIn(); i++ {
		params[i], err = c.buildParam(vt.In(i))
		if err != nil {
			return err
		}
	}

	v.Call(params)

	// 获取 providers
	return nil
}

// buildParam 构建参数
// 1. 从容器中获取 provider
// 2. 递归获取 provider 的参数值
// 3. 获取到参数之后执行函数
// 4. 将结果缓存并且返回结果
func (c *Container) buildParam(param reflect.Type) (val reflect.Value, err error) {
	if result, ok := c.results[param]; ok {
		return result, nil
	}

	provider, ok := c.providers[param]
	if !ok {
		return reflect.Value{}, fmt.Errorf("can not found provider: %s", param)
	}

	params := make([]reflect.Value, len(provider.params))
	for i, p := range provider.params {
		params[i], err = c.buildParam(p)
		if err != nil {
			return reflect.Value{}, err
		}
	}

	results := provider.value.Call(params)
	for _, result := range results {
		// 判断是否报错
		if IsError(result.Type()) && !result.IsNil() {
			return reflect.Value{}, fmt.Errorf("%+v call err: %+v", provider, result)
		}

		if !IsError(result.Type()) && !result.IsNil() {
			c.results[result.Type()] = result
		}
	}
	return c.results[param], nil
}
