package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/huseynovvusal/goch/internal/tui/shared"
)

type Model struct{}

func NewMainModel() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) View() string {
	return shared.HeaderStyle.Render("Goch - LAN Chat Application")
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}
