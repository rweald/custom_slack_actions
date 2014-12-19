# Custom Slack Actions

A collection of custom slack actions written in Go. Provides a basic HTTP server that accepts the incoming slack custom actions and calls the correct function to respond.

This is largely an experiment with Go. 


##Usage

Custom slack actions require an API endpoint that this project provides. Running this server to respond to actions is easy:

``` go run main.go ```


####Adding new actions
Adding actions is easy. Simply implement the ```SlackAction``` interface that can be found in [actions/base.go](https://github.com/rweald/custom_slack_actions/blob/master/actions/base.go)


##ToDo

- [ ] Add Makefile for compiling & installing binary. 
- [ ] Add testing for main server and existing actions
- [ ] Add some useful actions

