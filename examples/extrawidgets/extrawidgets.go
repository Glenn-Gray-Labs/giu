package main

import (
	g "github.com/Glenn-Gray-Labs/giu"
)

var (
	showPD bool    = true
	radius float32 = 20
)

func loop() {
	g.SingleWindow("Extra Widgets", g.Layout{
		g.Checkbox("Show ProgressIndicator", &showPD, nil),
		g.Condition(showPD, g.Layout{
			g.SliderFloat("Radius", &radius, 10, 100, ""),
			g.Line(
				g.ProgressIndicator("pd1", "", 20+radius, 20+radius, radius),
				g.ProgressIndicator("pd2", "", 20+radius, 20+radius, radius),
			),
			g.ProgressIndicator("pd3", "Loading...", 0, 0, radius),
		}, nil),
	})
}

func main() {
	wnd := g.NewMasterWindow("Extra Widgets", 800, 600, g.MasterWindowFlagsNotResizable, nil)
	wnd.Main(loop)
}
