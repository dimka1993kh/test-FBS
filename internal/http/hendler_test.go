package http_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHendler_fiboHandler(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/v1/fibo?x=1&y=3", nil)
		w := httptest.NewRecorder()
		ctx := context.Background()

		service, mockUsecase := GetService(t)
		mockUsecase.EXPECT().Fib(ctx, "1", "3").Return([]uint64{0, 1, 1}, nil)

		service.FiboHandler(w, req)

		require.Equal(t, w.Code, http.StatusOK)
		require.Equal(t, w.Body.String(), "[0,1,1]\n")
	})
}
