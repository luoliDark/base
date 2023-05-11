package pinyinutil

import "github.com/Lofanmi/pinyin-golang/pinyin"

// 返回中文的首字母 sep=为间隔符
func PinYinConvertAbbr(str, sep string) string {
	dict := pinyin.Dict{}
	return dict.Abbr(str, sep)
}

func PinYinCovertFull(str, sep string) string {
	dict := pinyin.Dict{}
	return dict.Name(str, sep).None()
}
