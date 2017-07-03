package main

import (
	"os"

	"github.com/lstipakov/salakieli/renderer"
)

func main() {
	renderer.Render("zlata", os.Stdout)
}
