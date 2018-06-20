package main

import (
	"github.com/dgx80/snippets/clip"
)

func main() {

	clip.Prepare()
	s := clip.CutLastWord()
	clip.Paste("$" + s)
}
