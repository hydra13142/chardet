package chardet

var big_map = map[rune]int{
	0xaaba: 0x0, // 的
	0xa440: 0x1, // 一
	0xac4f: 0x2, // 是
	0xa446: 0x3, // 了
	0xa662: 0x4, // 在
	0xa448: 0x5, // 人
	0xa4a3: 0x6, // 不
	0x98c7: 0x7, // 国
	0xa6b3: 0x8, // 有
	0xa4a4: 0x9, // 中
	0xa54c: 0xa, // 他
	0xa7da: 0xb, // 我
	0xa94d: 0xc, // 和
	0xa46a: 0xd, // 大
	0x9fc4: 0xe, // 个
	0xa457: 0xf, // 上
}

// big5 (chinese)
// [\x00-\x7F]
// [\xA1-\xF9][\x40-\x7E\xA1-\xFE]
type big5 struct {
	rune
	byte
	freq [16]int
}

func (b big5) String() string {
	return "big5"
}

func (b big5) Rune() rune {
	return b.rune
}

func (b *big5) Priority() float64 {
	s, f := 0, 0.0
	for _, x := range b.freq {
		s += x
	}
	if s == 0 {
		return 0
	}
	for i, x := range b.freq {
		k := float64(x)/float64(s)*100 - freq[i]
		if k >= 0 {
			f += k
		} else {
			f -= k
		}
	}
	return 1 - f/100
}

func (b *big5) Feed(x byte) (ans bool) {
	defer func() {
		if ans && b.byte == 0 {
			if i, ok := big_map[b.rune]; ok {
				b.freq[i]++
			}
		}
	}()
	if b.byte == 0 {
		if x >= 0x00 && x <= 0x7F {
			b.rune = rune(x)
			return true
		}
		if x >= 0xA1 && x <= 0xF9 {
			b.byte = 1
			b.rune = rune(x) << 8
			return true
		}
	} else {
		if (x >= 0x40 && x <= 0x7E) || (x >= 0xA1 && x <= 0xFE) {
			b.byte = 0
			b.rune |= rune(x)
			return true
		}
	}
	return false
}
