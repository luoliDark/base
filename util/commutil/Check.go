package commutil

import "regexp"

const (
	regular_mobile = `^1[3|4|5|6|7|8|9]\d{9}$`
)

// 匹配规则
// ^1第一位为一
// [345789]{1} 后接一位345789 的数字
// \\d \d的转义 表示数字 {9} 接9位
// $ 结束符
func CheckMobile(mobileNum string) bool {
	reg := regexp.MustCompile(regular_mobile)
	return reg.MatchString(mobileNum)
}
