package usecase_test

import (
	"context"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/require"
)

func TestService_Fib(t *testing.T) {
	t.Run("wrong X", func(t *testing.T) {
		service, _ := GetService(t)
		ctx := context.Background()
		x := ""
		y := "4"

		got, err := service.Fib(ctx, x, y)
		require.Equal(t, err.Error(), "error: Х должно быть целым числом")
		require.Nil(t, got)
	})

	t.Run("wrong Y", func(t *testing.T) {
		service, _ := GetService(t)
		ctx := context.Background()
		x := "1"
		y := ""

		got, err := service.Fib(ctx, x, y)
		require.Equal(t, err.Error(), "error: Y должно быть целым числом")
		require.Nil(t, got)
	})

	t.Run("wrong serial number X", func(t *testing.T) {
		service, _ := GetService(t)
		ctx := context.Background()
		x := "-1"
		y := "4"

		got, err := service.Fib(ctx, x, y)
		require.Equal(t, err.Error(), "error: порядковый номер X должен быть больше или равен 0")
		require.Nil(t, got)
	})

	t.Run("wrong serial number Y", func(t *testing.T) {
		service, _ := GetService(t)
		ctx := context.Background()
		x := "0"
		y := "-1"

		got, err := service.Fib(ctx, x, y)
		require.Equal(t, err.Error(), "error: порядковый номер Y должен быть больше или равен 0")
		require.Nil(t, got)
	})

	t.Run("ok (with redis)", func(t *testing.T) {
		service, mockRepo := GetService(t)
		ctx := context.Background()
		x := "0"
		y := "4"
		want := []uint64{0, 1, 1, 2, 3}

		mockRepo.EXPECT().HGet(ctx, "2").Return("1", nil)
		mockRepo.EXPECT().HGet(ctx, "3").Return("2", nil)
		mockRepo.EXPECT().HGet(ctx, "4").Return("3", nil)

		got, err := service.Fib(ctx, x, y)
		require.NoError(t, err)
		require.Equal(t, got, want)
	})

	t.Run("ok (without redis)", func(t *testing.T) {
		service, mockRepo := GetService(t)
		ctx := context.Background()
		x := "0"
		y := "4"
		want := []uint64{0, 1, 1, 2, 3}

		mockRepo.EXPECT().HGet(ctx, "2").Return("", redis.Nil)
		mockRepo.EXPECT().HSet(ctx, "2", uint64(1)).Return(nil)
		mockRepo.EXPECT().HGet(ctx, "3").Return("", redis.Nil)
		mockRepo.EXPECT().HSet(ctx, "3", uint64(2)).Return(nil)
		mockRepo.EXPECT().HGet(ctx, "4").Return("", redis.Nil)
		mockRepo.EXPECT().HSet(ctx, "4", uint64(3)).Return(nil)

		got, err := service.Fib(ctx, x, y)
		require.NoError(t, err)
		require.Equal(t, got, want)
	})
}
