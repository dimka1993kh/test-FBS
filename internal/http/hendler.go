package http

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"

	"test/internal/usecase"
)

type Hendler struct {
	usecase usecase.IUsecase
}

func NewHedler(usecase usecase.IUsecase) *Hendler {
	return &Hendler{usecase: usecase}
}

func (h *Hendler) FiboHandler(w http.ResponseWriter, r *http.Request) {
	var (
		response      Response
		errorResponse ErrorResponse
	)

	x, ok := r.URL.Query()["x"]
	if !ok {
		log.Logger.Error().Msg("Error: не введен Х")

		errorResponse.Code = http.StatusBadRequest
		errorResponse.Error = "не введен Х"

		err := errorResponse.ToJSON(w)
		if err != nil {
			log.Logger.Error().Err(err)
		}

		return
	}

	y, ok := r.URL.Query()["y"]
	if !ok {
		log.Logger.Error().Msg("Error: не введен Y")

		errorResponse.Code = http.StatusBadRequest
		errorResponse.Error = "не введен Y"

		err := errorResponse.ToJSON(w)
		if err != nil {
			log.Logger.Error().Err(err)
		}

		return
	}

	resp, err := h.usecase.Fib(context.Background(), x[0], y[0])
	if err != nil {
		log.Logger.Error().Msgf("fibo error: %s", err)

		errorResponse.Code = http.StatusBadRequest
		errorResponse.Error = err.Error()

		err = errorResponse.ToJSON(w)
		if err != nil {
			log.Logger.Error().Err(err)
		}

		return
	}

	w.Header().Set("Content-Type", "application/json")

	response.Response = resp
	response.Code = http.StatusOK

	err = response.ToJSON(w)
	if err != nil {
		log.Logger.Err(err)
	}
}

type Response struct {
	Code     int      `json:"code"`
	Response []uint64 `json:"response"`
}

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func (e *ErrorResponse) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(e)
}

func (r *Response) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(r)
}
