# Canvas
Canvas is a Go package for generating ascii images. See [GoDoc](https://godoc.org/github.com/ianlopshire/canvas) for details.

## Installation

```shell
$go get github.com/ianlopshire/canvas
```

## Usage

```go
import "github.com/ianlopshire/canvas"
```

### Example

```go
// Build an ascii hexagon using strings.
hexagonString := `***>-----<***` + "\n"
hexagonString += `**/       \**` + "\n"
hexagonString += `*/         \*` + "\n"
hexagonString += `<           >` + "\n"
hexagonString += `*\         /*` + "\n"
hexagonString += `**\       /**` + "\n"
hexagonString += `***>-----<***`

// The string is parsed into a new canvase.
// The second paramater is used to specify transparency.
hexagon := NewCanvasFromString(hexagonString, '*')

// Created a larger canvas that is completely filled with '~' runes.
board := NewFill(31, 21, '~')

// Arrange multiple hexagons onto the new canvas.
board.Compose(0, 4, hexagon)
board.Compose(0, 10, hexagon)
board.Compose(9, 1, hexagon)
board.Compose(9, 7, hexagon)
board.Compose(9, 13, hexagon)
board.Compose(18, 4, hexagon)
board.Compose(18, 10, hexagon)

// Print the board
fmt.Print(board)
```
*Outputs*
```
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
~~~~~~~~~~~~>-----<~~~~~~~~~~~~
~~~~~~~~~~~/       \~~~~~~~~~~~
~~~~~~~~~~/         \~~~~~~~~~~
~~~>-----<           >-----<~~~
~~/       \         /       \~~
~/         \       /         \~
<           >-----<           >
~\         /       \         /~
~~\       /         \       /~~
~~~>-----<           >-----<~~~
~~/       \         /       \~~
~/         \       /         \~
<           >-----<           >
~\         /       \         /~
~~\       /         \       /~~
~~~>-----<           >-----<~~~
~~~~~~~~~~\         /~~~~~~~~~~
~~~~~~~~~~~\       /~~~~~~~~~~~
~~~~~~~~~~~~>-----<~~~~~~~~~~~~
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
```
