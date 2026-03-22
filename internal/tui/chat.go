package tui

import (
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/huseynovvusal/goch/internal/tui/shared"
)

func (m Model) viewChatting() string {
	var headerTitle string
	if len(m.onlineUsers) > 0 {
		headerTitle = "CHATTING WITH: " + m.onlineUsers[m.selectedUserIndex].Name + " | STATUS: STABLE \U0001f4f6"
	} else {
		headerTitle = "CHATTING (No peer selected)"
	}
	header := shared.HeaderStyle.Render(headerTitle)
	footer := shared.FooterStyle.Render("[ esc: Back to Hub | ^k: Clear History | ^l: Export Log ]")

	var body string
	headerHeight := lipgloss.Height(header)
	footerHeight := lipgloss.Height(footer)
	inputHeight := lipgloss.Height(m.messageInput.View())

	viewportHeight := m.height - headerHeight - footerHeight - inputHeight - 2
	if viewportHeight < 5 {
		viewportHeight = 5
	}

	if len(m.chatMessages) == 0 {
		body = lipgloss.NewStyle().Height(viewportHeight).Render(shared.BodyStyle.Render("No messages yet.") + "\n")
	} else {
		var rows []string
		for _, msg := range m.chatMessages {
			isOut := msg.From.Name == m.username
			timestamp := shared.TimestampStyle.Render(time.Now().Format("15:04"))

			if isOut {
				bubble := shared.OutgoingBubbleStyle.Render(msg.Content)
				row := lipgloss.JoinHorizontal(lipgloss.Center, timestamp, bubble)
				rows = append(rows, lipgloss.NewStyle().Width(m.width - 2).Align(lipgloss.Right).Render(row))
			} else {
				bubble := shared.IncomingBubbleStyle.Render(msg.From.Name + "\n" + msg.Content)
				row := lipgloss.JoinHorizontal(lipgloss.Center, bubble, timestamp)
				rows = append(rows, lipgloss.NewStyle().Width(m.width - 2).Align(lipgloss.Left).Render(row))
			}
		}

		startIdx := 0
		if len(rows) > viewportHeight {
			startIdx = len(rows) - viewportHeight
		}
		body = lipgloss.NewStyle().Height(viewportHeight).Render(strings.Join(rows[startIdx:], "\n"))
	}

	inputView := m.messageInput.View()
	page := lipgloss.JoinVertical(lipgloss.Left, header, "", body, inputView, footer)
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, page)
}
