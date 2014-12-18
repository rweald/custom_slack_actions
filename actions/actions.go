package actions

import "errors"

type SlackAction interface {
	HandleAction(m *SlackMessage) (error, *SlackResponse)
}

func init() {
	RegisterHandler("hello", HelloAction{})
}

type HelloAction struct{}

func (h HelloAction) HandleAction(m *SlackMessage) (error, *SlackResponse) {
	if m.Text == "" {
		return errors.New("Don't know who to say hello to. Please specify a name"), nil
	}
	return nil, &SlackResponse{"Hello " + m.Text}
}

type PingAction struct{}

func (a PingAction) HandleAction(m *SlackMessage) (error, *SlackResponse) {
	return nil, &SlackResponse{"Pong"}
}

func init() {
	RegisterHandler("ping", PingAction{})
}
