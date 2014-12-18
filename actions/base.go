package actions

// SlackAction defines the interface that all other
// actions must abide by. Very basic right now as it is
// the responsibility of the action itself to parse any
// complex parts of the SlackMessage
type SlackAction interface {
	HandleAction(m *SlackMessage) (error, *SlackResponse)
}

func init() {
	RegisterHandler("hello", HelloAction{})
}
