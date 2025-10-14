package encryptutil

import (
	"bytes"
	"crypto/cipher"
	"fmt"

	"github.com/tjfoc/gmsm/sm4"
)

//明文数据填充
func paddingLastGroup(plainText []byte, blockSize int) []byte {
	//1.计算最后一个分组中明文后需要填充的字节数
	padNum := blockSize - len(plainText)%blockSize
	//2.将字节数转换为byte类型
	char := []byte{byte(padNum)}
	//3.创建切片并初始化
	newPlain := bytes.Repeat(char, padNum)
	//4.将填充数据追加到原始数据后
	newText := append(plainText, newPlain...)

	return newText
}

//去掉明文后面的填充数据
func unpaddingLastGroup(plainText []byte) []byte {
	//1.拿到切片中的最后一个字节
	length := len(plainText)
	lastChar := plainText[length-1]
	//2.将最后一个数据转换为整数
	number := int(lastChar)
	return plainText[:length-number]
}

//加密
func sm4Encrypt(plainText, key []byte) []byte {
	block, err := sm4.NewCipher(key)
	if err != nil {
		panic(err)
	}
	paddData := paddingLastGroup(plainText, block.BlockSize())
	iv := []byte("12345678qwertyui")
	blokMode := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(paddData))
	blokMode.CryptBlocks(cipherText, paddData)
	return cipherText
}

//解密
func sm4Dectypt(cipherText, key []byte) []byte {
	block, err := sm4.NewCipher(key)
	if err != nil {
		panic(err)
	}
	iv := []byte("12345678qwertyui")
	blockMode := cipher.NewCBCDecrypter(block, iv)
	blockMode.CryptBlocks(cipherText, cipherText)
	plainText := unpaddingLastGroup(cipherText)
	return plainText
}

func main() {
	src := []byte("这是对称加密SM4的CBC模式加解密测试")
	key := []byte("1q2w3e4r5t6y7u8i")
	//加密
	cipherText := sm4Encrypt(src, key)
	fmt.Println(string(cipherText))
	// 解密
	plainText := sm4Dectypt(cipherText, key)
	fmt.Println(string(plainText))
	flag := bytes.Equal(src, plainText)
	fmt.Println("解密是否成功：", flag)
}
