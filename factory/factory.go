package factory

// IRuleConfigParse 抽象接口类
type IRuleConfigParse interface {
	Parser(data []byte)
}
