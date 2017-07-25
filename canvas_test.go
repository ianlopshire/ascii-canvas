package canvas

import (
	"fmt"
)

func Example() {
	// Build an ascii hexagon
	hexagonString := `***>-----<***` + "\n"
	hexagonString += `**/       \**` + "\n"
	hexagonString += `*/         \*` + "\n"
	hexagonString += `<           >` + "\n"
	hexagonString += `*\         /*` + "\n"
	hexagonString += `**\       /**` + "\n"
	hexagonString += `***>-----<***`
	hexagon := NewCanvasFromString(hexagonString, '*')

	// Compose multiple hexagons onto a board
	board := NewFill(31, 21, '~')
	board.Compose(0, 4, hexagon)
	board.Compose(0, 10, hexagon)
	board.Compose(9, 1, hexagon)
	board.Compose(9, 7, hexagon)
	board.Compose(9, 13, hexagon)
	board.Compose(18, 4, hexagon)
	board.Compose(18, 10, hexagon)

	// print the board
	fmt.Print(board)

	// Output: ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// ~~~~~~~~~~~~>-----<~~~~~~~~~~~~
	// ~~~~~~~~~~~/       \~~~~~~~~~~~
	// ~~~~~~~~~~/         \~~~~~~~~~~
	// ~~~>-----<           >-----<~~~
	// ~~/       \         /       \~~
	// ~/         \       /         \~
	// <           >-----<           >
	// ~\         /       \         /~
	// ~~\       /         \       /~~
	// ~~~>-----<           >-----<~~~
	// ~~/       \         /       \~~
	// ~/         \       /         \~
	// <           >-----<           >
	// ~\         /       \         /~
	// ~~\       /         \       /~~
	// ~~~>-----<           >-----<~~~
	// ~~~~~~~~~~\         /~~~~~~~~~~
	// ~~~~~~~~~~~\       /~~~~~~~~~~~
	// ~~~~~~~~~~~~>-----<~~~~~~~~~~~~
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
}
