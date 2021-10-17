package http_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	us "test/internal/http"

	"github.com/stretchr/testify/require"
)

func TestErrorResponse_ToJSON(t *testing.T) {
	type fields struct {
		Code  int
		Error string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ok",
			fields: fields{
				Code:  http.StatusNotFound,
				Error: "testError",
			},
			want: "{\"code\":404,\"error\":\"testError\"}\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &us.ErrorResponse{
				Code:  tt.fields.Code,
				Error: tt.fields.Error,
			}
			w := &bytes.Buffer{}
			require.NoError(t, e.ToJSON(w))
			require.Equal(t, w.String(), tt.want)
		})
	}
}

func TestHendler_fiboHandler(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/v1/fibo?x=1&y=3", nil)
		w := httptest.NewRecorder()
		ctx := context.Background()

		service, mockUsecase := GetService(t)
		mockUsecase.EXPECT().Fib(ctx, "1", "3").Return([]uint64{0, 1, 1}, nil)

		service.FiboHandler(w, req)

		require.Equal(t, w.Code, http.StatusOK)
		require.Equal(t, w.Body.String(), "{\"code\":200,\"response\":[0,1,1]}\n")
	})
}
