package factory

import (
	"reflect"
	"testing"

	"github.com/boruns/go-design-pattern/factory"
	"github.com/boruns/go-design-pattern/factory/simple"
)

func TestNewIRuleConfigParser(t *testing.T) {
	type args struct {
		t string
	}

	tests := []struct {
		name string
		args args
		want factory.IRuleConfigParse
	}{
		{
			name: "json",
			args: args{t: "json"},
			want: simple.JsonRuleConfigParser{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: simple.YamlRuleConfigParser{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simple.NewIRuleConfigParser(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIRuleConfigParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
