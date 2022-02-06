package common

import (
	"regexp"
	"strings"
)

// convert a string of keys to an array of RemoteKeyCode
// special keys are: UP; DOWN; LEFT; RIGHT; HOME; BACK; VOLP; VOLM; MUTE; PWR; ENTER
// you can also directly use defined keys in types.go like KEYCODE_TV_INPUT_HDMI_1
// each special key is separated by a semicolon
// example: `A ;UP; tst;KEYCODE_TV_INPUT_HDMI_1` will be parsed as [RemoteKeyCode_KEYCODE_A, RemoteKeyCode_KEYCODE_SPACE, RemoteKeyCode_KEYCODE_DPAD_UP, RemoteKeyCode_KEYCODE_SPACE, RemoteKeyCode_KEYCODE_T, RemoteKeyCode_KEYCODE_S, RemoteKeyCode_KEYCODE_T, KEYCODE_TV_INPUT_HDMI_1]
func ParseKeys(data string) []RemoteKeyCode {
	ret := []RemoteKeyCode{}
	special := map[string]RemoteKeyCode{
		"UP":    RemoteKeyCode_KEYCODE_DPAD_UP,
		"DOWN":  RemoteKeyCode_KEYCODE_DPAD_DOWN,
		"LEFT":  RemoteKeyCode_KEYCODE_DPAD_LEFT,
		"RIGHT": RemoteKeyCode_KEYCODE_DPAD_RIGHT,
		"HOME":  RemoteKeyCode_KEYCODE_HOME,
		"BACK":  RemoteKeyCode_KEYCODE_BACK,
		"MUTE":  RemoteKeyCode_KEYCODE_MUTE,
		"VOLP":  RemoteKeyCode_KEYCODE_VOLUME_UP,
		"VOLM":  RemoteKeyCode_KEYCODE_VOLUME_DOWN,
		"PWR":   RemoteKeyCode_KEYCODE_POWER,
		"ENTER": RemoteKeyCode_KEYCODE_ENTER,
	}

	for _, k := range strings.Split(data, ";") {
		if len(k) > 8 && k[0:8] == "KEYCODE_" {
			// if KEYCODE is directly specified
			ret = append(ret, RemoteKeyCode(RemoteKeyCode_value[k]))
		} else {
			if val, ok := special[k]; ok {
				// if it is a special key
				ret = append(ret, val)
			} else {
				for _, x := range k {
					// loop over each character
					c := string(x)
					if regexp.MustCompile(`^[A-Z]{1}$`).MatchString(c) {
						// if A-Z
						ret = append(ret, RemoteKeyCode(RemoteKeyCode_value["KEYCODE_"+c]))
					} else if regexp.MustCompile(`^[0-9]{1}$`).MatchString(c) {
						// if digit
						ret = append(ret, RemoteKeyCode(RemoteKeyCode_value["KEYCODE_NUMPAD_"+c]))
					} else if c == " " {
						// if digit
						ret = append(ret, RemoteKeyCode_KEYCODE_SPACE)
					}
				}
			}
		}
	}
	return ret
}
