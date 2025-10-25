package cmd

import "github.com/spf13/cobra"

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Manage users in the chat application",
	Long:  `The users command allows you to add, remove, and list users in the chat application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Default action when no subcommands are provided
	},
}

func init() {
	rootCmd.AddCommand(usersCmd)
}
