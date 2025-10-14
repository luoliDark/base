package byaccount

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/luoliDark/base/redishelper/rediscache"
	"github.com/luoliDark/base/util/commutil"
)

func TestGetSSOAuthTokenV1(t *testing.T) {
	tokenInfo := SSOAuthTokenConf{Token: "1", Appid: "2", Timestamp: "3", ReqData: "4", Sign: "5"}
	a, err := json.Marshal(tokenInfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	rediscache.SetStringExpire(authKey+"1123123442323232ada", string(a), 7200)
}

func TestGetSSOAuthTokenV12(t *testing.T) {
	m, err := CheckOAuthToken("1123123442323232ada", "")
	fmt.Println(m, err)
}

func TestGetAuthTokenV1(t *testing.T) {
	timestamp := time.Now().UnixMilli()
	appid := "oa"
	desSecret := "6AwHVHYsjq2D1vnw3vDk1VWMyfadfdewrgjuvxa"
	sign := EncodeToStringBySHA256(desSecret + appid + commutil.ToString(timestamp))

	p1, p2 := GetAuthTokenV1("oa", commutil.ToString(timestamp), sign, "1")
	fmt.Println(p1, p2)
}

// EncodeToString 使用SHA - 256对输入的字符串进行哈希计算，并返回十六进制字符串
func EncodeToStringBySHA256(data string) string {
	// 将输入的字符串转换为字节切片
	bytes := []byte(data)
	// 创建一个新的SHA - 256哈希对象
	h := sha256.New()
	// 向哈希对象写入字节数据
	h.Write(bytes)
	// 计算最终的哈希值
	t := h.Sum(nil)
	// 将哈希值转换为十六进制字符串
	return hex.EncodeToString(t)
}
