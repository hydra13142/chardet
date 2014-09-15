package chardet

type iso2022jp struct {
	byte
}

func (i iso2022jp) String() string {
	return "iso-2022-jp"
}

func (i *iso2022jp) Feed(x byte) bool {
	switch i.byte {
	case 0:
		if x == 0x1b {
			i.byte = 1
			return true
		} else {
			if (x >= 0x00 && x <= 0x7F) || (x >= 0xA1 && x <= 0xDF) {
				return true
			}
		}
	case 1:
		if x == '(' {
			i.byte = 2
			return true
		} else if x == '$' {
			i.byte = 3
			return true
		}
	case 2:
		if x == 'B' {
			i.byte = 0
			return true
		} else if x == 'J' {
			i.byte = 4
			return true
		}
	case 3:
		if x == '@' || x == 'B' {
			i.byte = 4
			return true
		}
	case 4:
		if x == 0x1b {
			i.byte = 1
			return true
		} else {
			if x >= 0x21 && x <= 0x7E {
				i.byte = 5
				return true
			}
		}
	case 5:
		if x >= 0x21 && x <= 0x7E {
			i.byte = 4
			return true
		}
	}
	return false
}