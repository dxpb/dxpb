package github

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gopkg.in/go-playground/webhooks.v5/github"
)

// New returns an initialized GitHub connector.
func New(c chan Commit) (Connector, error) {
	hook, err := github.New(github.Options.Secret(viper.GetString("github.secret")))
	if err != nil {
		return Connector{}, err
	}

	x := Connector{
		h:       hook,
		Commits: c,
	}

	return x, nil
}

// WebHook returns the webhook handler for the github connector.
func (c *Connector) WebHook(ctx *gin.Context) {
	payload, err := c.h.Parse(ctx.Request, github.PushEvent)
	if err != nil {
		if err == github.ErrEventNotFound {
			// ok event wasn;t one of the ones asked to be parsed
		}
	}
	switch payload.(type) {
	case github.PushPayload:
		push := payload.(github.PushPayload)
		for _, commit := range push.Commits {
			send := Commit{
				Committer: commit.Committer.Username,
				Hash:      commit.ID[:12],
				Msg:       commit.Message,
			}
			c.Commits <- send
		}
	}
}
