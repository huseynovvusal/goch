package tui

import (
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/huseynovvusal/goch/internal/config"
	"github.com/huseynovvusal/goch/internal/discovery"
	"github.com/huseynovvusal/goch/internal/tui/shared"
)

type state int

const (
	stateEnterName state = iota
	stateShowUsers
	stateChat
)

type Model struct {
	state state

	nameInput     textinput.Model
	name          string
	nameSubmitted bool

	onlineUsers       []discovery.NetworkUser
	selectedUserIndex int

	messageInput textinput.Model
	messages     []Message
}

type Message struct {
	Content string
	From    discovery.NetworkUser
}

type UpdateUsersMsg []discovery.NetworkUser

func NewMainModel() Model {
	nameInput := textinput.New()
	nameInput.Placeholder = "Enter your name"
	nameInput.Focus()
	nameInput.CharLimit = 32
	nameInput.Width = 20

	messageInput := textinput.New()
	messageInput.Placeholder = "Type your message"
	messageInput.CharLimit = 256
	messageInput.Width = 50

	return Model{
		nameInput:    nameInput,
		messageInput: messageInput,
		state:        stateEnterName,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) View() string {
	if m.state == stateEnterName {
		return shared.HeaderStyle.Render("Goch - LAN Chat Application") +
			"\n\n" +
			shared.BodyStyle.Render(m.nameInput.View()) +
			"\n\n" +
			shared.FooterStyle.Render("(Press Enter to submit)")
	}

	self := discovery.GetSelfUser()
	selfInfo := ""
	if self.Name != "" && self.IP != "" {
		selfInfo = shared.InfoStyle.Render("You: " + self.Name + " (" + self.IP + ")")
	}

	header := shared.HeaderStyle.Render("Hello, " + m.name + "!" + "\n" + selfInfo)
	footer := shared.FooterStyle.Render("Press ctrl+c to quit.")

	var body string
	if len(m.onlineUsers) > 0 {
		users := shared.SubtitleStyle.Render("Online users:") + "\n"

		for _, user := range m.onlineUsers {
			prefix := " "
			style := shared.ListTextStyle

			if m.state == stateShowUsers && m.onlineUsers[m.selectedUserIndex] == user {
				prefix = "> "
				style = shared.ListSelectedTextStyle
			} else if m.state == stateChat && m.onlineUsers[m.selectedUserIndex] == user {
				prefix = "* "
				style = shared.ListSelectedTextStyle
			}

			users += style.Render(prefix+user.Name+" ("+user.IP+")") + "\n"
		}

		body = shared.BodyStyle.Render(users)
	} else {
		body = shared.BodyStyle.Render("No online users found.")
	}

	if m.state == stateChat {
		body += "\n" + shared.SubtitleStyle.Render("Chatting with "+m.onlineUsers[m.selectedUserIndex].Name+":") + "\n"
		if len(m.messages) == 0 {
			body += shared.BodyStyle.Render("No messages yet.") + "\n"
		} else {
			for _, msg := range m.messages {
				body += shared.MessageFromStyle.Render(msg.From.Name+": ") + shared.MessageContentStyle.Render(msg.Content) + "\n"
			}
		}

		body += "\n" + m.messageInput.View()
	}

	return header + "\n\n" + body + "\n\n" + footer

}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c":
			return m, tea.Quit
		case "up":
			if m.state == stateShowUsers {
				if len(m.onlineUsers) > 0 {
					if m.selectedUserIndex > 0 {
						m.selectedUserIndex--
					}
				}
			}
			return m, nil
		case "down":
			if m.state == stateShowUsers {
				if len(m.onlineUsers) > 0 {
					if m.selectedUserIndex < len(m.onlineUsers)-1 {
						m.selectedUserIndex++
					}
				}
			}
			return m, nil
		case "enter":

			switch m.state {
			case stateEnterName:
				if !m.nameSubmitted {
					name := m.nameInput.Value()

					m.name = name
					m.nameSubmitted = true
					m.state = stateShowUsers

					go discovery.BroadcastPresence(m.name, 8787)

					return m, tea.Tick(time.Duration(config.ONLINE_USERS_REFRESH_INTERVAL)*time.Second, func(t time.Time) tea.Msg {
						users := discovery.GetOnlineUsers()

						return UpdateUsersMsg(users)
					})
				}

			case stateShowUsers:
				if len(m.onlineUsers) > 0 {
					m.state = stateChat
					m.messageInput.Focus()
				}

			case stateChat:
				messageContent := m.messageInput.Value()
				if messageContent != "" {
					message := Message{
						Content: messageContent,
						From:    discovery.GetSelfUser(),
					}
					m.messages = append(m.messages, message)
					m.messageInput.SetValue("")

				}
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

	if m.state == stateChat {
		var cmd tea.Cmd
		m.messageInput, cmd = m.messageInput.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m Model) GetName() string {
	return m.name
}
