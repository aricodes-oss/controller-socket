package main

import (
	"csock/config"
	"github.com/bits-and-blooms/bitset"
	"github.com/gempir/go-twitch-irc/v3"
	"go.bug.st/serial"
)

type Context struct {
	Config         *config.Config
	SelectedPort   int32
	CommsConnected bool
	Socket         *serial.Port
	TwitchClient   *twitch.Client
	ButtonState    bitset.BitSet
	Status         string
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

	if c.Config.TwitchChannel == "" {
		c.Status = "Please enter a channel name before connecting."
		return
	}

	client := twitch.NewAnonymousClient()
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		messageCallback(c, message)
	})
	client.Join(c.Config.TwitchChannel)

	c.Status = "Connected to " + c.Config.TwitchChannel

	c.TwitchClient = client
	client.Connect()
}

func (c *Context) DisconnectTwitch() {
	if c.TwitchClient != nil {
		c.TwitchClient.Disconnect()
		c.TwitchClient = nil
		c.Status = "Twitch client disconnected"
	}
}
