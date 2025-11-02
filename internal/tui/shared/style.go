package shared

import "github.com/charmbracelet/lipgloss"

// Modern green-on-dark terminal aesthetic
var (
	BaseFgColor    = lipgloss.Color("#C5F5C9")   // soft light green for general text
	BaseBgColor    = lipgloss.Color("#1e2d2cff") // dark greenish background
	AccentColor    = lipgloss.Color("#00FF88")   // bright neon green accent
	HighlightColor = lipgloss.Color("#a1ffbbff") // lime highlight for selected or active
	SuccessColor   = lipgloss.Color("#2ECC71")   // emerald success tone
	ErrorColor     = lipgloss.Color("#FF5C5C")   // red for errors and alerts
	PromptColor    = lipgloss.Color("#00D4B1")   // aqua-green prompt color
	DimmedColor    = lipgloss.Color("#4E5D4E")   // muted green-gray for subtle info
	BorderColor    = lipgloss.Color("#2C3E2C")   // dark green-gray for borders
)

// Styles tuned for a CLI chat interface
var (
	HeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(AccentColor).
			Background(BaseBgColor).
			MarginBottom(1).
			MarginTop(1).
			Padding(1, 2)

	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(HighlightColor).
			MarginBottom(1)

	SubtitleStyle = lipgloss.NewStyle().
			Foreground(AccentColor).
			Bold(true).
			MarginBottom(1)

	LabelStyle = lipgloss.NewStyle().
			Foreground(SuccessColor).
			Padding(0, 1)

	InfoStyle = lipgloss.NewStyle().
			Foreground(BaseFgColor)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(ErrorColor).
			Bold(true).
			Padding(0, 1)

	PromptStyle = lipgloss.NewStyle().
			Foreground(PromptColor).
			Bold(true).
			Padding(0, 1)

	BodyStyle = lipgloss.NewStyle().
			Foreground(BaseFgColor)

	FooterStyle = lipgloss.NewStyle().
			Foreground(DimmedColor).
			MarginTop(1)

	ListTextStyle = lipgloss.NewStyle().
			Foreground(BaseFgColor).
			Padding(0, 1)

	ListSelectedTextStyle = lipgloss.NewStyle().
				Foreground(HighlightColor).
				Bold(true)

	InfoBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(BorderColor).
			Padding(1, 2).
			Margin(1, 0).
			Foreground(BaseFgColor)
)
