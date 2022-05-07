package main

import (
	g "github.com/AllenDang/giu"
	"go.bug.st/serial"
)

func availableSerialPorts() []string {
	ports, err := serial.GetPortsList()
	if err != nil {
		panic(err)
	}

	return ports
}

func onPortSelect(ctx *Context) {
	ctx.Config.SerialPortName = availableSerialPorts()[ctx.SelectedPort]
}

func loop(ctx *Context) {
	ports := availableSerialPorts()
	g.PushWindowPadding(16.0, 16.0)
	g.PushFramePadding(8.0, 8.0)

	g.SingleWindow().Layout(
		g.Row(
			g.Column(
				g.Label("Twitch Channel"),
				g.InputText(&ctx.Config.TwitchChannel).Size(g.Auto),
			),
		),
		g.Dummy(8.0, 8.0),
		g.Column(
			g.Label("Serial Port"),
			g.Combo("", ports[ctx.SelectedPort], ports, &ctx.SelectedPort).OnChange(func() {
				onPortSelect(ctx)
			}),
		),

		g.Dummy(8.0, 32.0),
		g.Row(
			g.Dummy(-100, 0),
			g.Button("Start Comms").OnClick(func() {
				go initComms(ctx)
			}),
		),
	)

	g.PopStyleV(2)
}

func uiStart(ctx *Context) {
	wnd := g.NewMasterWindow("Twitch to Gamecube Adapter", 640, 480, g.MasterWindowFlagsFloating+g.MasterWindowFlagsNotResizable)
	wnd.Run(func() {
		loop(ctx)
	})
}
