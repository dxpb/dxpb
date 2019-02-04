package main

import (
	"log"

	"github.com/dxpb/dxpb/internal/github"
	"github.com/dxpb/dxpb/internal/http"
	"github.com/dxpb/dxpb/internal/irc"
)

func main() {
	log.Println("hello!")
	Start()

	ircClient, err := irc.New()
	if err != nil {
		log.Panic(err)
	}

	web, err := http.New()
	if err != nil {
		log.Panic(err)
	}

	gitEvent := make(chan github.Commit)
	gh, err := github.New(gitEvent)
	if err != nil {
		log.Panic(err)
	}
	web.RegisterGinHook("/ghwebhook", "POST", gh.WebHook)

	log.Println("Connecting to IRC")
	go ircClient.Connect()

	go web.Serve()
	for {
		select {
		case event := <-gitEvent:
			ircClient.NoticeCommit(event.Committer, event.Hash, event.Msg)
		}
	}
}
