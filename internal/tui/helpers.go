package tui

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/huh"
	"strconv"
)

func dbPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "goch", "goch.db")
}

func dbExists() bool {
	_, err := os.Stat(dbPath())
	return err == nil
}

func initDB(username, bio, port string) {
	os.MkdirAll(filepath.Dir(dbPath()), 0755)
	os.WriteFile(dbPath(), []byte(fmt.Sprintf("%s\n%s\n%s", username, bio, port)), 0644)
}

func loadDB(m *Model) {
	data, err := os.ReadFile(dbPath())
	if err == nil {
		lines := strings.Split(string(data), "\n")
		if len(lines) >= 3 {
			m.username = lines[0]
			m.bio = lines[1]
			m.port = lines[2]
		}
	}
}

func initForm(m *Model, theme *huh.Theme) *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Username").
				Value(&m.username).
				Validate(func(s string) error {
					if len(s) < 3 {
						return errors.New("Min 3 chars")
					}
					return nil
				}),
			huh.NewInput().
				Title("Bio").
				Value(&m.bio).
				Validate(func(s string) error {
					if len(s) > 50 {
						return errors.New("Max 50 chars")
					}
					return nil
				}),
			huh.NewInput().
				Title("Port").
				Value(&m.port).
				Validate(func(s string) error {
					if _, err := strconv.Atoi(s); err != nil {
						return errors.New("must be an integer")
					}
					return nil
				}),
		),
	).WithTheme(theme)
}
