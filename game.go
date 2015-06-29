package main

import (
	"github.com/nsf/termbox-go"
	//"github.com/simulatedsimian/rect"
)

type Buffer interface {
	SetCell(x, y int, ch rune, fg, bg uint16)
	GetCell(x, y int) (rune, uint16, uint16)
	Size() (int, int)
}

type TermboxBuffer struct {
}

func (b *TermboxBuffer) SetCell(x, y int, ch rune, fg, bg uint16) {
	termbox.SetCell(x, y, ch, termbox.Attribute(fg), termbox.Attribute(bg))
}

func (b *TermboxBuffer) GetCell(x, y int) (rune, uint16, uint16) {
	cells := termbox.CellBuffer()
	w, _ := termbox.Size()
	var cell *termbox.Cell = &cells[y*w+x]
	return cell.Ch, uint16(cell.Fg), uint16(cell.Bg)
}

func (b *TermboxBuffer) Size() (int, int) {
	return termbox.Size()
}

type Cell struct {
	Ch rune
	Fg uint16
	Bg uint16
}

type MemBuffer struct {
	w, h  int
	cells []Cell
}

func MakeMemBuffer(w, h int) *MemBuffer {
	return &MemBuffer{w, h, make([]Cell, w*h)}
}

func (b *MemBuffer) SetCell(x, y int, ch rune, fg, bg uint16) {
	if x < 0 || x >= b.w {
		return
	}
	if y < 0 || y >= b.h {
		return
	}

	b.cells[y*b.w+x] = Cell{ch, fg, bg}
}

func (b *MemBuffer) GetCell(x, y int) (rune, uint16, uint16) {
	var cell *Cell = &b.cells[y*b.w+x]
	return cell.Ch, cell.Fg, cell.Bg
}

func (b *MemBuffer) Size() (int, int) {
	return b.w, b.h
}

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
