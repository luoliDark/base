package encryptutil

import (
	"crypto/rand"
	"fmt"

	"github.com/tjfoc/gmsm/sm2"
)

func Aaaa() {

	reader := rand.Reader
	reader.Read([]byte("aaaaa"))
	//生成私钥
	privateKey, e := sm2.GenerateKey(reader)
	if e != nil {
		fmt.Println("sm2 encrypt faild！")
	}

	//从私钥中获取公钥
	//pubkey := &privateKey.PublicKey

	aaaa, _ := privateKey.Sign(reader, []byte("bbbbb"), nil)
	fmt.Println(string(aaaa))

	//verify := pubkey.Verify([]byte(aaaa), nil)
	//fmt.Println(string(verify))
	//msg:=  []byte("i am   wek &&  i am The_Reader too 。")
	////用公钥加密msg
	//bytes, i := pubkey.Encrypt(msg)
	//
	//if i !=nil{
	//	fmt.Println("使用私钥加密失败！")
	//}
	//
	//fmt.Println("the encrypt msg  =  ",hex.EncodeToString(bytes))
	////用私钥解密msg
	//decrypt, i2 := privateKey.Decrypt(bytes)
	//
	//if i2 != nil{
	//
	//	fmt.Println("使用私钥解密失败！")
	//}
	//
	//fmt.Println( "the msg  = ", string(decrypt))

}
