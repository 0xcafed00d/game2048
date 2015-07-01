package main

import (
	"github.com/nsf/termbox-go"
	//	"github.com/simulatedsimian/rect"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	doQuit := false

	for !doQuit {
		ev := termbox.PollEvent()

		if ev.Type == termbox.EventKey {
			termbox.Flush()
		}

		if ev.Type == termbox.EventResize {
			termbox.Flush()
		}
	}

}
