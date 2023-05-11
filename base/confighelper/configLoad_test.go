package confighelper

import (
	"base/base/util/encryptutil"
	"fmt"
	"testing"
)

func TestGetIniConfig(t *testing.T) {
	k := "vleyun202088880123456789"
	userkey, _ := encryptutil.DesEncrypt([]byte("test"), []byte(k))
	fmt.Println(userkey)
}
