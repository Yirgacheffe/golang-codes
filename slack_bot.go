package main

import (
	"log"

	"github.com/slack-go/slack"
)

const (
	OAUTH_TOKEN = "<Token From Slack>"
	CHANNEL_ID  = "<Channel>"
)

func main() {

	api := slack.New(OAUTH_TOKEN)
	attachment := slack.Attachment{
		Pretext: "Pretext",
		Text:    "Hello from Golang~~!",
	}

	channelId, timestamp, err := api.PostMessage(
		CHANNEL_ID,
		slack.MsgOptionText("This is the only message", false),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		log.Fatalf("%s\n", err)
	}

	log.Printf("Message successfully sent to Channel %s at %s\n", channelId, timestamp)
}
