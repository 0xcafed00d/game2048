package main

import (
	"github.com/nsf/termbox-go"
	"github.com/simulatedsimian/game2048/glib"
	"github.com/simulatedsimian/rect"
)

func main() {
	gc := glib.GameCore{}

	var r rect.Rectangle
	x, y := 0, 0
	dx, dy := 1, 1

	b := glib.MakeMemBuffer(10, 10)
	glib.Fill(b, '0', termbox.ColorRed, termbox.ColorBlue, glib.ALL)

	gc.OnInit = func(gc *glib.GameCore) error {
		r = rect.WH(gc.BackBuffer.Size())
		return nil
	}

	gc.OnEvent = func(gc *glib.GameCore, ev *termbox.Event) error {
		if ev.Type == termbox.EventKey {
			if ev.Ch == 'q' {
				gc.DoQuit = true
			}
		}
		return nil
	}

	gc.OnTick = func(gc *glib.GameCore) error {
		glib.FillArea(gc.BackBuffer, r, 'x', termbox.ColorCyan, termbox.ColorGreen, glib.ALL)
		glib.BlitBuffer(b, gc.BackBuffer, x, y)
		x += dx
		y += dy

		if x > 60 || x <= 0 {
			dx = -dx
		}
		if y > 30 || y <= 0 {
			dy = -dy
		}

		DrawBox(gc.BackBuffer, rect.XYWH(0, 10, 9, 5), 0)
		DrawBox(gc.BackBuffer, rect.XYWH(10, 10, 9, 5), 1)
		DrawBox(gc.BackBuffer, rect.XYWH(20, 10, 9, 5), 2)
		DrawBox(gc.BackBuffer, rect.XYWH(30, 10, 9, 5), 3)
		return nil
	}

	err := gc.Run()
	if err != nil {
		panic(err)
	}
}
