package cmd

import (
	"fmt"
	"os"

	"github.com/huseynovvusal/goch/internal/discovery"
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

		var name string
		fmt.Print("Enter your name: ")
		fmt.Scanln(&name)
		fmt.Println("Broadcasting your presence on the network...")

		go discovery.ListenForPresence(8787)

		discovery.BroadcastPresence(name, 8787)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
