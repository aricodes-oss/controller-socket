package main

import (
	"fmt"

	"go.bug.st/serial"
)

func initComms(ctx *Context) {
	if ctx.CommsConnected {
		if *ctx.Socket != nil {
			(*ctx.Socket).Close()
			ctx.Socket = nil
		}

		ctx.CommsConnected = false
	}

	mode := &serial.Mode{
		BaudRate: 115200,
	}

	port, err := serial.Open(ctx.Config.SerialPortName, mode)
	if err != nil {
		ctx.CommsConnected = false
		ctx.Socket = nil
		return
	}

	port.Write([]byte("RACK\n"))

	response := make([]byte, 3)
	bytesRead := 0

	for bytesRead < 3 {
		read, err := port.Read(response)
		if err != nil {
			port.Close()
			return
		}

		bytesRead += read
	}

	if string(response) == "ACK" {
		ctx.Socket = &port
		ctx.CommsConnected = true
		fmt.Println("Completed handshake with " + ctx.Config.SerialPortName + "!")
	} else {
		fmt.Println("Handshake unsuccessful, socket closing")
		port.Close()
	}
}
