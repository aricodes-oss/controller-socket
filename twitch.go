package main

import (
	"fmt"
	"github.com/gempir/go-twitch-irc/v3"
)

func initTwitch(ctx *Context) {
	ctx.TwitchClient = twitch.NewAnonymousClient()
	fmt.Println("Connected!")
}
