package actions

// PingAction is a very basic action that listens for the /ping
// command and then responds with "pong" to indicate that the server
// is still alive
type PingAction struct{}

func (a PingAction) HandleAction(m *SlackMessage) (error, *SlackResponse) {
	return nil, &SlackResponse{"Pong"}
}

func init() {
	RegisterHandler("ping", PingAction{})
}
