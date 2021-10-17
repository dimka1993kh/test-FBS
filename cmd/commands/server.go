package commands

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"test/internal/grpc"
	"test/internal/http"
	"test/internal/repository"
	"test/internal/usecase"
)

const (
	RunGRPCServer = "run_grpc_server"
	RunHTTPServer = "run_http_server"
)

func initServerCmd(root *cobra.Command) {
	root.AddCommand(
		&cobra.Command{
			Use:   RunGRPCServer,
			Short: "Run gRPC server",
			RunE:  run,
		},
		&cobra.Command{
			Use:   RunHTTPServer,
			Short: "Run http server",
			RunE:  run,
		},
	)
}

func run(command *cobra.Command, _ []string) error {
	repCfg, err := repository.NewConfig()
	if err != nil {
		log.Logger.Fatal().Err(err)
	}

	rep := repository.NewRedis(repCfg)
	use := usecase.New(&usecase.Config{Repository: rep})

	grpcCfg, err := grpc.NewConfig()
	if err != nil {
		log.Logger.Fatal().Err(err)
	}

	grpcServer := grpc.NewServer(grpcCfg)

	httpCfg, err := http.NewConfig()
	if err != nil {
		log.Logger.Fatal().Err(err)
	}

	httpServer := http.NewServer(httpCfg)

	switch command.Use {
	case RunGRPCServer:
		grpcServer.Run(use)
	case RunHTTPServer:
		err := httpServer.Run(use)
		if err != nil {
			log.Logger.Fatal().Msgf("http server error: %s", err)
		}
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
	sign := <-quit
	log.Logger.Info().Msgf("Shutting down server... Reason: %s", sign.String())

	return nil
}
