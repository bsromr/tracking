package main

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	win                 = windows.NewLazyDLL("user32.dll")
	getWindowText       = win.NewProc("GetWindowTextW")
	getWindowTextLenght = win.NewProc("GetWindowTextLengthW")
)

type (
	HANDLE uintptr
	HWND   HANDLE
)

func main() {
	fmt.Println("Hi, I am waiting until you open any window")
	control := tW("GetForegroundWindow")
	control2 := getText(HWND(control))
	for {
		hwnd := tW("GetForegroundWindow")
		getText := getText(HWND(hwnd))
		if control != hwnd && control2 != "" && hwnd != 0 {
			fmt.Printf("Now you're in->>  %s // hwnd-> %v \n", getText, hwnd)
		}
		control = hwnd
		control2 = getText
		time.Sleep(500 * time.Millisecond)
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
