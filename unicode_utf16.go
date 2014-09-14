package chardet

// utf-16 big ending
// [\x00-\xD7\xE0-\xFF][\x00-\xFF]
// [\xD8-\xDB][\x00-\xFF][\xDC-\DF][\x00-\xFF]
type utf16be struct {
	byte
}

func (u utf16be) String() string {
	return "utf16-be"
}

func (u *utf16be) Feed(x byte) bool {
	switch u.byte {
	case 0:
		if (x >= 0x00 && x <= 0xD7) || (x >= 0xE0 && x <= 0xFF) {
			u.byte = 1
			return true
		}
		if x >= 0xD8 && x <= 0xDB {
			u.byte = 2
			return true
		}
	case 1:
		u.byte = 0
		return true
	case 2:
		u.byte = 3
		return true
	default:
		if x >= 0xDC && x <= 0xDF {
			u.byte = 1
			return true
		}
	}
	return false
}

// utf-16 little ending
// [\x00-\xFF][\x00-\xD7\xE0-\xFF]
// [\x00-\xFF][\xD8-\xDB][\x00-\xFF][\xDC-\DF]
type utf16le struct {
	byte
}

func (u utf16le) String() string {
	return "utf16-le"
}

func (u *utf16le) Feed(x byte) bool {
	switch u.byte {
	case 0:
		u.byte = 1
		return true
	case 1:
		if (x >= 0x00 && x <= 0xD7) || (x >= 0xE0 && x <= 0xFF) {
			u.byte = 0
			return true
		}
		if x >= 0xD8 && x <= 0xDB {
			u.byte = 2
			return true
		}
	case 2:
		u.byte = 3
		return true
	default:
		if x >= 0xDC && x <= 0xDF {
			u.byte = 0
			return true
		}
	}
	return false
}
