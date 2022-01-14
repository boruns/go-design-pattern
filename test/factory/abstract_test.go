package factory

import (
	"reflect"
	"testing"

	"github.com/boruns/go-design-pattern/factory/abstract"
)

func TestCreateRuleParser(t *testing.T) {
	tests := []struct {
		name string
		want abstract.IRuleConfigParser
	}{
		{
			name: "json",
			want: abstract.JsonRuleConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := abstract.JsonConfigParserFactory{}
			if got := j.CreateRuleParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateRuleParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateSystemParser(t *testing.T) {
	tests := []struct {
		name string
		want abstract.ISysConfigParser
	}{
		{
			name: "json",
			want: abstract.JsonSystemConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := abstract.JsonConfigParserFactory{}
			if got := j.CreateSystemParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSystemParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
