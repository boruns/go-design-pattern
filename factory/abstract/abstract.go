// Package abstract 抽象工厂模式
package abstract

type IRuleConfigParser interface {
	Parser(data []byte)
}

type JsonRuleConfigParser struct{}

func (j JsonRuleConfigParser) Parser(data []byte) {
	panic("implement me")
}

type ISysConfigParser interface {
	ParseSystem(data []byte)
}

type JsonSystemConfigParser struct{}

func (j JsonSystemConfigParser) ParseSystem(data []byte) {
	panic("implement me")
}

type IConfigParser interface {
	CreateRuleParser() IRuleConfigParser
	CreateSystemParser() ISysConfigParser
}

type JsonConfigParserFactory struct{}

func (j JsonConfigParserFactory) CreateRuleParser() IRuleConfigParser {
	return JsonRuleConfigParser{}
}

func (j JsonConfigParserFactory) CreateSystemParser() ISysConfigParser {
	return JsonSystemConfigParser{}
}
