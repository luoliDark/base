package commutil

import (
	"regexp"
	"strings"
	"unicode"
)

/**
 * @author : YiXin
 * @date 2022-05-12
 * @version V1.0
 **/

const SpecialChar = `[\t\n\r]`

// 替换前端显示的特殊字符
func ReplaceWebShowSpecialChar(val string) string {
	if val == "" {
		return ""
	}
	reg := regexp.MustCompile(SpecialChar)
	ismatch := reg.MatchString(val)
	if !ismatch {
		return val
	}
	if strings.Contains(val, "\t") {
		val = strings.ReplaceAll(val, "\t", "  ")
	}
	if strings.Contains(val, "\n") {
		val = strings.ReplaceAll(val, "\n", "\\n")
	}
	if strings.Contains(val, "\r") {
		val = strings.ReplaceAll(val, "\r", " ")
	}
	return val
}

func ReplaceSpecialChar(val, sep string) string {
	if val == "" {
		return ""
	}
	reg := regexp.MustCompile(SpecialChar)
	ismatch := reg.MatchString(val)
	if !ismatch {
		return val
	}
	s := reg.ReplaceAllString(val, sep)
	return s
}

// 特殊字符替换为空
func ReplaceSpecialCharTheEmpty(val string) string {
	return ReplaceSpecialChar(val, "")
}

//是否有特殊 字符, 汉字等
func HasSpecialCharacterByStr(val string) bool {
	for _, v := range val {
		if HasSpecialCharacter(v) {
			return true
		}
	}
	return false
}

/**
判断是否为字母： unicode.IsLetter(v)
判断是否为十进制数字： unicode.IsDigit(v)
判断是否为数字： unicode.IsNumber(v)
判断是否为空白符号： unicode.IsSpace(v)

判断是否为Unicode标点字符 :unicode.IsPunct(v)
判断是否为中文：unicode.Han(v)
*/
func HasSpecialCharacter(letter rune) bool {
	if unicode.IsPunct(letter) || unicode.IsSymbol(letter) ||
		unicode.Is(unicode.Han, letter) {
		return true
	}
	return false
}

// Web参数替换空值
func WebParamsReplaceNull(s string) string {
	if s == "undefined" || s == "null" {
		return ""
	}
	return s
}

// WebParamsIsNull
func WebParamsIsNotNull(s string) bool {
	return !WebParamsIsNull(s)
}

// Web请求参数是否空值
func WebParamsIsNull(s string) bool {
	if s == "undefined" || s == "null" || s == "" {
		return true
	}
	return false
}
