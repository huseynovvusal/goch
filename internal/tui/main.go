package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/huseynovvusal/goch/internal/tui/shared"
)

type Model struct {
	nameInput textinput.Model
	name      string
	submitted bool
}

func NewMainModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Enter your name"
	ti.Focus()
	ti.CharLimit = 32
	ti.Width = 20

	return Model{
		nameInput: ti,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) View() string {
	if m.submitted {
		return shared.HeaderStyle.Render("Hello, " + m.name + "!\nPress q or ctrl+c to quit.")
	}
	return shared.HeaderStyle.Render("Goch - LAN Chat Application\n\n" + m.
		nameInput.View() + "\n(Press Enter to submit)")
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			if !m.submitted {
				m.name = m.nameInput.Value()
				m.submitted = true
				return m, nil
			}
		}
	}

	if !m.submitted {
		var cmd tea.Cmd
		m.nameInput, cmd = m.nameInput.Update(msg)
		return m, cmd
	}

	return m, nil
}
