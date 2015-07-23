package main

import (
	"github.com/nsf/termbox-go"
	"github.com/simulatedsimian/rect"
)

var drawingChars = [][]rune{
	[]rune("++++--||"),
	[]rune("┌┐└┘──││"),
	[]rune("╭╮╰╯──││"),
	[]rune("▛▜▙▟▀▄▌▐"),
}

func DrawBox(buffer Buffer, area rect.Rectangle, mode int) {
	buffer.SetCell(area.Min.X, area.Min.Y, drawingChars[mode][0],
		termbox.ColorDefault, termbox.ColorDefault)
	buffer.SetCell(area.Max.X-1, area.Min.Y, drawingChars[mode][1],
		termbox.ColorDefault, termbox.ColorDefault)
	buffer.SetCell(area.Min.X, area.Max.Y-1, drawingChars[mode][2],
		termbox.ColorDefault, termbox.ColorDefault)
	buffer.SetCell(area.Max.X-1, area.Max.Y-1, drawingChars[mode][3],
		termbox.ColorDefault, termbox.ColorDefault)

	for x := area.Min.X + 1; x < area.Max.X-1; x++ {
		buffer.SetCell(x, area.Min.Y, drawingChars[mode][4],
			termbox.ColorDefault, termbox.ColorDefault)
		buffer.SetCell(x, area.Max.Y-1, drawingChars[mode][5],
			termbox.ColorDefault, termbox.ColorDefault)
	}

	for y := area.Min.Y + 1; y < area.Max.Y-1; y++ {
		buffer.SetCell(area.Min.X, y, drawingChars[mode][6],
			termbox.ColorDefault, termbox.ColorDefault)
		buffer.SetCell(area.Max.X-1, y, drawingChars[mode][7],
			termbox.ColorDefault, termbox.ColorDefault)
	}
}
