package main

import (
	"fmt"

	"golang.org/x/sys/windows"
)

var (
	win = windows.NewLazyDLL("user32.dll")
)

func main() {
	for {
		hwnd := tW("GetForegroundWindow")
		fmt.Println(hwnd)
	}
}

func tW(str string) uintptr {
	pr := win.NewProc(str)
	hwnd, _, _ := pr.Call()
	return hwnd
}
