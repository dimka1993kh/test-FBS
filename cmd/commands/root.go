package commands

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "test",
		Short: "test server",
	}

	initServerCmd(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err)
	}
}
