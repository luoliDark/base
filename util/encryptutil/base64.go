package encryptutil

import (
	"encoding/base64"
	"fmt"
)

//转为base64
func StringToBase64(str string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(str))
	return encoded
}

//base64还原
func Base64ToString(base64Str string) string {
	decoded, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return "解析base64失败"
	}
	return string(decoded)
}
