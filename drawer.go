package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

func printAt(x, y int, s string, fg, bg termbox.Attribute) {
	for _, r := range s {
		termbox.SetCell(x, y, r, fg, bg)
		x++
	}
}

func printAtDef(x, y int, s string) {
	printAt(x, y, s, termbox.ColorDefault, termbox.ColorDefault)
}

type BoardDrawer interface {
	DrawScores(scr, hi int)
	DrawBoardNow(gb *GameBoard)
	StartSlideTiles(tiles []Move, dir Direction)
	DoneSlideTiles() bool
}

// simple Drawer with no animation
type SimpleDrawer struct {
}

func (sd SimpleDrawer) DrawScores(scr, hi int) {
	printAtDef(0, 0, fmt.Sprintf("Scrore: %05d  Hi-Score: %05d", scr, hi))
}

func (sd SimpleDrawer) DrawBoardNow(gb *GameBoard) {

	for y := 0; y < BoardSize; y++ {
		for x := 0; x < BoardSize; x++ {
			printAtDef(x*6, y*2+1, fmt.Sprintf("%5d", gb[x][y].val))
		}
	}
}

func (sd SimpleDrawer) StartSlideTiles(tiles []Move, dir Direction) {

}

func (sd SimpleDrawer) DoneSlideTiles() bool {
	return true
}
