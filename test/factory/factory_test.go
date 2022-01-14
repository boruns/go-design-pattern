package factory

import (
	"reflect"
	"testing"

	"github.com/boruns/go-design-pattern/factory/factory"
)

func TestNewIRuleConfigParseFactory(t *testing.T) {
	type args struct {
		t string
	}

	tests := []struct {
		name string
		args args
		want factory.IRuleConfigParseFactory
	}{
		{
			name: "json",
			args: args{t: "json"},
			want: factory.JsonRuleConfigParserFactory{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: factory.YamlRuleConfigParserFactory{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factory.NewIRuleConfigParseFactory(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIRuleConfigParserFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}
