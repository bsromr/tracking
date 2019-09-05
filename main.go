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

var hwnd2 uintptr

func main() {
	fmt.Println("Hi, I am waiting until you open any window")
	control := tW("GetForegroundWindow")
	control2 := getText(HWND(control))
	for {
		hwnd := tW("GetForegroundWindow")
		getText := getText(HWND(hwnd))
		if (control2 != getText && hwnd != 0 && hwnd != 262234 && hwnd != 65852) || (control != hwnd && hwnd != 0 && hwnd != 262234 && hwnd != 65852) {
			//Here I am checking the number of 262234 and 65852. These are
			//window transition and
			fmt.Printf("Now you're in->>  %s // hwnd-> %v \n", getText, hwnd)
		}
		control = hwnd
		control2 = getText
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
