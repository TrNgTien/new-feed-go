package main

import (
	"fmt"
	"log"

	"github.com/TrNgTien/new-feed-go/internal/wiring"
	"github.com/spf13/cobra"
)

var (
	version string
)

func server() *cobra.Command {
	command := &cobra.Command{
		Use:  "new-feed-go",
		Long: "Start standalone server for new-feed-go",
		RunE: func(cmd *cobra.Command, _ []string) error {
			app, cleanup, err := wiring.InitializeServer()

			defer cleanup()

			if err != nil {
				return err
			}

			return app.Start()
		},
	}
	return command
}

func main() {
	rootCommand := &cobra.Command{
		Version: fmt.Sprintf("Version: %s", version),
	}
	rootCommand.AddCommand(
		server(),
	)

	if err := rootCommand.Execute(); err != nil {
		log.Panic(err)
	}
}
