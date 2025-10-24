package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goch",
	Short: "goch is a CLI-based chat application working over LAN.",
	Long: `goch is a simple command-line chat application that allows users to communicate
over a local area network (LAN). It supports multiple users, private messaging, and
various customization options.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Default action when no subcommands are provided
		fmt.Println("Welcome to goch! Use 'goch --help' to see available commands.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
