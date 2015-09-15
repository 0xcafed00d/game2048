package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nsf/termbox-go"
	"github.com/simulatedsimian/game2048/glib"
	//"github.com/simulatedsimian/rect"
)

func main() {
	gc := glib.GameCore{}
	gc.TickTime = 50 * time.Millisecond

	state := MakeGameState()
	state.Drawer = &SimpleDrawer{}

	gc.OnInit = func(gc *glib.GameCore) error {
		state.NewGame()
		return nil
	}

	gc.OnEvent = func(gc *glib.GameCore, ev *termbox.Event) error {
		if ev.Type == termbox.EventKey {

			switch {
			case ev.Ch == 'q' || ev.Ch == 'Q':
				gc.DoQuit = true
			case ev.Ch == 'r' || ev.Ch == 'R':
				state.NewGame()
			case ev.Ch == 'u' || ev.Ch == 'U':
				state.UndoMove()
			case ev.Key == termbox.KeyArrowLeft:
				state.DoMove(Left)
			case ev.Key == termbox.KeyArrowRight:
				state.DoMove(Right)
			case ev.Key == termbox.KeyArrowUp:
				state.DoMove(Up)
			case ev.Key == termbox.KeyArrowDown:
				state.DoMove(Down)
			}
		}
		return nil
	}

	gc.OnTick = func(gc *glib.GameCore) error {
		state.Tick()
		return nil
	}

	err := gc.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
