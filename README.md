chardet
=======

detect text encoding, like python chardet, but for go

本包用于探测文本的编码格式。

支持探测的格式有：
5种unicode格式：utf8、大小端utf16，大小端utf32；
4种中文格式：gb2312、gbk、gb18030、big5；
2种日文格式euc-jp、shift-jis；
不支持其他语种和编码格式的探测。

本包只有一个函数：

    func Check([]byte) string

该函数用于探测格式。如果无法探测到编码格式，会返回空字符串。
如果发现该文本符合多个编码格式，会优先返回unicode格式；
否则进一步检测字符分布，返回最匹配的。

各编码格式对应字符串如下

* **utf8**  "utf8"
* **utf16 big-ending** "utf16be"
* **utf16 little-ending** "utf16le"
* **utf32 big-ending** "utf32be"
* **utf32 little-ending** "utf32le"
* **gb2312** "gb2312"
* **gbk** "gbk"
* **gb18030** "gb18030"
* **big5** "big5"
* **euc-jp** "euc-jp"
* **shift-jis** "shift-jis"
