package tui

func (m Model) View() string {
	if m.state == stateOnboarding {
		return m.viewOnboarding()
	}

	if m.state == stateHub {
		return m.viewHub()
	}

	if m.state == stateChatting {
		return m.viewChatting()
	}

	return ""
}
