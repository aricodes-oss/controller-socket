package main

import (
	"fmt"

	g "github.com/AllenDang/giu"
	"go.bug.st/serial"
)

const INPUT_SIZE = 128
const V_SPACING = 16.0

func availableSerialPorts() []string {
	ports, err := serial.GetPortsList()
	if err != nil {
		panic(err)
	}

	return ports
}

func onPortSelect(ctx *Context) {
	if ctx.CommsConnected {
		fmt.Println("Closing stale serial connection...")

		if ctx.Socket != nil {
			(*ctx.Socket).Close()
		}

		ctx.CommsConnected = false
	}
}

func commsButton(ctx *Context) *g.ButtonWidget {
	if ctx.CommsConnected {
		return g.Button("Disconnect").OnClick(ctx.CloseComms)
	}

	return g.Button("Connect").OnClick(ctx.InitComms)
}

func loop(ctx *Context) {
	ports := availableSerialPorts()
	g.PushWindowPadding(16.0, 16.0)
	g.PushFramePadding(8.0, 8.0)
	defer g.PopStyleV(2)

	g.SingleWindow().Layout(
		g.Row(
			g.Column(
				g.Label("Twitch Channel"),
				g.InputText(&ctx.Config.TwitchChannel).Size(g.Auto).OnChange(ctx.Config.Save),
			),
		),
		g.Dummy(0.0, V_SPACING),
		g.Column(
			g.Label("Serial Port"),
			g.Combo("", ports[ctx.SelectedPort], ports, &ctx.SelectedPort).OnChange(func() {
				onPortSelect(ctx)
			}),
		),

		g.Dummy(8.0, V_SPACING),
		g.Column(
			g.Label("Button Press Duration (ms)"),
			g.InputInt(&ctx.Config.PressDuration).OnChange(ctx.Config.Save),
		),

		g.Dummy(8.0, V_SPACING),
		g.Label("Button Triggers"),
		g.Dummy(8.0, V_SPACING/3),
		g.Row(
			g.Column(
				g.Row(
					g.Label("A:"),
				),

				g.Row(
					g.Label("B:"),
				),
			),

			g.Column(
				g.InputText(&ctx.Config.ButtonTriggers.A).OnChange(ctx.Config.Save).Size(INPUT_SIZE),
				g.InputText(&ctx.Config.ButtonTriggers.B).OnChange(ctx.Config.Save).Size(INPUT_SIZE),
			),

			g.Dummy(8.0, 1.0),

			g.Column(
				g.Row(
					g.Label("X:"),
				),

				g.Row(
					g.Label("Y:"),
				),

				g.Row(
					g.Label("START:"),
				),
			),

			g.Column(
				g.InputText(&ctx.Config.ButtonTriggers.X).OnChange(ctx.Config.Save).Size(INPUT_SIZE),
				g.InputText(&ctx.Config.ButtonTriggers.Y).OnChange(ctx.Config.Save).Size(INPUT_SIZE),
				g.InputText(&ctx.Config.ButtonTriggers.START).OnChange(ctx.Config.Save).Size(INPUT_SIZE),
			),

			g.Dummy(8.0, 1.0),

			g.Column(
				g.Row(
					g.Label("L:"),
				),

				g.Row(
					g.Label("R:"),
				),

				g.Row(
					g.Label("Z:"),
				),
			),

			g.Column(
				g.InputText(&ctx.Config.ButtonTriggers.L).OnChange(ctx.Config.Save).Size(INPUT_SIZE),
				g.InputText(&ctx.Config.ButtonTriggers.R).OnChange(ctx.Config.Save).Size(INPUT_SIZE),
				g.InputText(&ctx.Config.ButtonTriggers.Z).OnChange(ctx.Config.Save).Size(INPUT_SIZE),
			),
		),

		g.Dummy(1.0, V_SPACING/2),

		g.Row(
			g.Column(
				g.Row(
					g.Label("UP:"),
				),
				g.Row(
					g.Label("DOWN:"),
				),
			),

			g.Column(
				g.InputText(&ctx.Config.ButtonTriggers.UP).OnChange(ctx.Config.Save).Size(INPUT_SIZE*1.5),
				g.InputText(&ctx.Config.ButtonTriggers.DOWN).OnChange(ctx.Config.Save).Size(INPUT_SIZE*1.5),
			),

			g.Dummy(8.0, 1.0),

			g.Column(
				g.Row(
					g.Label("LEFT:"),
				),
				g.Row(
					g.Label("RIGHT:"),
				),
			),

			g.Column(
				g.InputText(&ctx.Config.ButtonTriggers.LEFT).OnChange(ctx.Config.Save).Size(INPUT_SIZE*1.5),
				g.InputText(&ctx.Config.ButtonTriggers.RIGHT).OnChange(ctx.Config.Save).Size(INPUT_SIZE*1.5),
			),
		),

		g.Dummy(8.0, 32.0),
		g.Custom(func() {
			width, _ := g.GetAvailableRegion()

			buttonText := "Serial Connect"
			onClick := ctx.InitComms

			if ctx.CommsConnected {
				buttonText = "Serial Disconnect"
				onClick = ctx.CloseComms
			}

			buttonWidth, buttonHeight := g.CalcTextSize(buttonText)
			buttonWidth += 16
			buttonHeight += 16

			g.Row(
				g.Dummy(width-buttonWidth, 0),
				g.Button(buttonText).OnClick(onClick),
			).Build()
		}),
		g.Custom(func() {
			width, _ := g.GetAvailableRegion()

			buttonText := "Twitch Connect"
			onClick := ctx.ConnectTwitch

			if ctx.TwitchClient != nil {
				buttonText = "Twitch Disconnect"
				onClick = ctx.DisconnectTwitch
			}

			buttonWidth, buttonHeight := g.CalcTextSize(buttonText)
			buttonWidth += 16
			buttonHeight += 16

			g.Row(
				g.Dummy(width-buttonWidth, 0),
				g.Button(buttonText).OnClick(func() { go onClick() }),
			).Build()
		}),
		g.Label("Status: "+ctx.Status),
	)
}

func uiStart(ctx *Context) {
	wnd := g.NewMasterWindow("Twitch to Gamecube Adapter", 580, 610, g.MasterWindowFlagsFloating+g.MasterWindowFlagsNotResizable)
	wnd.Run(func() {
		loop(ctx)
	})
}
