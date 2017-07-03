package renderer

import (
	"fmt"
	"io"
)

// Render text in output stream
func Render(text string, output io.Writer) {
	fmt.Fprintf(output, text)
}
