package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/simulatedsimian/game2048/glib"
	"github.com/simulatedsimian/rect"
)

func printAt(buf glib.Buffer, x, y int, s string, fg, bg termbox.Attribute) {
	for _, r := range s {
		buf.SetCell(x, y, r, fg, bg)
		x++
	}
}

func printAtDef(buf glib.Buffer, x, y int, s string) {
	printAt(buf, x, y, s, termbox.ColorDefault, termbox.ColorDefault)
}

type BoardDrawer interface {
	DrawScores(scr, hi int)
	DrawBoardNow(gb *GameBoard)
	StartSlideTiles(tiles []Move, dir Direction)
	DoneSlideTiles() bool
	//	Size() (x, y int)
}

// simple Drawer with no animation
type SimpleDrawer struct {
}

func (sd SimpleDrawer) DrawScores(scr, hi int) {
	printAtDef(glib.TermboxBuffer, 0, 0, fmt.Sprintf("Scrore: %05d  Hi-Score: %05d", scr, hi))
}

func (sd SimpleDrawer) DrawBoardNow(gb *GameBoard) {

	for y := 0; y < BoardSize; y++ {
		for x := 0; x < BoardSize; x++ {
			printAtDef(glib.TermboxBuffer, x*6, y*2+1, fmt.Sprintf("%5d", gb[x][y].val))
		}
	}
}

func (sd SimpleDrawer) StartSlideTiles(tiles []Move, dir Direction) {
}

func (sd SimpleDrawer) DoneSlideTiles() bool {
	return true
}

// -----------------------------------------------------------------------------

type AnimatedDrawer struct {
	TileW   int
	TileH   int
	gb      GameBoard
	aniCntr int
	moving  []Move
	dir     Direction
}

func (d *AnimatedDrawer) drawTile(buf glib.Buffer, x, y, val int) {
	DrawBox(buf, rect.XYWH(x, y, d.TileW, d.TileH), 3)
	printAtDef(buf, x+1, y+d.TileH/2, fmt.Sprintf("%5d", val))
}

func (d *AnimatedDrawer) isTileMoving(x, y int) bool {
	for i := 0; i < len(d.moving); i++ {
		if x == d.moving[i].x && y == d.moving[i].y {
			return true
		}
	}
	return false
}

func (d *AnimatedDrawer) DrawScores(scr, hi int) {
	printAtDef(glib.TermboxBuffer, 0, 0, fmt.Sprintf("Scrore: %05d  Hi-Score: %05d", scr, hi))
}

func (d *AnimatedDrawer) DrawBoardNow(gb *GameBoard) {

	d.gb = *gb
	for y := 0; y < BoardSize; y++ {
		for x := 0; x < BoardSize; x++ {
			tilex := x * d.TileW
			tiley := y * d.TileH
			d.drawTile(glib.TermboxBuffer, tilex, tiley, gb[x][y].val)
		}
	}
}

func (d *AnimatedDrawer) StartSlideTiles(tiles []Move, dir Direction) {
	d.aniCntr = 0
	d.dir = dir
	d.moving = tiles
}

func (d *AnimatedDrawer) DoneSlideTiles() bool {
	return true
}
