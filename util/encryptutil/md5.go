//panl 2020.2.20 用于进行MD5加密

package encryptutil

import "github.com/wumansgy/goEncrypt"

//MDF5 256位加密
func EncryptSha256(str string) string {
	return goEncrypt.Sha256Hex([]byte(str))
}

//MDF5 512位加密
func DecryptSha512(str string) string {
	return goEncrypt.Sha512Hex([]byte(str))
}
