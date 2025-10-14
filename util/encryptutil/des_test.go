package encryptutil

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/wumansgy/goEncrypt"
)

func TestDesEncrypt(t *testing.T) {
	plaintext := []byte("床前明月光，疑是地上霜，举头望明月，学习go语言")
	fmt.Println("明文为：", string(plaintext))

	// 传入明文和自己定义的密钥，密钥为24字节 可以自己传入初始化向量,如果不传就使用默认的初始化向量,8字节
	cryptText, err := goEncrypt.TripleDesEncrypt(plaintext, []byte("wumansgy12345678asdfghjk"))
	if err != nil {
		fmt.Println("111", err)
	}
	fmt.Println("三重DES的CBC模式加密后的密文为:", base64.StdEncoding.EncodeToString(cryptText))

	// 传入密文和自己定义的密钥，需要和加密的密钥一样，不一样会报错 可以自己传入初始化向量,如果不传就使用默认的初始化向量,8字节
	newplaintext, err := goEncrypt.TripleDesDecrypt(cryptText, []byte("wumansgy12345678asdfghjk"))
	if err != nil {
		fmt.Println("222", err)
	}

	fmt.Println("三重DES的CBC模式解密完：", string(newplaintext))
}

func TestDesEncrypt2(t *testing.T) {
	hash := goEncrypt.Sha256Hex([]byte("test"))
	fmt.Println(hash)
}
