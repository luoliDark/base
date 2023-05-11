package confighelper

import (
	"fmt"
	"testing"

	"github.com/luoliDark/base/util/encryptutil"
)

func TestGetIniConfig(t *testing.T) {
	k := "vleyun202088880123456789"
	userkey, _ := encryptutil.DesEncrypt([]byte("test"), []byte(k))
	fmt.Println(userkey)
}
