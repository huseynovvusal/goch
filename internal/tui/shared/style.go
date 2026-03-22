package shared

import "github.com/charmbracelet/lipgloss"

// Modern green-on-dark terminal aesthetic
var (
	PrimaryColor   = lipgloss.Color("#D27D2D") // Earthy Orange/Brown
	SecondaryColor = lipgloss.Color("#DEB887") // Burlywood / Light Brown
	AccentColor    = lipgloss.Color("#FF8C00") // Vibrant Orange
	BaseFgColor    = lipgloss.Color("#FDF5E6") // Old Lace / Off-white
	BaseBgColor    = lipgloss.Color("#1C140D") // Very Dark Brown
	HighlightColor = lipgloss.Color("#FFD700") // Golden highlight
	SuccessColor   = lipgloss.Color("#8FBC8F") // Earthy Green
	ErrorColor     = lipgloss.Color("#CD5C5C") // Earthy Red
	PromptColor    = lipgloss.Color("#FFB90F") // Dark Goldenrod
	DimmedColor    = lipgloss.Color("#8B5A2B") // Muted Brown
	BorderColor    = lipgloss.Color("#4A3020") // Dark warm gray-brown border
)

// Styles tuned for a CLI chat interface
var (
	HeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(BaseBgColor).
			Background(PrimaryColor).
			MarginBottom(1).
			MarginTop(1).
			Padding(0, 1)

	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(HighlightColor).
			MarginBottom(1)

	SubtitleStyle = lipgloss.NewStyle().
			Foreground(SecondaryColor).
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

	MessageFromStyle = lipgloss.NewStyle().
				Foreground(SecondaryColor).
				Bold(true)

	MessageContentStyle = lipgloss.NewStyle().
				Foreground(BaseFgColor).
				PaddingLeft(1)

	PanelStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(BorderColor).
			Padding(0, 1)

	LeftColumnStyle = PanelStyle.Copy()

	RightColumnStyle = PanelStyle.Copy()

	BottomSectionStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(BorderColor).
				Padding(0, 1)

	OutgoingBubbleStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(PrimaryColor).
				Padding(0, 1).
				MarginBottom(1)

	IncomingBubbleStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(SecondaryColor).
				Padding(0, 1).
				MarginBottom(1)

	TimestampStyle = lipgloss.NewStyle().
			Foreground(DimmedColor).
			MarginLeft(1).
			MarginRight(1)
)
