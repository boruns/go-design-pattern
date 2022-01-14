// Package simple 简单工厂模式
package simple

import "github.com/boruns/go-design-pattern/factory"

// jsonRuleConfigParse json 格式解析
type JsonRuleConfigParser struct{}

// Parse 实现 Parse 接口
func (j JsonRuleConfigParser) Parser(data []byte) {
	panic("implement me")
}

// yamlRuleConfigParse yaml 格式解析
type YamlRuleConfigParser struct{}

func (y YamlRuleConfigParser) Parser(data []byte) {
	panic("implement me")
}

// NewIRuleConfigParse 工厂类入口
func NewIRuleConfigParser(t string) factory.IRuleConfigParse {
	switch t {
	case "json":
		return JsonRuleConfigParser{}
	case "yaml":
		return YamlRuleConfigParser{}
	}
	return nil
}
