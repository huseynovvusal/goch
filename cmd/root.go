package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/huseynovvusal/goch/internal/chat"
	"github.com/huseynovvusal/goch/internal/db"
	"github.com/huseynovvusal/goch/internal/discovery"
	"github.com/huseynovvusal/goch/internal/tui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goch",
	Short: "goch is a CLI-based chat application working over LAN.",
	Long: `goch is a simple command-line chat application that allows users to communicate
over a local area network (LAN). It supports multiple users, private messaging, and
various customization options.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := db.Config{BroadCastPort: db.DefaultBroadcastPort, ChatPort: db.DefaultChatPort}
		store := db.NewConfigStore()
		if store.Exists() {
			if loadedCfg, err := store.Load(); err == nil {
				cfg = loadedCfg.WithDefaults()
			}
		}

		go discovery.ListenForPresence(cfg.BroadCastPort)

		chatMessages := make(chan chat.NetworkMessage)
		go chat.ListenForChatMessages(chatMessages, cfg.ChatPort)

		p := tea.NewProgram(tui.NewMainModel(chatMessages), tea.WithAltScreen())
		mainModel, err := p.Run()
		if err != nil {
			fmt.Println("Error running TUI:", err)
			os.Exit(1)
		}

		_, ok := mainModel.(tui.Model)
		if !ok {
			fmt.Println("Error asserting TUI model")
			os.Exit(1)
		}

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
