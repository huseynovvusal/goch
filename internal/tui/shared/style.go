package shared

import "github.com/charmbracelet/lipgloss"

var (
	BaseFgColor    = lipgloss.Color("#FFFFFF")
	BaseBgColor    = lipgloss.Color("#1E1E1E")
	AccentColor    = lipgloss.Color("#FF6AC1")
	BorderColor    = lipgloss.Color("#44475A")
	HighlightColor = lipgloss.Color("#50FA7B")
	ErrorColor     = lipgloss.Color("#FF5555")
)

var (
	HeaderStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(AccentColor).
		Background(BaseBgColor).
		Padding(1, 2)
)
