chardet
=======

detect text encoding, like python chardet, but for go

本包用于探测文本的编码格式。

支持探测的格式有：
5种unicode格式：UTF8、大小端UTF16，大小端UTF32；
4种中文格式：hz-gb2312、gbk、gb18030、big5；
3种日文格式euc-jp、shift-jis、iso-2022-jp；
1种日文格式euc-kr；
不支持其他语种和编码格式的探测。

本包只有一个函数：

    func Check([]byte) string

该函数用于探测格式。如果无法探测到编码格式，会返回空字符串。
如果发现该文本符合多个编码格式，会优先返回utf-8格式；
否则进一步检测字符分布，返回最匹配的。

各编码格式对应字符串如下

* **UTF8**  "utf-8"
* **UTF16 Big-Ending** "utf-16BE"
* **UTF16 Little-Ending** "utf-16LE"
* **UTF32 Big-Ending** "utf-32BE"
* **UTF32 Little-Ending** "utf-32LE"
* **HZ-GB2312** "hz-gb2312"
* **GBK** "gbk"
* **GB18030** "gb18030"
* **BIG5** "big5"
* **Euc-Kr** "euc-kr"
* **Euc-Jp** "euc-jp"
* **Shift-Jis** "shift-jis"
* **ISO-2022-JP** "iso-2022-jp"
