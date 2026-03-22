package tui

import (
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/huseynovvusal/goch/internal/discovery"
	"github.com/huseynovvusal/goch/internal/tui/shared"
)

func (m Model) viewChatting() string {
	var headerTitle string
	if len(m.onlineUsers) > 0 {
		headerTitle = "CHATTING WITH: " + m.onlineUsers[m.selectedUserIndex].Name + " | STATUS: STABLE"
	} else {
		headerTitle = "CHATTING (No peer selected)"
	}
	header := shared.HeaderStyle.Render(headerTitle)
	footer := shared.FooterStyle.Render("[ esc: Back to Hub | ^k: Clear History ]")

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
			isOut := msg.From.Name == m.username && msg.From.IP == discovery.GetSelfUser().IP

			// Handle cases where the timestamp was never initialized (e.g. dummy data / empty value)
			ts := msg.Timestamp
			if ts.IsZero() {
				ts = time.Now()
			}
			timestamp := shared.TimestampStyle.Render(ts.Format("15:04"))

			if isOut {
				bubble := shared.OutgoingBubbleStyle.Render(msg.Content)
				row := lipgloss.JoinHorizontal(lipgloss.Center, timestamp, bubble)
				rows = append(rows, lipgloss.NewStyle().Width(m.width-2).Align(lipgloss.Right).Render(row))
			} else {
				senderName := lipgloss.NewStyle().Foreground(shared.HighlightColor).Render(msg.From.Name)
				bubble := shared.IncomingBubbleStyle.Render(senderName + "\n" + msg.Content)
				row := lipgloss.JoinHorizontal(lipgloss.Center, bubble, timestamp)
				rows = append(rows, lipgloss.NewStyle().Width(m.width-2).Align(lipgloss.Left).Render(row))
			}
		}

		fullRender := strings.Join(rows, "\n")
		lines := strings.Split(fullRender, "\n")

		startIdx := len(lines) - viewportHeight - m.uiScrollOffset
		if startIdx < 0 {
			startIdx = 0
		}

		endIdx := startIdx + viewportHeight
		if endIdx > len(lines) {
			endIdx = len(lines)
		}

		var visibleLines []string
		if len(lines) > 0 && startIdx < len(lines) {
			visibleLines = lines[startIdx:endIdx]
		}

		body = lipgloss.NewStyle().Height(viewportHeight).Render(strings.Join(visibleLines, "\n"))
	}

	inputView := m.messageInput.View()
	page := lipgloss.JoinVertical(lipgloss.Left, header, "", body, inputView, footer)
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, page)
}
