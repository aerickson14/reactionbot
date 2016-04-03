package main

import (
	"os"

	slackbot "github.com/BeepBoopHQ/go-slackbot"
	"github.com/nlopes/slack"
	"golang.org/x/net/context"
)

const (
	WithTyping    = slackbot.WithTyping
	WithoutTyping = slackbot.WithoutTyping
)

type ReactionBot slackbot.Bot

func (b *ReactionBot) New(slackToken string) *ReactionBot {
	b := &ReactionBot{Client: slack.New(slackToken)}
	return b
}

func main() {
	bot := ReactionBot.New(os.Getenv("SLACK_TOKEN"))
	
	bot.Hear(".*").MessageHandler(CatchAllHandler)
	bot.Run()
}

func (b *ReactionBot) ReplyWithReaction(evt *slack.MessageEvent, reaction string, typing bool) {
	item := NewMessageItem(evt.Msg.Channel, evt.Msg)
	b.Client.AddReaction(reaction, item)
}

func CatchAllHandler(ctx context.Context, bot *ReactionBot, evt *slack.MessageEvent) {
	bot.ReplyWithReaction(evt, "+:+1:", WithoutTyping)
}
