package tui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/huseynovvusal/goch/internal/tui/shared"
)

func (m Model) viewHub() string {
	// --- STATUS BAR ---
	statusBar := shared.FooterStyle.Render("[ ^q: Quit | s: Settings | Enter: Start Chatting ]")

	// Calculate remaining height for columns
	availHeight := m.height - lipgloss.Height(statusBar)
	if availHeight < 5 {
		availHeight = 5
	}
	colHeight := availHeight - 2
	if colHeight < 2 {
		colHeight = 2
	}

	// --- LEFT COLUMN: Network Nodes ---
	leftWidth := int(float64(m.width) * 0.3)
	if leftWidth < 20 {
		leftWidth = 20
	}

	leftContent := shared.SubtitleStyle.Render("NETWORK NODES") + "\n"
	if len(m.onlineUsers) > 0 {
		for _, user := range m.onlineUsers {
			indicator := lipgloss.NewStyle().Foreground(shared.SuccessColor).Render("●")
			nameStr := fmt.Sprintf("%s (%s)", user.Name, user.IP)
			if m.onlineUsers[m.selectedUserIndex] == user {
				nameStr = lipgloss.NewStyle().Foreground(shared.HighlightColor).Bold(true).Render("> " + nameStr)
			} else {
				nameStr = "  " + nameStr
			}
			latency := lipgloss.NewStyle().Foreground(shared.DimmedColor).Render("[2ms]")
			leftContent += fmt.Sprintf("%s %s %s\n", indicator, nameStr, latency)
		}
	} else {
		leftContent += lipgloss.NewStyle().Foreground(shared.DimmedColor).Render("Scanning for peers...")
	}
	leftCol := shared.LeftColumnStyle.Width(leftWidth - 4).Height(colHeight).Render(leftContent)

	// --- RIGHT COLUMN: Selected Peer Details ---
	rightWidth := m.width - leftWidth
	if rightWidth < 20 {
		rightWidth = 20
	}

	rightContent := shared.SubtitleStyle.Render("SELECTED PEER DETAILS") + "\n\n"
	if len(m.onlineUsers) > 0 {
		selected := m.onlineUsers[m.selectedUserIndex]

		header := lipgloss.NewStyle().Foreground(shared.PrimaryColor).Bold(true).Render(selected.Name)
		bio := lipgloss.NewStyle().Foreground(shared.DimmedColor).Render("Bio: No bio available")
		ip := shared.InfoStyle.Render("IP: " + selected.IP)
		lastSeen := shared.InfoStyle.Render("Last Seen: Just now")

		rightContent += lipgloss.JoinVertical(lipgloss.Left, header, bio, "\n", ip, "Port: 8989", lastSeen)
	} else {
		rightContent += lipgloss.NewStyle().Foreground(shared.DimmedColor).Render("No peer selected.")
	}
	rightCol := shared.RightColumnStyle.Width(rightWidth - 4).Height(colHeight).Render(rightContent)

	columns := lipgloss.JoinHorizontal(lipgloss.Top, leftCol, rightCol)

	page := lipgloss.JoinVertical(lipgloss.Left, columns, statusBar)
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, page)
}
