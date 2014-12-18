package actions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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

type SlackResponse struct {
	Text string
}

var commands = make(map[string]*SlackAction)

func RegisterHandler(pattern string, sa SlackAction) {
	commands[pattern] = &sa
}

func handleInboundSlackAction(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	cs := r.Form["command"]

	if len(cs) < 1 {
		http.Error(w, "Must specify a command", http.StatusNotFound)
		return
	}

	requested_action := cs[0]

	if requested_action == "" {
		http.Error(w, "Must specify a command", http.StatusNotFound)
		return
	}

	var sr *SlackResponse = nil

	for k, v := range commands {
		if k == requested_action {
			err, s := (*v).HandleAction(&SlackMessage{})

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			sr = s
			break
		}
	}

	resp, err := json.Marshal(sr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func RunSlackActionServer() {
	fmt.Println("Server Starting")
	http.HandleFunc("/", handleInboundSlackAction)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
