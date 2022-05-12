package main

import (
	"fmt"
	"github.com/gempir/go-twitch-irc/v3"
	"time"
)

func (ctx *Context) Release(b Button) {
	if !ctx.ButtonState.Test(uint(b)) {
		return
	}

	ctx.ButtonState.Clear(uint(b))
	(*ctx.Socket).Write([]byte(fmt.Sprintf("r%d\n", b)))
}

func (ctx *Context) Press(b Button) {
	if ctx.ButtonState.Test(uint(b)) {
		return
	}

	ctx.ButtonState.Set(uint(b))
	(*ctx.Socket).Write([]byte(fmt.Sprintf("p%d\n", b)))
	go func() {
		time.Sleep(time.Millisecond * time.Duration(ctx.Config.PressDuration))
		ctx.Release(b)
	}()
}

func messageCallback(ctx *Context, message twitch.PrivateMessage) {
	triggers := ctx.Config.ButtonTriggers
	switch message.Message {
	case triggers.A:
		ctx.Press(A)
	case triggers.B:
		ctx.Press(B)
	case triggers.X:
		ctx.Press(X)
	case triggers.Y:
		ctx.Press(Y)

	case triggers.L:
		ctx.Press(L)
	case triggers.R:
		ctx.Press(R)
	case triggers.Z:
		ctx.Press(Z)

	case triggers.START:
		ctx.Press(START)

	case triggers.UP:
		ctx.Press(UP)
	case triggers.DOWN:
		ctx.Press(DOWN)
	case triggers.LEFT:
		ctx.Press(LEFT)
	case triggers.RIGHT:
		ctx.Press(RIGHT)
	}
}
