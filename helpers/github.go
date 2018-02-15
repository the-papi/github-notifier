package helpers

import (
	"golang.org/x/oauth2"
	"context"
	"github.com/google/go-github/github"
	"time"
	"log"
	"github.com/gen2brain/beeep"
)

type GithubNotifier struct {
	client  *github.Client
	context context.Context
}

func NewGithubNotifier(apiToken string) (*GithubNotifier) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: apiToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return &GithubNotifier{
		client:  client,
		context: ctx,
	}
}

func (g *GithubNotifier) ListenNotifications(notificationChannel *chan *github.Notification, wakeUpInterval time.Duration) {
    for {
		opts := github.NotificationListOptions{
			Since: time.Now().Add(-wakeUpInterval),
		}

		notifications, _, err := g.client.Activity.ListNotifications(g.context, &opts)

		if err != nil {
			log.Fatal(err)
		}

		for _, v := range notifications {
			beeep.Notify("[" + *v.Subject.Type + "] " + *v.Repository.FullName, *v.Subject.Title, "./icons/octocat.png")
		}

		time.Sleep(wakeUpInterval)
	}
}
