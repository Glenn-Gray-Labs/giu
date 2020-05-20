package main

import (
	g "github.com/Glenn-Gray-Labs/giu"
	"github.com/Glenn-Gray-Labs/giu/imgui"
)

func loop() {
	imgui.ShowDemoWindow(nil)
}

func main() {
	wnd := g.NewMasterWindow("Widgets", 1024, 768, 0, nil)
	wnd.Main(loop)
}
