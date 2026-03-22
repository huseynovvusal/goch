package tui

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/huseynovvusal/goch/internal/chat"
	"github.com/huseynovvusal/goch/internal/config"
	"github.com/huseynovvusal/goch/internal/discovery"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	select {
	case chatMsg := <-m.chatMessagesChan:
		m.chatMessages = append(m.chatMessages, chatMsg)
	default:
	}

	// Handle global keys
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	if m.state == stateOnboarding {
		form, cmd := m.form.Update(msg)
		if f, ok := form.(*huh.Form); ok {
			m.form = f
		}
		if m.form.State == huh.StateCompleted {
			m.state = stateHub
			initDB(m.username, m.bio, m.port)

			go discovery.BroadcastPresence(m.username, config.BROADCAST_PORT)

			return m, tea.Batch(cmd, tea.Tick(time.Duration(config.ONLINE_USERS_REFRESH_INTERVAL)*time.Second, func(t time.Time) tea.Msg {
				users := discovery.GetOnlineUsers()
				return UpdateUsersMsg(users)
			}))
		}
		return m, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.state == stateHub && len(m.onlineUsers) > 0 && m.selectedUserIndex > 0 {
				m.selectedUserIndex--
			}
			return m, nil
		case "down":
			if m.state == stateHub && len(m.onlineUsers) > 0 && m.selectedUserIndex < len(m.onlineUsers)-1 {
				m.selectedUserIndex++
			}
			return m, nil
		case "q":
			if m.state == stateHub {
				return m, tea.Quit
			}
		case "s":
			if m.state == stateHub {
				// Settings mocked or ignored for now
			}
		case "ctrl+k":
			if m.state == stateChatting {
				m.chatMessages = make([]chat.NetworkMessage, 0)
			}
		case "esc":
			if m.state == stateChatting {
				m.state = stateHub
			}
		case "enter":
			if m.state == stateHub {
				if len(m.onlineUsers) > 0 {
					m.state = stateChatting
					m.messageInput.Focus()
				}
			} else if m.state == stateChatting {
				messageContent := strings.TrimSpace(m.messageInput.Value())

				if messageContent != "" {
					toUser := m.onlineUsers[m.selectedUserIndex]
					fromUser := discovery.GetSelfUser()
					chat.SendChatMessage(messageContent, toUser, fromUser)

					message := chat.NetworkMessage{
						Content: messageContent,
						From:    discovery.GetSelfUser(),
					}

					m.chatMessages = append(m.chatMessages, message)
					m.messageInput.SetValue("")
				}
			}
		}
	case UpdateUsersMsg:
		m.onlineUsers = msg
		addDummyData(&m)

		return m, tea.Tick(time.Duration(config.ONLINE_USERS_REFRESH_INTERVAL)*time.Second, func(t time.Time) tea.Msg {
			return UpdateUsersMsg(discovery.GetOnlineUsers())
		})
	}

	if m.state == stateChatting {
		var cmd tea.Cmd
		m.messageInput, cmd = m.messageInput.Update(msg)
		return m, cmd
	}

	return m, nil
}
