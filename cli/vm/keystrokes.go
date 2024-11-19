/*
Copyright (c) 2014-2016 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vm

import (
	"context"
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

type hidKey struct {
	Code         int32
	ShiftPressed bool
}

// stolen from
// https://gist.github.com/MightyPork/6da26e382a7ad91b5496ee55fdc73db2#file-usb_hid_keys-h-L110
const (
	KEY_MOD_LCTRL          = 0x01
	KEY_MOD_LSHIFT         = 0x02
	KEY_MOD_LALT           = 0x04
	KEY_MOD_LMETA          = 0x08
	KEY_MOD_RCTRL          = 0x10
	KEY_MOD_RSHIFT         = 0x20
	KEY_MOD_RALT           = 0x40
	KEY_MOD_RMETA          = 0x80
	KEY_NONE               = 0x00
	KEY_ERR_OVF            = 0x01
	KEY_A                  = 0x04
	KEY_B                  = 0x05
	KEY_C                  = 0x06
	KEY_D                  = 0x07
	KEY_E                  = 0x08
	KEY_F                  = 0x09
	KEY_G                  = 0x0a
	KEY_H                  = 0x0b
	KEY_I                  = 0x0c
	KEY_J                  = 0x0d
	KEY_K                  = 0x0e
	KEY_L                  = 0x0f
	KEY_M                  = 0x10
	KEY_N                  = 0x11
	KEY_O                  = 0x12
	KEY_P                  = 0x13
	KEY_Q                  = 0x14
	KEY_R                  = 0x15
	KEY_S                  = 0x16
	KEY_T                  = 0x17
	KEY_U                  = 0x18
	KEY_V                  = 0x19
	KEY_W                  = 0x1a
	KEY_X                  = 0x1b
	KEY_Y                  = 0x1c
	KEY_Z                  = 0x1d
	KEY_1                  = 0x1e
	KEY_2                  = 0x1f
	KEY_3                  = 0x20
	KEY_4                  = 0x21
	KEY_5                  = 0x22
	KEY_6                  = 0x23
	KEY_7                  = 0x24
	KEY_8                  = 0x25
	KEY_9                  = 0x26
	KEY_0                  = 0x27
	KEY_ENTER              = 0x28
	KEY_ESC                = 0x29
	KEY_BACKSPACE          = 0x2a
	KEY_TAB                = 0x2b
	KEY_SPACE              = 0x2c
	KEY_MINUS              = 0x2d
	KEY_EQUAL              = 0x2e
	KEY_LEFTBRACE          = 0x2f
	KEY_RIGHTBRACE         = 0x30
	KEY_BACKSLASH          = 0x31
	KEY_HASHTILDE          = 0x32
	KEY_SEMICOLON          = 0x33
	KEY_APOSTROPHE         = 0x34
	KEY_GRAVE              = 0x35
	KEY_COMMA              = 0x36
	KEY_DOT                = 0x37
	KEY_SLASH              = 0x38
	KEY_CAPSLOCK           = 0x39
	KEY_F1                 = 0x3a
	KEY_F2                 = 0x3b
	KEY_F3                 = 0x3c
	KEY_F4                 = 0x3d
	KEY_F5                 = 0x3e
	KEY_F6                 = 0x3f
	KEY_F7                 = 0x40
	KEY_F8                 = 0x41
	KEY_F9                 = 0x42
	KEY_F10                = 0x43
	KEY_F11                = 0x44
	KEY_F12                = 0x45
	KEY_SYSRQ              = 0x46
	KEY_SCROLLLOCK         = 0x47
	KEY_PAUSE              = 0x48
	KEY_INSERT             = 0x49
	KEY_HOME               = 0x4a
	KEY_PAGEUP             = 0x4b
	KEY_DELETE             = 0x4c
	KEY_END                = 0x4d
	KEY_PAGEDOWN           = 0x4e
	KEY_RIGHT              = 0x4f
	KEY_LEFT               = 0x50
	KEY_DOWN               = 0x51
	KEY_UP                 = 0x52
	KEY_NUMLOCK            = 0x53
	KEY_KPSLASH            = 0x54
	KEY_KPASTERISK         = 0x55
	KEY_KPMINUS            = 0x56
	KEY_KPPLUS             = 0x57
	KEY_KPENTER            = 0x58
	KEY_KP1                = 0x59
	KEY_KP2                = 0x5a
	KEY_KP3                = 0x5b
	KEY_KP4                = 0x5c
	KEY_KP5                = 0x5d
	KEY_KP6                = 0x5e
	KEY_KP7                = 0x5f
	KEY_KP8                = 0x60
	KEY_KP9                = 0x61
	KEY_KP0                = 0x62
	KEY_KPDOT              = 0x63
	KEY_102ND              = 0x64
	KEY_COMPOSE            = 0x65
	KEY_POWER              = 0x66
	KEY_KPEQUAL            = 0x67
	KEY_F13                = 0x68
	KEY_F14                = 0x69
	KEY_F15                = 0x6a
	KEY_F16                = 0x6b
	KEY_F17                = 0x6c
	KEY_F18                = 0x6d
	KEY_F19                = 0x6e
	KEY_F20                = 0x6f
	KEY_F21                = 0x70
	KEY_F22                = 0x71
	KEY_F23                = 0x72
	KEY_F24                = 0x73
	KEY_OPEN               = 0x74
	KEY_HELP               = 0x75
	KEY_PROPS              = 0x76
	KEY_FRONT              = 0x77
	KEY_STOP               = 0x78
	KEY_AGAIN              = 0x79
	KEY_UNDO               = 0x7a
	KEY_CUT                = 0x7b
	KEY_COPY               = 0x7c
	KEY_PASTE              = 0x7d
	KEY_FIND               = 0x7e
	KEY_MUTE               = 0x7f
	KEY_VOLUMEUP           = 0x80
	KEY_VOLUMEDOWN         = 0x81
	KEY_KPCOMMA            = 0x85
	KEY_RO                 = 0x87
	KEY_KATAKANAHIRAGANA   = 0x88
	KEY_YEN                = 0x89
	KEY_HENKAN             = 0x8a
	KEY_MUHENKAN           = 0x8b
	KEY_KPJPCOMMA          = 0x8c
	KEY_HANGEUL            = 0x90
	KEY_HANJA              = 0x91
	KEY_KATAKANA           = 0x92
	KEY_HIRAGANA           = 0x93
	KEY_ZENKAKUHANKAKU     = 0x94
	KEY_KPLEFTPAREN        = 0xb6
	KEY_KPRIGHTPAREN       = 0xb7
	KEY_LEFTCTRL           = 0xe0
	KEY_LEFTSHIFT          = 0xe1
	KEY_LEFTALT            = 0xe2
	KEY_LEFTMETA           = 0xe3
	KEY_RIGHTCTRL          = 0xe4
	KEY_RIGHTSHIFT         = 0xe5
	KEY_RIGHTALT           = 0xe6
	KEY_RIGHTMETA          = 0xe7
	KEY_MEDIA_PLAYPAUSE    = 0xe8
	KEY_MEDIA_STOPCD       = 0xe9
	KEY_MEDIA_PREVIOUSSONG = 0xea
	KEY_MEDIA_NEXTSONG     = 0xeb
	KEY_MEDIA_EJECTCD      = 0xec
	KEY_MEDIA_VOLUMEUP     = 0xed
	KEY_MEDIA_VOLUMEDOWN   = 0xee
	KEY_MEDIA_MUTE         = 0xef
	KEY_MEDIA_WWW          = 0xf0
	KEY_MEDIA_BACK         = 0xf1
	KEY_MEDIA_FORWARD      = 0xf2
	KEY_MEDIA_STOP         = 0xf3
	KEY_MEDIA_FIND         = 0xf4
	KEY_MEDIA_SCROLLUP     = 0xf5
	KEY_MEDIA_SCROLLDOWN   = 0xf6
	KEY_MEDIA_EDIT         = 0xf7
	KEY_MEDIA_SLEEP        = 0xf8
	KEY_MEDIA_COFFEE       = 0xf9
	KEY_MEDIA_REFRESH      = 0xfa
	KEY_MEDIA_CALC         = 0xfb
)

var hidKeyMap = map[string]int32{
	"KEY_MOD_LCTRL":          KEY_MOD_LCTRL,
	"KEY_MOD_LSHIFT":         KEY_MOD_LSHIFT,
	"KEY_MOD_LALT":           KEY_MOD_LALT,
	"KEY_MOD_LMETA":          KEY_MOD_LMETA,
	"KEY_MOD_RCTRL":          KEY_MOD_RCTRL,
	"KEY_MOD_RSHIFT":         KEY_MOD_RSHIFT,
	"KEY_MOD_RALT":           KEY_MOD_RALT,
	"KEY_MOD_RMETA":          KEY_MOD_RMETA,
	"KEY_NONE":               KEY_NONE,
	"KEY_ERR_OVF":            KEY_ERR_OVF,
	"KEY_A":                  KEY_A,
	"KEY_B":                  KEY_B,
	"KEY_C":                  KEY_C,
	"KEY_D":                  KEY_D,
	"KEY_E":                  KEY_E,
	"KEY_F":                  KEY_F,
	"KEY_G":                  KEY_G,
	"KEY_H":                  KEY_H,
	"KEY_I":                  KEY_I,
	"KEY_J":                  KEY_J,
	"KEY_K":                  KEY_K,
	"KEY_L":                  KEY_L,
	"KEY_M":                  KEY_M,
	"KEY_N":                  KEY_N,
	"KEY_O":                  KEY_O,
	"KEY_P":                  KEY_P,
	"KEY_Q":                  KEY_Q,
	"KEY_R":                  KEY_R,
	"KEY_S":                  KEY_S,
	"KEY_T":                  KEY_T,
	"KEY_U":                  KEY_U,
	"KEY_V":                  KEY_V,
	"KEY_W":                  KEY_W,
	"KEY_X":                  KEY_X,
	"KEY_Y":                  KEY_Y,
	"KEY_Z":                  KEY_Z,
	"KEY_1":                  KEY_1,
	"KEY_2":                  KEY_2,
	"KEY_3":                  KEY_3,
	"KEY_4":                  KEY_4,
	"KEY_5":                  KEY_5,
	"KEY_6":                  KEY_6,
	"KEY_7":                  KEY_7,
	"KEY_8":                  KEY_8,
	"KEY_9":                  KEY_9,
	"KEY_0":                  KEY_0,
	"KEY_ENTER":              KEY_ENTER,
	"KEY_ESC":                KEY_ESC,
	"KEY_BACKSPACE":          KEY_BACKSPACE,
	"KEY_TAB":                KEY_TAB,
	"KEY_SPACE":              KEY_SPACE,
	"KEY_MINUS":              KEY_MINUS,
	"KEY_EQUAL":              KEY_EQUAL,
	"KEY_LEFTBRACE":          KEY_LEFTBRACE,
	"KEY_RIGHTBRACE":         KEY_RIGHTBRACE,
	"KEY_BACKSLASH":          KEY_BACKSLASH,
	"KEY_HASHTILDE":          KEY_HASHTILDE,
	"KEY_SEMICOLON":          KEY_SEMICOLON,
	"KEY_APOSTROPHE":         KEY_APOSTROPHE,
	"KEY_GRAVE":              KEY_GRAVE,
	"KEY_COMMA":              KEY_COMMA,
	"KEY_DOT":                KEY_DOT,
	"KEY_SLASH":              KEY_SLASH,
	"KEY_CAPSLOCK":           KEY_CAPSLOCK,
	"KEY_F1":                 KEY_F1,
	"KEY_F2":                 KEY_F2,
	"KEY_F3":                 KEY_F3,
	"KEY_F4":                 KEY_F4,
	"KEY_F5":                 KEY_F5,
	"KEY_F6":                 KEY_F6,
	"KEY_F7":                 KEY_F7,
	"KEY_F8":                 KEY_F8,
	"KEY_F9":                 KEY_F9,
	"KEY_F10":                KEY_F10,
	"KEY_F11":                KEY_F11,
	"KEY_F12":                KEY_F12,
	"KEY_SYSRQ":              KEY_SYSRQ,
	"KEY_SCROLLLOCK":         KEY_SCROLLLOCK,
	"KEY_PAUSE":              KEY_PAUSE,
	"KEY_INSERT":             KEY_INSERT,
	"KEY_HOME":               KEY_HOME,
	"KEY_PAGEUP":             KEY_PAGEUP,
	"KEY_DELETE":             KEY_DELETE,
	"KEY_END":                KEY_END,
	"KEY_PAGEDOWN":           KEY_PAGEDOWN,
	"KEY_RIGHT":              KEY_RIGHT,
	"KEY_LEFT":               KEY_LEFT,
	"KEY_DOWN":               KEY_DOWN,
	"KEY_UP":                 KEY_UP,
	"KEY_NUMLOCK":            KEY_NUMLOCK,
	"KEY_KPSLASH":            KEY_KPSLASH,
	"KEY_KPASTERISK":         KEY_KPASTERISK,
	"KEY_KPMINUS":            KEY_KPMINUS,
	"KEY_KPPLUS":             KEY_KPPLUS,
	"KEY_KPENTER":            KEY_KPENTER,
	"KEY_KP1":                KEY_KP1,
	"KEY_KP2":                KEY_KP2,
	"KEY_KP3":                KEY_KP3,
	"KEY_KP4":                KEY_KP4,
	"KEY_KP5":                KEY_KP5,
	"KEY_KP6":                KEY_KP6,
	"KEY_KP7":                KEY_KP7,
	"KEY_KP8":                KEY_KP8,
	"KEY_KP9":                KEY_KP9,
	"KEY_KP0":                KEY_KP0,
	"KEY_KPDOT":              KEY_KPDOT,
	"KEY_102ND":              KEY_102ND,
	"KEY_COMPOSE":            KEY_COMPOSE,
	"KEY_POWER":              KEY_POWER,
	"KEY_KPEQUAL":            KEY_KPEQUAL,
	"KEY_F13":                KEY_F13,
	"KEY_F14":                KEY_F14,
	"KEY_F15":                KEY_F15,
	"KEY_F16":                KEY_F16,
	"KEY_F17":                KEY_F17,
	"KEY_F18":                KEY_F18,
	"KEY_F19":                KEY_F19,
	"KEY_F20":                KEY_F20,
	"KEY_F21":                KEY_F21,
	"KEY_F22":                KEY_F22,
	"KEY_F23":                KEY_F23,
	"KEY_F24":                KEY_F24,
	"KEY_OPEN":               KEY_OPEN,
	"KEY_HELP":               KEY_HELP,
	"KEY_PROPS":              KEY_PROPS,
	"KEY_FRONT":              KEY_FRONT,
	"KEY_STOP":               KEY_STOP,
	"KEY_AGAIN":              KEY_AGAIN,
	"KEY_UNDO":               KEY_UNDO,
	"KEY_CUT":                KEY_CUT,
	"KEY_COPY":               KEY_COPY,
	"KEY_PASTE":              KEY_PASTE,
	"KEY_FIND":               KEY_FIND,
	"KEY_MUTE":               KEY_MUTE,
	"KEY_VOLUMEUP":           KEY_VOLUMEUP,
	"KEY_VOLUMEDOWN":         KEY_VOLUMEDOWN,
	"KEY_KPCOMMA":            KEY_KPCOMMA,
	"KEY_RO":                 KEY_RO,
	"KEY_KATAKANAHIRAGANA":   KEY_KATAKANAHIRAGANA,
	"KEY_YEN":                KEY_YEN,
	"KEY_HENKAN":             KEY_HENKAN,
	"KEY_MUHENKAN":           KEY_MUHENKAN,
	"KEY_KPJPCOMMA":          KEY_KPJPCOMMA,
	"KEY_HANGEUL":            KEY_HANGEUL,
	"KEY_HANJA":              KEY_HANJA,
	"KEY_KATAKANA":           KEY_KATAKANA,
	"KEY_HIRAGANA":           KEY_HIRAGANA,
	"KEY_ZENKAKUHANKAKU":     KEY_ZENKAKUHANKAKU,
	"KEY_KPLEFTPAREN":        KEY_KPLEFTPAREN,
	"KEY_KPRIGHTPAREN":       KEY_KPRIGHTPAREN,
	"KEY_LEFTCTRL":           KEY_LEFTCTRL,
	"KEY_LEFTSHIFT":          KEY_LEFTSHIFT,
	"KEY_LEFTALT":            KEY_LEFTALT,
	"KEY_LEFTMETA":           KEY_LEFTMETA,
	"KEY_RIGHTCTRL":          KEY_RIGHTCTRL,
	"KEY_RIGHTSHIFT":         KEY_RIGHTSHIFT,
	"KEY_RIGHTALT":           KEY_RIGHTALT,
	"KEY_RIGHTMETA":          KEY_RIGHTMETA,
	"KEY_MEDIA_PLAYPAUSE":    KEY_MEDIA_PLAYPAUSE,
	"KEY_MEDIA_STOPCD":       KEY_MEDIA_STOPCD,
	"KEY_MEDIA_PREVIOUSSONG": KEY_MEDIA_PREVIOUSSONG,
	"KEY_MEDIA_NEXTSONG":     KEY_MEDIA_NEXTSONG,
	"KEY_MEDIA_EJECTCD":      KEY_MEDIA_EJECTCD,
	"KEY_MEDIA_VOLUMEUP":     KEY_MEDIA_VOLUMEUP,
	"KEY_MEDIA_VOLUMEDOWN":   KEY_MEDIA_VOLUMEDOWN,
	"KEY_MEDIA_MUTE":         KEY_MEDIA_MUTE,
	"KEY_MEDIA_WWW":          KEY_MEDIA_WWW,
	"KEY_MEDIA_BACK":         KEY_MEDIA_BACK,
	"KEY_MEDIA_FORWARD":      KEY_MEDIA_FORWARD,
	"KEY_MEDIA_STOP":         KEY_MEDIA_STOP,
	"KEY_MEDIA_FIND":         KEY_MEDIA_FIND,
	"KEY_MEDIA_SCROLLUP":     KEY_MEDIA_SCROLLUP,
	"KEY_MEDIA_SCROLLDOWN":   KEY_MEDIA_SCROLLDOWN,
	"KEY_MEDIA_EDIT":         KEY_MEDIA_EDIT,
	"KEY_MEDIA_SLEEP":        KEY_MEDIA_SLEEP,
	"KEY_MEDIA_COFFEE":       KEY_MEDIA_COFFEE,
	"KEY_MEDIA_REFRESH":      KEY_MEDIA_REFRESH,
	"KEY_MEDIA_CALC":         KEY_MEDIA_CALC,
}

var hidCharacterMap = map[string]hidKey{
	"a": {KEY_A, false},
	"b": {KEY_B, false},
	"c": {KEY_C, false},
	"d": {KEY_D, false},
	"e": {KEY_E, false},
	"f": {KEY_F, false},
	"g": {KEY_G, false},
	"h": {KEY_H, false},
	"i": {KEY_I, false},
	"j": {KEY_J, false},
	"k": {KEY_K, false},
	"l": {KEY_L, false},
	"m": {KEY_M, false},
	"n": {KEY_N, false},
	"o": {KEY_O, false},
	"p": {KEY_P, false},
	"q": {KEY_Q, false},
	"r": {KEY_R, false},
	"s": {KEY_S, false},
	"t": {KEY_T, false},
	"u": {KEY_U, false},
	"v": {KEY_V, false},
	"w": {KEY_W, false},
	"x": {KEY_X, false},
	"y": {KEY_Y, false},
	"z": {KEY_Z, false},
	"1": {KEY_1, false},
	"2": {KEY_2, false},
	"3": {KEY_3, false},
	"4": {KEY_4, false},
	"5": {KEY_5, false},
	"6": {KEY_6, false},
	"7": {KEY_7, false},
	"8": {KEY_8, false},
	"9": {KEY_9, false},
	"0": {KEY_0, false},
	"A": {KEY_A, true},
	"B": {KEY_B, true},
	"C": {KEY_C, true},
	"D": {KEY_D, true},
	"E": {KEY_E, true},
	"F": {KEY_F, true},
	"G": {KEY_G, true},
	"H": {KEY_H, true},
	"I": {KEY_I, true},
	"J": {KEY_J, true},
	"K": {KEY_K, true},
	"L": {KEY_L, true},
	"M": {KEY_M, true},
	"N": {KEY_N, true},
	"O": {KEY_O, true},
	"P": {KEY_P, true},
	"Q": {KEY_Q, true},
	"R": {KEY_R, true},
	"S": {KEY_S, true},
	"T": {KEY_T, true},
	"U": {KEY_U, true},
	"V": {KEY_V, true},
	"W": {KEY_W, true},
	"X": {KEY_X, true},
	"Y": {KEY_Y, true},
	"Z": {KEY_Z, true},
	"!": {KEY_1, true},
	"@": {KEY_2, true},
	"#": {KEY_3, true},
	"$": {KEY_4, true},
	"%": {KEY_5, true},
	"^": {KEY_6, true},
	"&": {KEY_7, true},
	"*": {KEY_8, true},
	"(": {KEY_9, true},
	")": {KEY_0, true},
	" ": {KEY_SPACE, false},
	"-": {KEY_MINUS, false},
	"_": {KEY_MINUS, true},
	"=": {KEY_EQUAL, false},
	"+": {KEY_EQUAL, true},
	"[": {KEY_LEFTBRACE, false},
	"{": {KEY_LEFTBRACE, true},
	"]": {KEY_RIGHTBRACE, false},
	"}": {KEY_RIGHTBRACE, true},
	`\`: {KEY_BACKSLASH, false},
	"|": {KEY_BACKSLASH, true},
	";": {KEY_SEMICOLON, false},
	":": {KEY_SEMICOLON, true},
	"'": {KEY_APOSTROPHE, false},
	`"`: {KEY_APOSTROPHE, true},
	"`": {KEY_GRAVE, false},
	"~": {KEY_GRAVE, true},
	",": {KEY_COMMA, false},
	"<": {KEY_COMMA, true},
	".": {KEY_DOT, false},
	">": {KEY_DOT, true},
	"/": {KEY_SLASH, false},
	"?": {KEY_SLASH, true},
}

type keystrokes struct {
	*flags.VirtualMachineFlag

	UsbHidCodeValue int32
	UsbHidCodes     string
	UsbHidString    string
	LeftControl     bool
	LeftShift       bool
	LeftAlt         bool
	LeftGui         bool
	RightControl    bool
	RightShift      bool
	RightAlt        bool
	RightGui        bool
}

func init() {
	cli.Register("vm.keystrokes", &keystrokes{})
}

func (cmd *keystrokes) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	f.StringVar(&cmd.UsbHidString, "s", "", "Raw String to Send")
	f.StringVar(&cmd.UsbHidCodes, "c", "", "USB HID Codes (hex) or aliases, comma separated")
	f.Var(flags.NewInt32(&cmd.UsbHidCodeValue), "r", "Raw USB HID Code Value (int32)")
	f.BoolVar(&cmd.LeftControl, "lc", false, "Enable/Disable Left Control")
	f.BoolVar(&cmd.LeftShift, "ls", false, "Enable/Disable Left Shift")
	f.BoolVar(&cmd.LeftAlt, "la", false, "Enable/Disable Left Alt")
	f.BoolVar(&cmd.LeftGui, "lg", false, "Enable/Disable Left Gui")
	f.BoolVar(&cmd.RightControl, "rc", false, "Enable/Disable Right Control")
	f.BoolVar(&cmd.RightShift, "rs", false, "Enable/Disable Right Shift")
	f.BoolVar(&cmd.RightAlt, "ra", false, "Enable/Disable Right Alt")
	f.BoolVar(&cmd.RightGui, "rg", false, "Enable/Disable Right Gui")
}

func (cmd *keystrokes) Usage() string {
	return "VM"
}

func (cmd *keystrokes) Description() string {
	description := `Send Keystrokes to VM.

Examples:
 Default Scenario
  govc vm.keystrokes -vm $vm -s "root" 	# writes 'root' to the console
  govc vm.keystrokes -vm $vm -c 0x15 	# writes an 'r' to the console
  govc vm.keystrokes -vm $vm -r 1376263 # writes an 'r' to the console
  govc vm.keystrokes -vm $vm -c 0x28 	# presses ENTER on the console
  govc vm.keystrokes -vm $vm -c 0x4c -la=true -lc=true 	# sends CTRL+ALT+DEL to console
  govc vm.keystrokes -vm $vm -c 0x15,KEY_ENTER # writes an 'r' to the console and press ENTER

List of available aliases:
`
	keys := make([]string, 0)
	for key, _ := range hidKeyMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for i, key := range keys {
		if i > 0 {
			description += ", "
		}
		description += key
	}
	return description + "\n"
}

func (cmd *keystrokes) Process(ctx context.Context) error {
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *keystrokes) Run(ctx context.Context, f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return flag.ErrHelp
	}

	err = cmd.processUserInput(ctx, vm)
	if err != nil {
		return err
	}
	return nil
}

func (cmd *keystrokes) processUserInput(ctx context.Context, vm *object.VirtualMachine) error {
	if err := cmd.checkValidInputs(); err != nil {
		return err
	}

	codes, err := cmd.processUsbCode()

	if err != nil {
		return err
	}

	var keyEventArray []types.UsbScanCodeSpecKeyEvent
	for _, code := range codes {
		leftShiftSetting := false
		if code.ShiftPressed || cmd.LeftShift {
			leftShiftSetting = true
		}
		modifiers := types.UsbScanCodeSpecModifierType{
			LeftControl:  &cmd.LeftControl,
			LeftShift:    &leftShiftSetting,
			LeftAlt:      &cmd.LeftAlt,
			LeftGui:      &cmd.LeftGui,
			RightControl: &cmd.RightControl,
			RightShift:   &cmd.RightShift,
			RightAlt:     &cmd.RightAlt,
			RightGui:     &cmd.RightGui,
		}
		keyEvent := types.UsbScanCodeSpecKeyEvent{
			UsbHidCode: code.Code,
			Modifiers:  &modifiers,
		}
		keyEventArray = append(keyEventArray, keyEvent)
	}

	spec := types.UsbScanCodeSpec{
		KeyEvents: keyEventArray,
	}

	_, err = vm.PutUsbScanCodes(ctx, spec)

	return err
}

func (cmd *keystrokes) processUsbCode() ([]hidKey, error) {
	if cmd.rawCodeProvided() {
		return []hidKey{{cmd.UsbHidCodeValue, false}}, nil
	}

	if cmd.hexCodeProvided() {
		var retKeyArray []hidKey
		for _, c := range strings.Split(cmd.UsbHidCodes, ",") {
			var s int32
			lookupvalue, ok := hidKeyMap[c]
			if ok {
				s = intToHidCode(lookupvalue)
			} else {
				var err error
				s, err = hexStringToHidCode(c)
				if err != nil {
					return nil, err
				}
			}
			retKeyArray = append(retKeyArray, hidKey{s, false})
		}
		return retKeyArray, nil
	}

	if cmd.stringProvided() {
		var retKeyArray []hidKey
		for _, c := range cmd.UsbHidString {
			lookupValue, ok := hidCharacterMap[string(c)]
			if !ok {
				return nil, fmt.Errorf("invalid Character %s in String: %s", string(c), cmd.UsbHidString)
			}
			lookupValue.Code = intToHidCode(lookupValue.Code)
			retKeyArray = append(retKeyArray, lookupValue)
		}
		return retKeyArray, nil
	}
	return nil, nil
}

func hexStringToHidCode(hex string) (int32, error) {
	s, err := strconv.ParseInt(hex, 0, 32)
	if err != nil {
		return 0, err
	}
	return intToHidCode(int32(s)), nil
}

func intToHidCode(v int32) int32 {
	var s int32 = v << 16
	s |= 7
	return s
}

func (cmd *keystrokes) checkValidInputs() error {
	// poor man's boolean XOR -> A xor B xor C = A'BC' + AB'C' + A'B'C + ABC
	if (!cmd.rawCodeProvided() && cmd.hexCodeProvided() && !cmd.stringProvided()) || // A'BC'
		(cmd.rawCodeProvided() && !cmd.hexCodeProvided() && !cmd.stringProvided()) || // AB'C'
		(!cmd.rawCodeProvided() && !cmd.hexCodeProvided() && cmd.stringProvided()) || // A'B'C
		(cmd.rawCodeProvided() && cmd.hexCodeProvided() && cmd.stringProvided()) { // ABC
		return nil
	}
	return fmt.Errorf("specify only 1 argument")
}

func (cmd keystrokes) rawCodeProvided() bool {
	return cmd.UsbHidCodeValue != 0
}

func (cmd keystrokes) hexCodeProvided() bool {
	return cmd.UsbHidCodes != ""
}

func (cmd keystrokes) stringProvided() bool {
	return cmd.UsbHidString != ""
}
