package chardet

// utf-32 big ending
// \x00[\x00-\x0F][\x00-\xFF]{2}
type utf32be struct {
	byte
}

func (u utf32be) String() string {
	return "utf32-be"
}

func (u *utf32be) Feed(x byte) bool {
	switch u.byte {
	case 0:
		if x == 0x00 {
			u.byte = 1
			return true
		}
	case 1:
		if x >= 0x00 && x <= 0x0F {
			u.byte = 2
			return true
		}
	case 2:
		u.byte = 3
		return true
	default:
		u.byte = 0
		return true
	}
	return false
}

// utf-32 little ending
// [\x00-\xFF]{2}[\x00-\x0F]\x00
type utf32le struct {
	byte
}

func (u utf32le) String() string {
	return "utf32-le"
}

func (u *utf32le) Feed(x byte) bool {
	switch u.byte {
	case 0:
		u.byte = 1
		return true
	case 1:
		u.byte = 2
		return true
	case 2:
		if x >= 0x00 && x <= 0x0F {
			u.byte = 3
			return true
		}
	default:
		if x == 0x00 {
			u.byte = 0
			return true
		}
	}
	return false
}
