package byaccount

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/luoliDark/base/util/commutil"

	"github.com/luoliDark/base/loghelper"
	"github.com/luoliDark/base/redishelper/rediscache"
	uuid "github.com/satori/go.uuid"
)

const authKey = "custom_key_api_oauth_access_token_"

var getTokenLock sync.Mutex

type SSOAuthTokenConf struct {
	Appid     string `json:"appid"`
	Timestamp string `json:"timestamp"`
	Sign      string `json:"sign"`
	Token     string `json:"token"`
	EntCode   string `json:"entcode"`
	EntID     string `json:"entid"`
	RequestId string `json:"request_id"`
}

type ApiOAuthTokenResult struct {
	EntCode     string `json:"entcode"`
	AccessToken string `json:"accessToken"`
	RequestId   string `json:"request_id"`
	ExpiresIn   string `json:"expires_in"`
}

// CheckOAuthToken 检查Token方法
// 检查Token是否存在，是否有效，是否过期，是否被禁用等，
// 验证逻辑：从Redis中获取Token信息，验证Token是否存在，是否有效，是否过期，是否被禁用等。
// @param token 要检查的Token
// @return error 错误信息
func CheckOAuthToken(token string, entcode string) (map[string]string, error) {
	var err error
	// 参数基础校验
	tokenInfo := make(map[string]string)
	if token == "" {
		return tokenInfo, errors.New("缺少必要参数")
	}
	tokenInfoStr := rediscache.GetString(authKey + token)
	if len(tokenInfoStr) == 0 || tokenInfoStr == "" {
		return tokenInfo, errors.New("token不存在")
	}
	err = json.Unmarshal([]byte(tokenInfoStr), &tokenInfo)
	if err != nil {
		return tokenInfo, errors.New("token解析失败")
	}
	// 验证Token是否存在
	if tokenInfo["appid"] == "" {
		return tokenInfo, errors.New("token不存在")
	}
	return tokenInfo, nil
}

// GetAuthTokenV1 获取Token方法
// 使用HMAC-SHA256生成签名，验证签名的有效性，生成Token并返回给客户端，
// 签名逻辑：请求参数 + 应用密钥，然后使用SHA-256哈希算法生成签名。
// @param appid 应用ID
// @param timestamp 时间戳
// @param reqData 请求数据
// @param sign 签名
// @return Token 生成的Token
// @return error 错误信息
func GetAuthTokenV1(appid, timestamp, sign, random string) (rs ApiOAuthTokenResult, err error) {
	// 参数基础校验
	if appid == "" || timestamp == "" || sign == "" {
		return rs, errors.New("缺少必要参数")
	}
	// 获取配置中的DES密钥
	appinfo := rediscache.GetHashMap(0, 0, "sys_appinfo", appid)
	if len(appinfo) == 0 || appinfo["appid"] == "" {
		return rs, errors.New("应用不存在:" + appid)
	}
	if appinfo["secret"] == "" {
		return rs, errors.New("应用未设置完成，请联系管理员！")
	}
	desSecret := appinfo["secret"]
	// 验证签名有效性
	expectedSig := generateSignature(desSecret + appid + commutil.ToString(timestamp))
	if expectedSig != sign {
		return rs, errors.New("签名验证失败")
	}

	getTokenLock.Lock()
	// 生成Token
	token := uuid.NewV4().String()
	getTokenLock.Unlock()
	reqID := commutil.GetUUID()

	// 保存Token到Redis，设置过期时间为2小时
	loghelper.ByInfo("获取Token成功", fmt.Sprintf("请求参数：appid=%s,timestamp=%s,sign=%s,entid=%s,reqID=%s",
		appid, timestamp, sign, appinfo["entid"], reqID), "")

	tokenInfo := SSOAuthTokenConf{Token: token, Appid: appid, Timestamp: timestamp,
		Sign: sign, RequestId: reqID, EntCode: appinfo["entid"], EntID: appinfo["entid"]}
	a, err := json.Marshal(tokenInfo)
	if err != nil {
		return rs, errors.New("生成Token失败")
	}
	rediscache.SetStringExpire(authKey+token, string(a), 7200)

	rs = ApiOAuthTokenResult{
		EntCode:     appinfo["entid"],
		AccessToken: token,
		RequestId:   reqID,
		ExpiresIn:   "7200",
	}
	return rs, nil
}

// 生成签名方法
// 使用HMAC-SHA256生成签名，验证签名的有效性，生成Token并返回给客户端，
// 签名逻辑：请求参数 + 应用密钥，然后使用SHA-256哈希算法生成签名。
func generateSignature(nonce string) string {
	hash := sha256.Sum256([]byte(nonce))
	return hex.EncodeToString(hash[:])
}
