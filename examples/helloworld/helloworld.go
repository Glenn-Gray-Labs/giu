package main

import (
	"fmt"

	g "github.com/Glenn-Gray-Labs/giu"
)

func onClickMe() {
	fmt.Println("Hello world!")
}

func onImSoCute() {
	fmt.Println("Im sooooooo cute!!")
}

func loop() {
	g.SingleWindow("hello world", g.Layout{
		g.Label("Hello world from giu"),
		g.Line(
			g.Button("Click Me", onClickMe),
			g.Button("I'm so cute", onImSoCute)),
	})
}

func main() {
	wnd := g.NewMasterWindow("Hello world", 400, 200, g.MasterWindowFlagsNotResizable, nil)
	wnd.Main(loop)
}
