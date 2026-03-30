package tui

import (
	"context"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/huseynovvusal/goch/internal/chat"
	"github.com/huseynovvusal/goch/internal/config"
	"github.com/huseynovvusal/goch/internal/db"
	"github.com/huseynovvusal/goch/internal/discovery"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	select {
	case chatMsg := <-m.chatMessagesChan:
		m.chatMessages = append(m.chatMessages, chatMsg)
		if m.msgStore != nil {
			_ = m.msgStore.SaveMessage(context.Background(), chatMsg, chatMsg.From.IP)
		}
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

	if m.state == stateOnboarding || m.state == stateSettings {
		form, cmd := m.form.Update(msg)
		if f, ok := form.(*huh.Form); ok {
			m.form = f
		}
		if m.form.State == huh.StateCompleted {
			pastState := m.state
			m.state = stateHub

			m.username = m.form.GetString("username")
			m.bio = m.form.GetString("bio")
			m.port = m.form.GetString("port")

			store := db.NewConfigStore()
			_ = store.Save(db.Config{
				Username: m.username,
				Bio:      m.bio,
				Port:     m.port,
			})

			if pastState == stateOnboarding {
				go discovery.BroadcastPresence(m.username, config.BROADCAST_PORT)

				return m, tea.Batch(cmd, tea.Tick(time.Duration(config.ONLINE_USERS_REFRESH_INTERVAL)*time.Second, func(t time.Time) tea.Msg {
					users := discovery.GetOnlineUsers()
					return UpdateUsersMsg(users)
				}))
			}
			return m, cmd
		}
		return m, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.state == stateHub && len(m.onlineUsers) > 0 && m.selectedUserIndex > 0 {
				m.selectedUserIndex--
			} else if m.state == stateChatting {
				if len(m.chatMessages) > 0 {
					m.uiScrollOffset++

					// If we are at the top of the currently fetched block, fetch more!
					if m.uiScrollOffset+m.height/2 >= len(m.chatMessages) && m.msgStore != nil {
						peer := m.onlineUsers[m.selectedUserIndex]
						m.chatOffset += 50
						olderMsgs, err := m.msgStore.GetMessages(context.Background(), peer.IP, 50, m.chatOffset)
						if err == nil && len(olderMsgs) > 0 {
							m.chatMessages = append(olderMsgs, m.chatMessages...)
						} else {
							m.chatOffset -= 50
						}
					}
				}
			}
			return m, nil
		case "down":
			if m.state == stateHub && len(m.onlineUsers) > 0 && m.selectedUserIndex < len(m.onlineUsers)-1 {
				m.selectedUserIndex++
			} else if m.state == stateChatting {
				if m.uiScrollOffset > 0 {
					m.uiScrollOffset--
				}
			}
			return m, nil
		case "ctrl+q":
			if m.state == stateHub {
				return m, tea.Quit
			}
		case "s":
			if m.state == stateHub {
				m.state = stateSettings
				m.form = initForm(&m, huh.ThemeDracula())
				return m, m.form.Init()
			}
		case "ctrl+k":
			if m.state == stateChatting {
				m.chatMessages = make([]chat.NetworkMessage, 0)
			}
		case "esc":
			if m.state == stateChatting || m.state == stateSettings {
				m.state = stateHub
			}
		case "enter":
			if m.state == stateHub {
				if len(m.onlineUsers) > 0 {
					m.state = stateChatting
					m.messageInput.Focus()

					m.chatOffset = 0
					m.uiScrollOffset = 0
					if m.msgStore != nil {
						peer := m.onlineUsers[m.selectedUserIndex]
						msgs, err := m.msgStore.GetMessages(context.Background(), peer.IP, 50, m.chatOffset)
						if err == nil {
							m.chatMessages = msgs
						}
					}
				}
			} else if m.state == stateChatting {
				messageContent := strings.TrimSpace(m.messageInput.Value())

				if messageContent != "" {
					toUser := m.onlineUsers[m.selectedUserIndex]
					fromUser := discovery.GetSelfUser()
					_ = chat.SendChatMessage(messageContent, toUser, fromUser)

					message := chat.NetworkMessage{
						Content:   messageContent,
						From:      discovery.GetSelfUser(),
						Timestamp: time.Now(),
					}

					m.chatMessages = append(m.chatMessages, message)
					if m.msgStore != nil {
						_ = m.msgStore.SaveMessage(context.Background(), message, toUser.IP)
					}

					m.uiScrollOffset = 0
					m.messageInput.SetValue("")
				}
			}
		}
	case UpdateUsersMsg:
		m.onlineUsers = msg
		// addDummyData(&m)

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
