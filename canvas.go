// Package canvas provides utilities for drawing ascii images
package canvas

import (
	"strings"
)

// Canvas is an x by y grid of runes.
// The top left corner of the grid is (0,0)
type Canvas [][]NullRune

// NewCanvas returns a new instance of Canvas with the given demention.
// Each position in the canvas will be filled with a transparent rune.
func NewCanvas(width, height int) Canvas {
	c := make([][]NullRune, height)

	for i := range c {
		c[i] = make([]NullRune, width)
	}

	return c
}

// NewFill NewCanvas returns a new instance of Canvas with the given demention.
// Each position in the canvas will be filled with the given rune.
func NewFill(width, height int, value rune) Canvas {
	c := make([][]NullRune, height)

	for i := range c {
		row := make([]NullRune, width)
		for j := range row {
			row[j].Set(value)
		}
		c[i] = row
	}

	return c
}

// NewCanvasFromString returns a new instance of Canvas based on the provided string.
//
// The width of the new Canvas will be the length of the longest line.
// The height of the new canvas will be the number of lines in the provided string.
// Any rune in the string with a value equal to the transparent paramater will be treated as transparent.
// If a line in the provided string is shorter than the longest line, subsequent runes will be treated a transparent.
func NewCanvasFromString(s string, transparent rune) Canvas {
	lines := strings.Split(s, "\n")
	maxLen := 0

	for i := range lines {
		if len(lines[i]) > maxLen {
			maxLen = len(lines[i])
		}
	}

	c := NewCanvas(maxLen, len(lines))

	for y := range lines {
		rs := []rune(lines[y])

		for x := range rs {
			if rs[x] != transparent {
				c.Set(x, y, rs[x])
			}
		}
	}

	return c
}

// String returns a string representation of the Canvas
func (c Canvas) String() string {
	s := ""

	for _, l := range c {
		for _, r := range l {
			if v, ok := r.Value(); ok {
				s += string(v)
			} else {
				s += " "
			}
		}
		s += "\n"
	}

	return s
}

// Set sets the value at the given position
func (c Canvas) Set(x, y int, value rune) {
	c[y][x].Set(value)
}

// Clear sets the given position as transparent
func (c Canvas) Clear(x, y int) {
	c[y][x].Clear()
}

// Value returns the value at the given position.
// If the second return value is false, the position is transparent.
func (c Canvas) Value(x, y int) (rune, bool) {
	return c[y][x].Value()
}

// Height returns the height of the canvas
func (c Canvas) Height() int {
	return len(c)
}

// Width returns the width of the canvas
func (c Canvas) Width() int {
	return len(c[0])
}

// Compose layers the given Canvas on top of the current canvas.
// The (0,0) position of the given Canvas will start be placed at (xOffest, yOffest)
// Transparent runes in the given Canvas are ignored.
// Any runes that overflow off the edge of the current canvas will be ignored.
func (c Canvas) Compose(xOffest, yOffest int, area Canvas) {
	for i := range area {
		for j := range area[i] {
			if value, ok := area[i][j].Value(); ok {
				if i+yOffest < c.Height() && j+xOffest < c.Width() {
					c[i+yOffest][j+xOffest].Set(value)
				}
			}
		}
	}
}
