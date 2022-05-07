package main

import (
	"csock/config"
	"go.bug.st/serial"
)

type Context struct {
	Config         *config.Config
	SelectedPort   int32
	CommsConnected bool
	Socket         *serial.Port
}
