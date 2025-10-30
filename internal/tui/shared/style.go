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
			Padding(1, 2).
			MarginBottom(1)

	InfoStyle = lipgloss.NewStyle().
			Foreground(BaseFgColor).
			Background(BaseBgColor).
			Padding(0, 1)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(ErrorColor).
			Background(BaseBgColor).
			Bold(true).
			Padding(0, 1)

	PromptStyle = lipgloss.NewStyle().
			Foreground(HighlightColor).
			Background(BaseBgColor).
			Bold(true).
			Padding(0, 1)

	BodyStyle = lipgloss.NewStyle().
			Foreground(BaseFgColor)

	FooterStyle = lipgloss.NewStyle().
			Foreground(AccentColor).
			MarginTop(1)
)
