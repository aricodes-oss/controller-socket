package main

import (
	"fmt"
	"github.com/gempir/go-twitch-irc/v3"
	"time"
)

func (ctx *Context) Release(b Button) {
	if !ctx.Pressed(b) {
		return
	}

	ctx.ButtonState.Clear(uint(b))
	(*ctx.Socket).Write([]byte(fmt.Sprintf("r%d\n", b)))
}

func (ctx *Context) Press(b Button) {
	if ctx.Pressed(b) {
		return
	}

	ctx.ButtonState.Set(uint(b))
	(*ctx.Socket).Write([]byte(fmt.Sprintf("p%d\n", b)))
}

func (ctx *Context) PressAndRelease(b Button) {
	go func() {
		ctx.Press(b)
		time.Sleep(time.Millisecond * time.Duration(ctx.Config.PressDuration))
		ctx.Release(b)
	}()
}

func (ctx *Context) Toggle(b Button) {
	if ctx.Pressed(b) {
		ctx.Release(b)
	} else {
		ctx.Press(b)
	}
}

func (ctx *Context) Pressed(b Button) bool {
	return ctx.ButtonState.Test(uint(b))
}

func HandleDirection(ctx *Context, current, opposite Button) {
	socdAllowed := ctx.Config.AllowSOCD
	holdDirections := ctx.Config.HoldDirections

	// No sanity checks applied
	if socdAllowed {
		if holdDirections {
			// SOCD && Hold
			ctx.Toggle(current)
		} else {
			// SOCD && !Hold
			ctx.PressAndRelease(current)
		}
	} else {
		// !SOCD
		if ctx.Pressed(opposite) {
			ctx.Release(opposite)
		}

		if holdDirections {
			ctx.Toggle(current)
		} else {
			ctx.PressAndRelease(current)
		}
	}
}

func messageCallback(ctx *Context, message twitch.PrivateMessage) {
	triggers := ctx.Config.ButtonTriggers
	switch message.Message {
	case triggers.A:
		ctx.PressAndRelease(A)
	case triggers.B:
		ctx.PressAndRelease(B)
	case triggers.X:
		ctx.PressAndRelease(X)
	case triggers.Y:
		ctx.PressAndRelease(Y)

	case triggers.L:
		ctx.PressAndRelease(L)
	case triggers.R:
		ctx.PressAndRelease(R)
	case triggers.Z:
		ctx.PressAndRelease(Z)

	case triggers.START:
		ctx.PressAndRelease(START)

	case triggers.UP:
		HandleDirection(ctx, UP, DOWN)
	case triggers.DOWN:
		HandleDirection(ctx, DOWN, UP)
	case triggers.LEFT:
		HandleDirection(ctx, LEFT, RIGHT)
	case triggers.RIGHT:
		HandleDirection(ctx, RIGHT, LEFT)

	case triggers.CUP:
		HandleDirection(ctx, UP, DOWN)
	case triggers.CDOWN:
		HandleDirection(ctx, DOWN, UP)
	case triggers.CLEFT:
		HandleDirection(ctx, LEFT, RIGHT)
	case triggers.CRIGHT:
		HandleDirection(ctx, RIGHT, LEFT)
	}
}
