package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var win = windows.NewLazyDLL("user32.dll")
var getWindowText = win.NewProc("GetWindowTextW")
var getWindowTextLenght = win.NewProc("GetWindowTextLengthW")

type (
	HANDLE uintptr
	HWND   HANDLE
)

func main() {
	for {
		hwnd := tW("GetForegroundWindow")
		getText := getText(HWND(hwnd))
		fmt.Printf("Now you're in->>  %s // hwnd-> %v \n", getText, hwnd)
	}
}

func tW(str string) uintptr {
	pr := win.NewProc(str)
	hwnd, _, _ := pr.Call()
	return hwnd
}

func getText(hwnd HWND) string {
	len, _, _ := getWindowTextLenght.Call(uintptr(hwnd))
	r := int(len) + 1 //have to +1
	buf := make([]uint16, r)
	getWindowText.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&buf[0])), uintptr(r))
	return syscall.UTF16ToString(buf)
}
