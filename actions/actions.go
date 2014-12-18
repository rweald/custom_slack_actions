package actions

type SlackAction interface {
	HandleAction(m *SlackMessage) (error, *SlackResponse)
}

func init() {
	RegisterHandler("hello", HelloAction{})
}

type HelloAction struct{}

func (h HelloAction) HandleAction(m *SlackMessage) (error, *SlackResponse) {
	return nil, &SlackResponse{"Hello World"}
}
