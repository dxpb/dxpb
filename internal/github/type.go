package github

import (
	"gopkg.in/go-playground/webhooks.v5/github"
)

// Connector implements all the functionality needed to respond to
// GitHub interactions.
type Connector struct {
	h       *github.Webhook
	Commits chan Commit
}

// Commit inclues information about a particular commit returned from
// the GitHub API.
type Commit struct {
	Committer string
	Hash      string
	Msg       string
}
