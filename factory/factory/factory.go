// Package factory 工厂模式
package factory

import (
	fac "github.com/boruns/go-design-pattern/factory"
	"github.com/boruns/go-design-pattern/factory/simple"
)

// IRuleConfigParseFactory 工厂抽象类
type IRuleConfigParseFactory interface {
	CreateParser() fac.IRuleConfigParse
}

// JsonRuleConfigParserFactory json 工厂类
type JsonRuleConfigParserFactory struct{}

// CreateParser 实现抽象类的 CreateParser
func (j JsonRuleConfigParserFactory) CreateParser() fac.IRuleConfigParse {
	return simple.JsonRuleConfigParser{}
}

// YamlRuleConfigParserFactory yaml 工厂类
type YamlRuleConfigParserFactory struct{}

// CreateParser 实现抽象类的 CreateParser
func (y YamlRuleConfigParserFactory) CreateParser() fac.IRuleConfigParse {
	return simple.YamlRuleConfigParser{}
}

// NewIRuleConfigParseFactory 获取工厂类方法
func NewIRuleConfigParseFactory(t string) IRuleConfigParseFactory {
	switch t {
	case "json":
		return JsonRuleConfigParserFactory{}
	case "yaml":
		return YamlRuleConfigParserFactory{}
	}
	return nil
}
