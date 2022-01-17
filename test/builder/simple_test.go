package builder

import (
	"testing"

	"github.com/boruns/go-design-pattern/builder"
	"github.com/boruns/go-design-pattern/builder/simple"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResourcePoolConfigBuilder(t *testing.T) {
	tests := []struct {
		name    string
		builder *simple.ResourcePoolConfigBuilder
		want    *builder.ResourcePoolConfig
		wantErr bool
	}{
		{
			name: "name empty",
			builder: &simple.ResourcePoolConfigBuilder{
				Name:     "",
				MaxTotal: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "maxIdle < minIdle",
			builder: &simple.ResourcePoolConfigBuilder{
				Name:     "test",
				MaxIdle:  1,
				MinIdle:  10,
				MaxTotal: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			builder: &simple.ResourcePoolConfigBuilder{
				Name: "test",
			},
			want: &builder.ResourcePoolConfig{
				Name:     "test",
				MaxTotal: builder.DefaultMaxTotal,
				MaxIdle:  builder.DefaultMaxIdle,
				MinIdle:  builder.DefaultMinIdle,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.builder.Build()
			require.Equalf(t, tt.wantErr, err != nil, "Build() error = %v, wantErr %v", err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
