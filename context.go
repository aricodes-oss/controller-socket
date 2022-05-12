package main

import (
	"csock/config"
	"fmt"
	"github.com/gempir/go-twitch-irc/v3"
	"go.bug.st/serial"
)

type Context struct {
	Config         *config.Config
	SelectedPort   int32
	CommsConnected bool
	Socket         *serial.Port
	TwitchClient   *twitch.Client
}

func (c *Context) SerialPortName() string {
	ports, err := serial.GetPortsList()
	if err != nil {
		panic(err)
	}

	return ports[c.SelectedPort]
}

func (c *Context) CloseComms() {
	if c.Socket != nil {
		(*c.Socket).Close()
	}

	c.CommsConnected = false
}

func (c *Context) ConnectTwitch() {
	if c.TwitchClient != nil {
		c.TwitchClient.Disconnect()
	}

	client := twitch.NewAnonymousClient()
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(message.Message)
	})
	client.Join(c.Config.TwitchChannel)

	c.TwitchClient = client
	client.Connect()
}

func (c *Context) DisconnectTwitch() {
	if c.TwitchClient != nil {
		c.TwitchClient.Disconnect()
		c.TwitchClient = nil
	}
}
