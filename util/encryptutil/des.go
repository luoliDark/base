//panl 2020.2.20 用于进行DES 对称加密

package encryptutil

import (
	"encoding/base64"
	"fmt"
	"github.com/wumansgy/goEncrypt"
	"strings"
)

//DES加密
func DesEncrypt(plainText, key []byte) (string, error) {

	cryptText, err := goEncrypt.TripleDesEncrypt(plainText, key)
	if err != nil {
		return "", err
	}

	str := base64.StdEncoding.EncodeToString(cryptText)
	str = strings.ReplaceAll(str, "+", ".zzsoft888.") //+号
	return str, nil
}

//DES解密
func DesDecrypt(cipherText string, key []byte) (string, error) {

	cipherText = strings.ReplaceAll(cipherText, ".zzsoft888.", "+") //转回+号
	word, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	newplaintext, err := goEncrypt.TripleDesDecrypt(word, key)
	if err != nil {
		return "", err
	}

	return string(newplaintext), nil
}

//对数据进行脱敏显示
func TuoMing(str string) string {

	str = strings.ReplaceAll(str, " ", "")

	var result string
	var sb strings.Builder
	index := 1
	for _, r := range str {
		index++
		s := fmt.Sprintf("%c", r)
		if index%3 == 0 {
			sb.WriteString(s)
		} else {
			sb.WriteString("*")
		}
	}
	result = sb.String()
	return result
}
