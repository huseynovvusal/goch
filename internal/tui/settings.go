package tui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/huseynovvusal/goch/internal/tui/shared"
)

func (m Model) viewSettings() string {
	logo := lipgloss.NewStyle().Foreground(shared.PrimaryColor).Bold(true).Render(shared.GoatLogo)
	tagline := lipgloss.NewStyle().Foreground(shared.DimmedColor).Render("SETTINGS")
	formView := m.form.View()
	footer := lipgloss.NewStyle().Foreground(shared.DimmedColor).Render("Press Enter to save changes, Esc to cancel...")

	page := lipgloss.JoinVertical(lipgloss.Center, logo, tagline, "\n", formView, "\n", footer)
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, page)
}
