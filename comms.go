package main

import (
	"go.bug.st/serial"
)

func (ctx *Context) InitComms() {
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

	port, err := serial.Open(ctx.SerialPortName(), mode)
	if err != nil {
		ctx.CommsConnected = false
		ctx.Socket = nil
		ctx.Status = "Serial failure - try a different port?"
		return
	}

	port.Write([]byte("RACK\n"))

	response := make([]byte, 6)
	bytesRead := 0
	cycles := 0

	for bytesRead < 6 {
		read, err := port.Read(response)
		if err != nil {
			ctx.Status = "Serial failure - try a different port?"
			port.Close()
			return
		}

		bytesRead += read
		cycles += 1

		if cycles > 20 {
			ctx.Status = "Failure to receive ACK from serial port"
			return
		}
	}

	if string(response) == "ACKGCN" {
		ctx.Socket = &port
		ctx.CommsConnected = true
		ctx.Status = "Completed handshake with " + ctx.SerialPortName() + "!"
	} else {
		ctx.Status = "Handshake unsuccessful, socket closing"
		port.Close()
	}
}
