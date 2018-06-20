package clip

import "github.com/go-vgo/robotgo"

var clipMem string

func Prepare() {
	robotgo.MicroSleep(200)
	robotgo.KeyTap("escape")
	wait()
}

func Paste(s string) {
	storeClipboard()

	robotgo.WriteAll(s)
	robotgo.PasteStr(s)
	restoreClipboard()
}

func CutLastWord() string {
	storeClipboard()

	robotgo.KeyTap("left", "control", "shift")
	wait()
	robotgo.KeyTap("x", "control")
	wait()
	s, _ := robotgo.ReadAll()
	restoreClipboard()
	return s
}

func storeClipboard() {
	clipMem, _ = robotgo.ReadAll()
}

func restoreClipboard() {
	robotgo.WriteAll(clipMem)
}

func wait() {
	robotgo.MicroSleep(5)
}
