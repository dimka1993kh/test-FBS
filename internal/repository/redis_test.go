package repository_test

import (
	"context"
	"testing"

	"test/internal/repository"

	"github.com/stretchr/testify/require"
)

func TestRedis_HGet(t *testing.T) {
	type args struct {
		ctx   context.Context
		key   string
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ok",
			args: args{
				ctx:   context.Background(),
				key:   "testKey",
				value: "testValue",
			},
			want: "testValue",
		},
	}
	for _, tt := range tests {
		cfg, err := repository.NewConfig()
		require.NoError(t, err)
		t.Run(tt.name, func(t *testing.T) {
			r := repository.NewRedis(cfg)
			err = r.HSet(tt.args.ctx, tt.args.key, tt.args.value)
			got, err := r.HGet(tt.args.ctx, tt.args.key)
			require.NoError(t, err)
			require.Equal(t, got, tt.want)
		})
	}
}
