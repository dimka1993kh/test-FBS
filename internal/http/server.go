package http

import (
	"net"
	"net/http"

	"github.com/caarlos0/env/v6"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"test/internal/usecase"
)

type Server struct {
	host   string
	port   string
	router *mux.Router
}

type Config struct {
	Host string `env:"HTTP_HOST" envDefault:"0.0.0.0"`
	Port string `env:"HTTP_PORT" envDefault:"1002"`
}

func NewConfig() (*Config, error) {
	var cfg Config

	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func NewServer(cfg *Config) *Server {
	return &Server{
		host:   cfg.Host,
		port:   cfg.Port,
		router: mux.NewRouter(),
	}
}

func (s *Server) Run(usecase usecase.IUsecase) error {
	http.Handle("/", s.router)
	s.router.HandleFunc("/api/v1/fibo", NewHedler(usecase).FiboHandler)

	log.Logger.Info().Msgf("Server is listening on port %s", s.port)

	err := http.ListenAndServe(net.JoinHostPort(s.host, s.port), nil)
	if err != nil {
		return err
	}

	return nil
}
