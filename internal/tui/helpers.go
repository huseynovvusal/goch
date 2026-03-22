package tui

import (
	"errors"
	"strconv"

	"github.com/charmbracelet/huh"
)

func initForm(m *Model, theme *huh.Theme) *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Key("username").
				Title("Username").
				Validate(func(s string) error {
					if len(s) < 3 {
						return errors.New("Min 3 chars")
					}
					return nil
				}),
			huh.NewInput().
				Key("bio").
				Title("Bio").
				Validate(func(s string) error {
					if len(s) > 50 {
						return errors.New("Max 50 chars")
					}
					return nil
				}),
			huh.NewInput().
				Key("port").
				Title("Port").
				Validate(func(s string) error {
					if _, err := strconv.Atoi(s); err != nil {
						return errors.New("must be an integer")
					}
					return nil
				}),
		),
	).WithTheme(theme)
}
