package grpc_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	fibo "test/proto"
)

func TestFiboService_GetFibo(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		in := &fibo.FiboRequest{
			X: "4",
			Y: "8",
		}
		ctx := context.Background()
		want := &fibo.FigoResponse{
			Response: []uint64{3, 5, 8},
		}

		service, mockUsecase := GetService(t)
		mockUsecase.EXPECT().Fib(ctx, "4", "8").Return([]uint64{3, 5, 8}, nil)

		res, err := service.GetFibo(ctx, in)
		require.NoError(t, err)
		require.Equal(t, res, want)
	})
}
