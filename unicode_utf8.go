package chardet

// utf-8
// [\x00-\x7F]
// [\xC0-\xDF][\x80-\xBF]
// [\xE0-\xEF][\x80-\xBF]{2}
// [\xF0-\xF7][\x80-\xBF]{3}
type utf8 struct {
	byte
}

func (u utf8) String() string {
	return "utf8"
}

func (u *utf8) Feed(x byte) bool {
	if u.byte == 0 {
		if x >= 0x00 && x <= 0x7F {
			return true
		}
		if x >= 0xC0 && x <= 0xDF {
			u.byte = 1
			return true
		}
		if x >= 0xE0 && x <= 0xEF {
			u.byte = 2
			return true
		}
		if x >= 0xF0 && x <= 0xF7 {
			u.byte = 3
			return true
		}
	} else {
		if x >= 0x80 && x <= 0xBF {
			u.byte -= 1
			return true
		}
	}
	return false
}
