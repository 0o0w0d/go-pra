package screen

var ScreenX = 680		//
var ScreenY = 720		//
var width = 1080

const ColorDepth = 16		//
const velocity = 8

func ResizeScreen(x, y int)  {		//
	ScreenX, ScreenY = x, y
}

func changeWidth(w int)  {
	width = w
}