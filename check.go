package chardet

type detect interface {
	String() string
	Feed(byte) bool
	Priority() float64
}

func check(data []byte, lst []detect) []detect {
	for _, c := range data {
		for i, l := 0, len(lst); i < l; {
			if !lst[i].Feed(c) {
				copy(lst[i:], lst[i+1:])
				l--
				lst = lst[:l]
			} else {
				i++
			}
		}
	}
	if len(lst) == 0 {
		return nil
	}
	return lst
}

func Check(data []byte) string {
	if s := checkbom(data); s != "" {
		return s
	}
	lb := check(data, []detect{&utf8{}, &utf16BE{}, &utf16LE{}, &utf32BE{}, &utf32LE{}, &hzgb2312{}})
	if len(lb) > 0 {
		x, y := -1, -100.0
		for i, l := range lb {
			if r := l.Priority(); y < r {
				x, y = i, r
			}
		}
		return lb[x].String()
	}
	lp := check(data, []detect{&gbk{}, &big5{}, &eucJP{}, &shiftJIS{}, &iso2022JP{}, &eucKR{}, &gb18030{}})
	if len(lp) > 0 {
		x, y := -1, -100.0
		for i, l := range lp {
			if r := l.Priority(); y < r {
				x, y = i, r
			}
		}
		return lp[x].String()
	}
	return ""
}
