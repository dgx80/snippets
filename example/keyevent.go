package main

import (
	"fmt"
	"sync"

	"github.com/go-vgo/robotgo"
)

var wg sync.WaitGroup

func main() {

	wg.Add(1)
	fmt.Println("start")
	go kmet()
	wg.Wait()

	//clip.Prepare()
	//s := clip.CutLastWord()
	//clip.Paste("$" + s)
}

func kmet() {
	defer wg.Done()
	keve := robotgo.AddEvent("control+k")
	if keve == 0 {
		fmt.Println("you press...", "k")
	}
}
