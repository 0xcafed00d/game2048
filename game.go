package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/simulatedsimian/rect"
	"time"
)

func main() {

	start := time.Now()
	for {
		time.Sleep(time.Second)
		fmt.Println((time.Now().UnixNano() - start.UnixNano()) / int64(time.Millisecond))
	}

	js, jserr := OpenJoystick(0)

	if jserr == nil {
		fmt.Println("Axis Count: ", js.AxisCount())
		fmt.Println("Button Count: ", js.ButtonCount())
		fmt.Println("Name: ", js.Name())

		for {
			jinfo := js.Read()
			fmt.Printf("%v\n", jinfo)
			time.Sleep(time.Millisecond * 100)
		}
	} else {
		fmt.Println(jserr)
		return
	}

	gc := GameCore{}

	var r rect.Rectangle
	x, y := 0, 0
	dx, dy := 1, 1

	b := MakeMemBuffer(10, 10)
	Fill(b, '0', termbox.ColorRed, termbox.ColorBlue, ALL)

	gc.OnInit = func(gc *GameCore) error {
		r = rect.WH(gc.BackBuffer.Size())
		r.Expand(rect.Vec{-1, -1})
		return nil
	}

	gc.OnEvent = func(gc *GameCore, ev *termbox.Event) error {
		if ev.Type == termbox.EventKey {
			if ev.Ch == 'q' {
				gc.DoQuit = true
			}
		}
		return nil
	}

	gc.OnTick = func(gc *GameCore) error {
		FillArea(gc.BackBuffer, r, 'x', termbox.ColorCyan, termbox.ColorGreen, ALL)
		BlitBuffer(b, gc.BackBuffer, x, y)
		x += dx
		y += dy

		if x > 60 || x <= 0 {
			dx = -dx
		}
		if y > 30 || y <= 0 {
			dy = -dy
		}
		return nil
	}

	err := gc.Run()
	if err != nil {
		panic(err)
	}
}
