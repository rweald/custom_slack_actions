package actions

import "fmt"

type SlackMessage struct {
	Token       string
	TeamId      string
	ChannelId   string
	ChannelName string
	UserId      string
	Username    string
	Command     string
	Text        string
}

type SlackCommand struct{}

//var commands = make()

func (sc SlackCommand) HandleAction(m *SlackMessage) (error, string) {
	return nil, ""
}

func Run() {
	fmt.Println("Server Starting")
}
