package encryptutil

import (
	"fmt"
	"testing"

	"github.com/luoliDark/base/confighelper"
	"github.com/wumansgy/goEncrypt"
)

func TestDesEncrypt(t *testing.T) {
	//plaintext := []byte("床前明月光，疑是地上霜，举头望明月，学习go语言")
	//fmt.Println("明文为：", string(plaintext))
	//
	//// 传入明文和自己定义的密钥，密钥为24字节 可以自己传入初始化向量,如果不传就使用默认的初始化向量,8字节
	//cryptText, err := goEncrypt.TripleDesEncrypt(plaintext, []byte("wumansgy12345678asdfghjk"))
	//if err != nil {
	//	fmt.Println("111", err)
	//}
	//fmt.Println("三重DES的CBC模式加密后的密文为:", base64.StdEncoding.EncodeToString(cryptText))
	//
	//a:=base64.StdEncoding.EncodeToString(cryptText)
	//b,_:=base64.StdEncoding.DecodeString(a)
	//// 传入密文和自己定义的密钥，需要和加密的密钥一样，不一样会报错 可以自己传入初始化向量,如果不传就使用默认的初始化向量,8字节
	//newplaintext, err := goEncrypt.TripleDesDecrypt(b, []byte("wumansgy12345678asdfghjk"))
	//if err != nil {
	//	fmt.Println("222", err)
	//}
	//
	//fmt.Println("三重DES的CBC模式解密完：", string(newplaintext))

	k := confighelper.GetDesKey()
	//fmt.Print(k)
	ucode := "964"
	usk, _ := DesEncrypt([]byte(ucode), []byte(k))
	fmt.Println("工号：" + ucode + "--》生成的密文为：" + usk)

	usercode, _ := DesDecrypt(usk, []byte(k))
	fmt.Println("解析出来工号=", usercode)

	//
	//a, _ := DesDecrypt("b2huCnge04s=", []byte(k))
	//fmt.Println(a)

	//s := TuoMing("要 要有一65在顶替寺压至hjj7")
	//fmt.Println(s)
}

func TestDesEncrypt2(t *testing.T) {
	hash := goEncrypt.Sha256Hex([]byte("test"))
	fmt.Println(hash)
}

func TestDesDecrypt(t *testing.T) {
	//i7dSTVJrfSs=	hNdGDlhEJ68=,hNdGDlhEJ68=
	//usk=i7dSTVJrfSs=\u0026usk=0fpZDWdaf/s=\u0026usk=xpwn9QJ+1RE=\u0026usk=oEI10uvMe0s=\u0026usk=lyJTownInXU=\u0026usk=Z6S7TfwSq5Q=\u0026usk=TdsdDijwmRQ=\u0026usk=O0j90tOGTF8=\u0026usk=O2kN7tvPsNs=\u0026usk=i7dSTVJrfSs=
	k := "vleyun202088880123456789"
	fmt.Println(DesDecrypt("hNdGDlhEJ68=", []byte(k)))
	fmt.Println(DesDecrypt("qFRCU/siyT3dlLwhcccx0Q==", []byte(k)))
}
