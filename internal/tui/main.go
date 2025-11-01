package tui

import (
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/huseynovvusal/goch/internal/config"
	"github.com/huseynovvusal/goch/internal/discovery"
	"github.com/huseynovvusal/goch/internal/tui/shared"
)

type Model struct {
	nameInput         textinput.Model
	name              string
	nameSubmitted     bool
	onlineUsers       []discovery.NetworkUser
	selectedUserIndex int
}

type UpdateUsersMsg []discovery.NetworkUser

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
	if m.nameSubmitted {
		header := shared.HeaderStyle.Render("Hello, " + m.name + "!")
		footer := shared.FooterStyle.Render("Press q or ctrl+c to quit.")

		var body string
		if len(m.onlineUsers) > 0 {
			users := shared.SubtitleStyle.Render("Online users:") + "\n"

			for _, user := range m.onlineUsers {
				prefix := " "
				style := shared.ListTextStyle

				if m.selectedUserIndex < len(m.onlineUsers) && user == m.onlineUsers[m.selectedUserIndex] {
					prefix = "> "
					style = shared.ListSelectedTextStyle
				}

				users += style.Render(prefix+user.Name+" ("+user.IP+")") + "\n"
			}

			body = shared.BodyStyle.Render(users)
		} else {
			body = shared.BodyStyle.Render("No online users found.")
		}

		return header + "\n\n" + body + "\n\n" + footer
	}

	return shared.HeaderStyle.Render("Goch - LAN Chat Application") +
		"\n\n" +
		shared.BodyStyle.Render(m.nameInput.View()) +
		"\n\n" +
		shared.FooterStyle.Render("(Press Enter to submit)")
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.nameSubmitted && len(m.onlineUsers) > 0 {
				if m.selectedUserIndex > 0 {
					m.selectedUserIndex--
				}
			}
			return m, nil
		case "down", "j":
			if m.nameSubmitted && len(m.onlineUsers) > 0 {
				if m.selectedUserIndex < len(m.onlineUsers)-1 {
					m.selectedUserIndex++
				}
			}
			return m, nil
		case "enter":
			if !m.nameSubmitted {
				name := m.nameInput.Value()

				m.name = name
				m.nameSubmitted = true

				go discovery.BroadcastPresence(m.name, 8787)

				return m, tea.Tick(time.Duration(config.ONLINE_USERS_REFRESH_INTERVAL)*time.Second, func(t time.Time) tea.Msg {
					users := discovery.GetOnlineUsers()

					return UpdateUsersMsg(users)
				})
			}

		}
	case UpdateUsersMsg:
		m.onlineUsers = msg.(UpdateUsersMsg)

		return m, tea.Tick(time.Duration(config.ONLINE_USERS_REFRESH_INTERVAL)*time.Second, func(t time.Time) tea.Msg {
			return UpdateUsersMsg(discovery.GetOnlineUsers())
		})
	}

	if !m.nameSubmitted {
		var cmd tea.Cmd
		m.nameInput, cmd = m.nameInput.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m Model) GetName() string {
	return m.name
}
