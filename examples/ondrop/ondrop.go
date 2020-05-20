package main

import (
	"fmt"
	"strings"

	g "github.com/Glenn-Gray-Labs/giu"
)

var (
	dropInFiles string
)

func loop() {
	g.SingleWindow("On Drop Demo", g.Layout{
		g.Label("Drop file to this window"),
		g.InputTextMultiline("#DroppedFiles", &dropInFiles, -1, -1, g.InputTextFlagsReadOnly, nil, nil),
	})
}

func onDrop(names []string) {
	var sb strings.Builder
	for _, n := range names {
		sb.WriteString(fmt.Sprintf("%s\n", n))
	}

	dropInFiles = sb.String()
	g.Update()
}

func main() {
	wnd := g.NewMasterWindow("On Drop Demo", 600, 400, g.MasterWindowFlagsNotResizable, nil)
	wnd.SetDropCallback(onDrop)
	wnd.Main(loop)
}
