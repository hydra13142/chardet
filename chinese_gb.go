package chardet

var gbk_map = map[rune]int{
	0xb5c4: 0x0, // 的
	0xd2bb: 0x1, // 一
	0xcac7: 0x2, // 是
	0xc1cb: 0x3, // 了
	0xd4da: 0x4, // 在
	0xc8cb: 0x5, // 人
	0xb2bb: 0x6, // 不
	0xb9fa: 0x7, // 国
	0xd3d0: 0x8, // 有
	0xd6d0: 0x9, // 中
	0xcbfb: 0xa, // 他
	0xced2: 0xb, // 我
	0xbacd: 0xc, // 和
	0xb4f3: 0xd, // 大
	0xb8f6: 0xe, // 个
	0xc9cf: 0xf, // 上
}

// gb2312 (chinese)
// [\x00-\x7F]
// [\xA1-\xF7][\xA1-\xFE]
type gb2312 struct {
	byte
}

func (g gb2312) String() string {
	return "gb2312"
}
func (g *gb2312) Feed(x byte) bool {
	if g.byte == 0 {
		if x >= 0x00 && x <= 0x7F {
			return true
		}
		if x >= 0xA1 && x <= 0xF7 {
			g.byte = 1
			return true
		}
	} else {
		if x >= 0xA1 && x <= 0xFE {
			g.byte = 0
			return true
		}
	}
	return false
}

func (g gb2312) Priority() float64 {
	return 0
}

// gbk (chinese)
// [\x00-\x7F]
// [\x81-\xFE][\x40-\x7E\x80-\xFE]
type gbk struct {
	byte
	rune
	freq [16]int
}

func (g gbk) String() string {
	return "gbk"
}

func (g *gbk) Feed(x byte) (ans bool) {
	defer func() {
		if ans && g.byte == 0 {
			if i, ok := gbk_map[g.rune]; ok {
				g.freq[i]++
			}
		}
	}()
	if g.byte == 0 {
		if x >= 0x00 && x <= 0x7F {
			g.rune = rune(x)
			return true
		}
		if x >= 0x81 && x <= 0xFE {
			g.byte = 1
			g.rune = rune(x) << 8
			return true
		}
	} else {
		if (x >= 0x40 && x <= 0x7E) || (x >= 0x80 && x <= 0xFE) {
			g.byte = 0
			g.rune |= rune(x)
			return true
		}
	}
	return false
}

func (g *gbk) Priority() float64 {
	s, f := 0, 0.0
	for _, x := range g.freq {
		s += x
	}
	if s == 0 {
		return 0
	}
	for i, x := range g.freq {
		k := float64(x)/float64(s)*100 - freq[i]
		if k >= 0 {
			f += k
		} else {
			f -= k
		}
	}
	return 1 - f/100
}

// gb18030 (chinese)
// [\x00-\x7F]
// [\x81-\xFE][\x40-\x7E\x80-\xFE]
// [\x81-\xFE][\x30-\x39][\x81-\xFE][\x30-\x39]
type gb18030 struct {
	byte
}

func (g gb18030) String() string {
	return "gb18030"
}

func (g *gb18030) Feed(x byte) bool {
	switch g.byte {
	case 0:
		if x >= 0x00 && x <= 0x7F {
			return true
		}
		if x >= 0x81 && x <= 0xFE {
			g.byte = 1
			return true
		}
	case 1:
		if (x >= 0x40 && x <= 0x7E) || (x >= 0x80 && x <= 0xFE) {
			g.byte = 0
			return true
		}
		if x >= 0x30 && x <= 0x39 {
			g.byte = 2
			return true
		}
	case 2:
		if x >= 0x81 && x <= 0xFE {
			g.byte = 3
			return true
		}
	default:
		if x >= 0x30 && x <= 0x39 {
			g.byte = 0
			return true
		}
	}
	return false
}

func (g gb18030) Priority() float64 {
	return 0
}
