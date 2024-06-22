package main

import (
	"fmt"
	"os"

	"github.com/MohamedDhiaZoghlami/go-commerce/server"
	"github.com/spf13/cobra"

	"github.com/sirupsen/logrus"
)

var commandServe *cobra.Command

var (
	argPort string
)

func init() {
	commandServe = &cobra.Command{
		Use:   "serve",
		Short: "Connect to the storage and begin serving requests.",
		Long:  ``,
		Run: func(commandServe *cobra.Command, args []string) {
			if err := serve(); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}
		},
	}
	commandServe.Flags().StringVar(&argPort, "port", ":8080", "rest address to listen on")
}

func serve() error {
	logger := logrus.New()
	s, err := server.NewServer(logger)
	if err != nil {
		return err
	}
	if err := s.Run(argPort); err != nil {
		return err
	}
	return nil
}
