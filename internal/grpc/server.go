package grpc

import (
	"net"

	"test/internal/usecase"
	fibo "test/proto"

	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type Server struct {
	host string
	port string
}

func NewServer(cfg *Config) *Server {
	return &Server{
		host: cfg.Host,
		port: cfg.Port,
	}
}

type Config struct {
	Host string `env:"GRPS_HOST" envDefault:"0.0.0.0"`
	Port string `env:"GRPC_PORT" envDefault:"1001"`
}

func NewConfig() (*Config, error) {
	var cfg Config

	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (s *Server) Run(usecase usecase.IUsecase) {
	server := grpc.NewServer()

	instance := NewFiboService(usecase)

	fibo.RegisterFiboServiceServer(server, instance)

	listener, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		log.Logger.Fatal().Msgf("Unable to create grpc listener: %s", err)
	}

	log.Logger.Info().Msgf("Server is listening on port %s", s.port)

	if err = server.Serve(listener); err != nil {
		log.Logger.Fatal().Msgf("Unable to start server: %s", err)
	}
}
