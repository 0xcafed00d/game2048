package main

import (
	"github.com/nsf/termbox-go"
	"github.com/simulatedsimian/rect"
	"time"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	tbuffer := TermboxBuffer{}

	r := rect.WH(tbuffer.Size())
	r.Expand(rect.Vec{-1, -1})
	FillArea(tbuffer, r, 'x', (uint16)(termbox.ColorCyan), (uint16)(termbox.ColorGreen), CHAR)

	b := MakeMemBuffer(10, 10)
	Fill(b, '0', (uint16)(termbox.ColorRed), (uint16)(termbox.ColorBlue), ALL)

	BlitBuffer(b, tbuffer, 1, 1)

	termbox.Flush()

	doQuit := false

	ticker := time.NewTicker(time.Millisecond * 10)

	x, y := 0, 0
	dx, dy := 1, 1

	for !doQuit {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				if ev.Ch == 'q' {
					doQuit = true
				}
				termbox.Flush()
			}

			if ev.Type == termbox.EventResize {
				termbox.Flush()
			}

		case <-ticker.C:
			FillArea(tbuffer, r, 'x', (uint16)(termbox.ColorCyan), (uint16)(termbox.ColorGreen), ALL)
			BlitBuffer(b, tbuffer, x, y)
			termbox.Flush()
			x += dx
			y += dy

			if x > 60 || x <= 0 {
				dx = -dx
			}
			if y > 30 || y <= 0 {
				dy = -dy
			}
		}
	}
}
