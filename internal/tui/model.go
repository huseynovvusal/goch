package tui

import (
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/huseynovvusal/goch/internal/chat"
	"github.com/huseynovvusal/goch/internal/config"
	"github.com/huseynovvusal/goch/internal/db"
	"github.com/huseynovvusal/goch/internal/discovery"
)

type state int

const (
	stateOnboarding state = iota
	stateHub
	stateChatting
	stateSettings
)

type Model struct {
	state state

	// Onboarding
	form     *huh.Form
	username string
	bio      string
	port     string

	// Hub
	onlineUsers       []discovery.NetworkUser
	selectedUserIndex int

	// Screen dimensions
	width  int
	height int

	// Chat
	messageInput     textinput.Model
	chatMessages     []chat.NetworkMessage
	chatMessagesChan chan chat.NetworkMessage
	msgStore         *db.MessageStore
	chatOffset       int
	uiScrollOffset   int
}

type UpdateUsersMsg []discovery.NetworkUser

func NewMainModel(chatMessagesChan chan chat.NetworkMessage) Model {
	messageInput := textinput.New()
	messageInput.Placeholder = "Enter your message..."
	messageInput.CharLimit = 256
	messageInput.Width = 50

	msgStore, _ := db.NewMessageStore()

	m := Model{
		messageInput:     messageInput,
		chatMessages:     []chat.NetworkMessage{},
		chatMessagesChan: chatMessagesChan,
		msgStore:         msgStore,
		chatOffset:       0,
		uiScrollOffset:   0,
		username:         "",
		bio:              "",
		port:             "7070",
	}

	store := db.NewConfigStore()
	if store.Exists() {
		m.state = stateHub
		if cfg, err := store.Load(); err == nil {
			m.username = cfg.Username
			m.bio = cfg.Bio
			m.port = cfg.Port
		}
		// addDummyData(&m)
	} else {
		m.state = stateOnboarding
		theme := huh.ThemeDracula()
		m.form = initForm(&m, theme)
	}

	return m
}

func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, textinput.Blink)

	if m.state == stateOnboarding {
		cmds = append(cmds, m.form.Init())
	} else if m.state == stateHub {
		go discovery.BroadcastPresence(m.username, config.BROADCAST_PORT)
		cmds = append(cmds, tea.Tick(time.Duration(config.ONLINE_USERS_REFRESH_INTERVAL)*time.Second, func(t time.Time) tea.Msg {
			users := discovery.GetOnlineUsers()
			return UpdateUsersMsg(users)
		}))
	}

	return tea.Batch(cmds...)
}

func (m Model) GetName() string {
	return m.username
}
