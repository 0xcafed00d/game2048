package main

import (
	"github.com/nsf/termbox-go"
	"github.com/simulatedsimian/rect"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	tbuffer := TermboxBuffer{}

	r := rect.WH(tbuffer.Size())
	r.Expand(rect.Vec{-1, -1})
	FillArea(tbuffer, r, 'x', (uint16)(termbox.ColorCyan), (uint16)(termbox.ColorGreen), CHAR)

	b := MakeMemBuffer(10, 10)
	Fill(b, '0', (uint16)(termbox.ColorRed), (uint16)(termbox.ColorBlue), ALL)

	BlitBuffer(b, tbuffer, 1, 1)

	termbox.Flush()

	doQuit := false

	for !doQuit {
		ev := termbox.PollEvent()

		if ev.Type == termbox.EventKey {
			if ev.Ch == 'q' {
				doQuit = true
			}
			termbox.Flush()
		}

		if ev.Type == termbox.EventResize {
			termbox.Flush()
		}
	}

}
