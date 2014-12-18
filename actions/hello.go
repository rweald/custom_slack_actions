package actions

import "errors"

// HelloAction is a first take on an action that actually uses the text
// provided by the slack hook.
type HelloAction struct{}

func (h HelloAction) HandleAction(m *SlackMessage) (error, *SlackResponse) {
	if m.Text == "" {
		return errors.New("Don't know who to say hello to. Please specify a name"), nil
	}
	return nil, &SlackResponse{"Hello " + m.Text}
}
