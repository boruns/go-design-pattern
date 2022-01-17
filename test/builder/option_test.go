package builder

import (
	"testing"

	"github.com/boruns/go-design-pattern/builder"
	"github.com/boruns/go-design-pattern/builder/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResourcePoolConfigOption(t *testing.T) {
	type args struct {
		name string
		opts []option.ResourcePoolConfigOption
	}
	tests := []struct {
		name    string
		args    args
		want    *builder.ResourcePoolConfig
		wantErr bool
	}{
		{
			name: "maxTotal < maxIdle",
			args: args{
				name: "test",
				opts: []option.ResourcePoolConfigOption{
					option.SetMaxTotal(10),
					option.SetMaxIdle(20),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				name: "success",
				opts: []option.ResourcePoolConfigOption{
					option.SetMinIdle(5),
					option.SetMaxIdle(10),
					option.SetMaxTotal(20),
				},
			},
			want: &builder.ResourcePoolConfig{
				Name:     "success",
				MaxTotal: 20,
				MaxIdle:  10,
				MinIdle:  5,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := option.NewResourcePoolConfig(tt.args.name, tt.args.opts...)
			require.Equalf(t, tt.wantErr, err != nil, "error = %v, wantErr %v", err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
