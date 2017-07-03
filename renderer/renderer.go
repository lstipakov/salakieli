package renderer

import (
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"strings"
)

const (
	letterWidth    = 60
	letterHeight   = 80
	letterInterval = letterWidth / 4
	margin         = 10
)

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

// Bresenham's algorithm, http://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
func drawLine(img *image.Paletted, x0 int, y0 int, x1 int, y1 int) {
	dx := abs(x1 - x0)
	dy := abs(y1 - y0)
	sx, sy := 1, 1
	if x0 >= x1 {
		sx = -1
	}
	if y0 >= y1 {
		sy = -1
	}
	err := dx - dy

	for {
		img.SetColorIndex(x0, y0, 1)
		if x0 == x1 && y0 == y1 {
			return
		}
		e2 := err * 2
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

func drawChar(ch byte, img *image.Paletted, startX int, startY int) {
	_, ok := dict[ch]
	if !ok {
		// do nothing if char does not exist in dict
		return
	}

	if (dict[ch] & edge1) != 0 {
		drawLine(img, startX, startY, startX+letterWidth/2, startY)
	}
	if (dict[ch] & edge2) != 0 {
		drawLine(img, startX+letterWidth/2, startY, startX+letterWidth, startY)
	}
	if (dict[ch] & edge3) != 0 {
		drawLine(img, startX+letterWidth, startY, startX+letterWidth, startY+letterHeight/2)
	}
	if (dict[ch] & edge4) != 0 {
		drawLine(img, startX+letterWidth, startY+letterHeight/2, startX+letterWidth, startY+letterHeight)
	}
	if (dict[ch] & edge5) != 0 {
		drawLine(img, startX+letterWidth/2, startY+letterHeight, startX+letterWidth, startY+letterHeight)
	}
	if (dict[ch] & edge6) != 0 {
		drawLine(img, startX, startY+letterHeight, startX+letterWidth/2, startY+letterHeight)
	}
	if (dict[ch] & edge7) != 0 {
		drawLine(img, startX, startY+letterHeight/2, startX, startY+letterHeight)
	}
	if (dict[ch] & edge8) != 0 {
		drawLine(img, startX, startY, startX, startY+letterHeight/2)
	}

	if (dict[ch] & edge9) != 0 {
		drawLine(img, startX, startY, startX+letterWidth/2, startY+letterHeight/2)
	}
	if (dict[ch] & edge10) != 0 {
		drawLine(img, startX+letterWidth/2, startY, startX+letterWidth/2, startY+letterHeight/2)
	}
	if (dict[ch] & edge11) != 0 {
		drawLine(img, startX+letterWidth/2, startY+letterHeight/2, startX+letterWidth, startY)
	}
	if (dict[ch] & edge12) != 0 {
		drawLine(img, startX+letterWidth/2, startY+letterHeight/2, startX+letterWidth, startY+letterHeight/2)
	}
	if (dict[ch] & edge13) != 0 {
		drawLine(img, startX+letterWidth/2, startY+letterHeight/2, startX+letterWidth, startY+letterHeight)
	}
	if (dict[ch] & edge14) != 0 {
		drawLine(img, startX+letterWidth/2, startY+letterHeight/2, startX+letterWidth/2, startY+letterHeight)
	}
	if (dict[ch] & edge15) != 0 {
		drawLine(img, startX, startY+letterHeight, startX+letterWidth/2, startY+letterHeight/2)
	}
	if (dict[ch] & edge16) != 0 {
		drawLine(img, startX, startY+letterHeight/2, startX+letterWidth/2, startY+letterHeight/2)
	}
}

// Render text in output stream
func Render(text string, output io.Writer) {
	text = strings.ToUpper(text)

	width := len(text)*letterWidth + (len(text)-1)*letterInterval + margin*2
	height := letterHeight + margin*2

	rect := image.Rect(0, 0, width, height)
	img := image.NewPaletted(rect, []color.Color{color.White, color.Black})

	for i := 0; i < len(text); i++ {
		ch := byte(text[i])
		startX := margin + i*(letterWidth+letterInterval)
		startY := margin
		drawChar(ch, img, startX, startY)
	}

	jpeg.Encode(output, img, nil)
}
