package main

import (
	"github.com/nsf/termbox-go"
	"github.com/simulatedsimian/rect"
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

const (
	CHAR    = 1
	FG      = 2
	BG      = 4
	ALL     = CHAR | FG | BG
	ATTRIBS = FG | BG
)

func FillArea(dst Buffer, area rect.Rectangle, ch rune, fg, bg uint16, filltype int) {
	area, ok := rect.Intersection(rect.WH(dst.Size()), area)
	if !ok {
	}
}

func Fill(dst Buffer, ch rune, fg, bg uint16, filltype int) {
	FillArea(dst, rect.WH(dst.Size()), ch, fg, bg, filltype)
}

func BlitBuffer(src, dst Buffer, xpos, ypos int) {
	srcSz := rect.WH(src.Size())
	srcSz.Translate(rect.Vec{xpos, ypos})

	dstSz := rect.WH(dst.Size())

	if intersect, ok := rect.Intersection(srcSz, dstSz); ok {
		_ = intersect
	}
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
