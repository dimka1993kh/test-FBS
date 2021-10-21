package http

import (
	"context"
	"net/http"

	"github.com/rs/zerolog/log"

	"test/internal/usecase"
)

type Hendler struct {
	usecase usecase.IUsecase
}

func NewHandler(usecase usecase.IUsecase) *Hendler {
	return &Hendler{usecase: usecase}
}

func (h *Hendler) FiboHandler(w http.ResponseWriter, r *http.Request) {
	x, ok := r.URL.Query()["x"]
	if !ok {
		log.Logger.Error().Msg("error: не введен Х")
		ResponseJSON(w, "error: не введен Х", http.StatusBadRequest)

		return
	}

	y, ok := r.URL.Query()["y"]
	if !ok {
		log.Logger.Error().Msg("error: не введен Y")
		ResponseJSON(w, "error: не введен Y", http.StatusBadRequest)

		return
	}

	resp, err := h.usecase.Fib(context.Background(), x[0], y[0])
	if err != nil {
		log.Logger.Error().Msgf("fibo error: %s", err)
		ResponseJSON(w, err.Error(), http.StatusBadRequest)

		return
	}

	ResponseJSON(w, resp, http.StatusOK)
}
