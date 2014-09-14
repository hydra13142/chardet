package chardet

type detect interface {
	String() string
	Feed(byte) bool
}

type detectplus interface {
	detect
	Priority() float64
}

func checkUnicode(data []byte) string {
	lst := []detect{&utf8{}, &utf16be{}, &utf16le{}, &utf32be{}, &utf32le{}}
	for _, c := range data {
		for i, l := 0, len(lst); i < l; {
			if !lst[i].Feed(c) {
				l--
				lst[i] = lst[l]
				lst = lst[:l]
			} else {
				i++
			}
		}
	}
	if len(lst) > 0 {
		return lst[0].String()
	}
	return ""
}

func checkChinese(data []byte) (string, float64) {
	lst := []detectplus{&gb2312{}, &gbk{}, &gb18030{}, &big5{}}
	for _, c := range data {
		for i, l := 0, len(lst); i < l; {
			if !lst[i].Feed(c) {
				l--
				lst[i] = lst[l]
				lst = lst[:l]
			} else {
				i++
			}
		}
	}
	if len(lst) != 0 {
		a, b := -10.0, 0
		for i, c := range lst {
			if p := c.Priority(); a < p {
				a, b = p, i
			}
		}
		return lst[b].String(), a
	}
	return "", -2
}

func checkJapanese(data []byte) (string, float64) {
	lst := []detectplus{newEucJP(), newShiftJIS()}
	for _, c := range data {
		for i, l := 0, len(lst); i < l; {
			if !lst[i].Feed(c) {
				l--
				lst[i] = lst[l]
				lst = lst[:l]
			} else {
				i++
			}
		}
	}
	if len(lst) != 0 {
		a, b := -10.0, 0
		for i, c := range lst {
			if p := c.Priority(); a < p {
				a, b = p, i
			}
		}
		return lst[b].String(), a
	}
	return "", -2
}

// Check用于检测文本的编码格式。
//
// 可检测"utf8", "utf16be", "utf16le", "utf32be", "utf32le"共计5种标准格式；
// 可检测"gb2312", "gbk", "gb18030", "big5"共计4种中文格式；
// 可检测"euc-jp", "shift-jis"共计2种日文格式；
// 不支持其他语种的编码格式。
//
// 对满足多种编码的文本，会采用字符分布进一步检测，取优先级最高的。
//
// 如果检测到，会返回上述的字符串说明格式；否则返回空字符串。
func Check(data []byte) string {
	var m, n string
	var p, q float64
	if len(data) == 0 {
		return ""
	}
	if m = checkBom(data); m != "" {
		return m
	}
	if m = checkUnicode(data); m != "" {
		return m
	}
	if m, p = checkChinese(data); m != "" && p > 0 {
		return m
	}
	if n, q = checkJapanese(data); n != "" && q > 0 {
		return n
	}
	if m == "" {
		return n
	}
	if n == "" {
		return m
	}
	if p > q {
		return m
	} else {
		return n
	}
}
