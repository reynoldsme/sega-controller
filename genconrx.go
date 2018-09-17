package main

import (
	"bufio"
	"fmt"
	"net"
	"syscall"
)

// find and load user32.dll
var dll = syscall.NewLazyDLL("user32.dll")

//  access to the specific named procedure
var procKeyBd = dll.NewProc("keybd_event")

func main() {

	// 192.168.11.36
	ip := "192.168.11.64"
	port := "9000"

	fmt.Println("GenConRX: Recieves commands from a Sega Genesis controler over TCP ")
	fmt.Println(" and emulates windows keypresses")
	fmt.Println("'use nc -lk 9000 < /dev/ttyUSB*' on the remote server")
	fmt.Println("Attempting to connect to remote server...")

	// Open TCP connection to remote host.
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("Connection failed. Exiting.")
		fmt.Println(err)
	} else {
		fmt.Println("Connection opened.")
	}

	scanner := bufio.NewScanner(conn)
	// For every '\n' terminated line recieved, check the state and emulate a key event.
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		line := scanner.Text()

		// get the state of each button
		u := line[1:2]
		d := line[2:3]
		l := line[3:4]
		r := line[4:5]
		c := line[5:6]

		// send an emulated win32 keydown or keyup
		doKeys(u, VK_UP)
		doKeys(d, VK_DOWN)
		doKeys(l, VK_LEFT)
		doKeys(r, VK_RIGHT)
		doKeys(c, VK_A)

		//fmt.Println(u + d + l + r + c)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}

func doKeys(key string, keycode int) {
	if key != "0" {
		downKey(keycode)
		fmt.Println(key + " pressed")
	} else {
		upKey(keycode)
		fmt.Println(key + " released")
	}
}

// send a win32 keydown event
func downKey(key int) {
	vkey := key + 0x80
	procKeyBd.Call(uintptr(key), uintptr(vkey), 0, 0)
}

// send a win32 keyup event
func upKey(key int) {
	vkey := key + 0x80
	procKeyBd.Call(uintptr(key), uintptr(vkey), _KEYEVENTF_KEYUP, 0)
}

// win32 modifier keymap
const (
	_VK_SHIFT        = 0x10
	_VK_CTRL         = 0x11
	_VK_ALT          = 0x12
	_VK_LSHIFT       = 0xA0
	_VK_RSHIFT       = 0xA1
	_VK_LCONTROL     = 0xA2
	_VK_RCONTROL     = 0xA3
	_KEYEVENTF_KEYUP = 0x0002
)

// win32 base keymap
const (
	VK_BACK  = 0x08
	VK_TAB   = 0x09
	VK_ENTER = 0x0D
	VK_ESC   = 0x1B
	VK_LEFT  = 0x25
	VK_UP    = 0x26
	VK_RIGHT = 0x27
	VK_DOWN  = 0x28
	VK_A     = 0x41
	VK_B     = 0x42
	VK_C     = 0x43
	VK_D     = 0x44
	VK_E     = 0x45
	VK_F     = 0x46
	VK_G     = 0x47
	VK_H     = 0x48
	VK_I     = 0x49
	VK_J     = 0x4A
	VK_K     = 0x4B
	VK_L     = 0x4C
	VK_M     = 0x4D
	VK_N     = 0x4E
	VK_O     = 0x4F
	VK_P     = 0x50
	VK_Q     = 0x51
	VK_R     = 0x52
	VK_S     = 0x53
	VK_T     = 0x54
	VK_U     = 0x55
	VK_V     = 0x56
	VK_W     = 0x57
	VK_X     = 0x58
	VK_Y     = 0x59
	VK_Z     = 0x5A
	VK_LWIN  = 0x5B
	VK_RWIN  = 0x5C
)
