package sso

import (
	"encoding/json"
	"errors"
	"fmt"
	"runtime/debug"
	"strings"
	"time"

	"github.com/luoliDark/base/confighelper"
	"github.com/luoliDark/base/loghelper"
	"github.com/luoliDark/base/redishelper"
	"github.com/luoliDark/base/redishelper/rediscache"
	"github.com/luoliDark/base/sso/ssologin/byaccount"
	"github.com/luoliDark/base/sso/ssologin/common"
	"github.com/luoliDark/base/sysmodel"
	"github.com/luoliDark/base/sysmodel/restentity"
	"github.com/luoliDark/base/util/commutil"
	"github.com/luoliDark/base/util/encryptutil"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
)

var UserCheckErrorBySessionEmpty = errors.New("获取用户Session为空,请重新登录")

// 接口事件调用时，根据eventPar 获取用户,
// 用户token失效 改为从db重新获取用户信息。
func GetUserByEventPar(eventPar *restentity.EventPar) (sysmodel.SSOUser, error) {
	token := eventPar.Sid
	user, err := GetUserByToken(token)
	if err == nil {
		//接口事件类 取传过来的entid
		if !g.IsEmpty(eventPar.EntId) {
			user.EntID = eventPar.EntId
		}
	}
	return user, err
}

// 根据token获取用户信息
func GetUserByToken(token string) (sysmodel.SSOUser, error) {

	userId := redishelper.GetString(confighelper.GetEnterpriseID(), confighelper.GetSessionDbIndex(), token)

	if g.IsEmpty(userId) {

		sidArr := strings.Split(token, ":")
		if len(sidArr) > 2 {

			loginDate := sidArr[2] //取登录日期

			//如果是当天登录的过，说明没有过期， 就再重新从redis获取3次
			if loginDate == commutil.GetNowYYDDMM() {
				b := debug.Stack()
				stack := string(b)

				for i := 1; i <= 2; i++ {

					time.Sleep(time.Duration(1) * time.Second) //延时1秒

					//延时后重新获取
					userId = redishelper.GetString(confighelper.GetEnterpriseID(), confighelper.GetSessionDbIndex(), token)

					go common.InsertLoginLog(userId, "easyfa服务GetUserByToken执行失败，userId="+userId+" Sid="+token+
						"当天登录信息获取失败，偿试第"+commutil.ToString(i)+"次重获redis信息"+
						" 程序stack="+stack, token, "")

					if !g.IsEmpty(userId) {
						break //获取成功后不再遍历
					}

				}

			}
		}
	}

	if userId == "" {
		return sysmodel.SSOUser{}, UserCheckErrorBySessionEmpty
	}

	user, err := GetUserFormUserId(userId)
	if err != nil {
		go common.InsertLoginLog(userId, "easyfa服务GetUserByToken 执行失败 userId="+userId+" Sid="+token+"未从redis获取到用户对象", token, "")
		return user, err
	} else {
		/// 取消Token不一致 检查，以支持多端登录 ， 用户信息的SID设置为传的Token ，例如：后续接口调用时才能保证SID为同一个。
		user.SId = token
		return user, nil //登录成功
	}
}

// 根据userid 获取user对象
func GetUserFormUserId(UserId string) (sysmodel.SSOUser, error) {

	userjson := redishelper.GetString(confighelper.GetEnterpriseID(), confighelper.GetSessionDbIndex(), UserId)
	if userjson != "" {
		userbyte := []byte(userjson)
		newUser := sysmodel.SSOUser{}
		err := json.Unmarshal(userbyte, &newUser)
		if err != nil {
			loghelper.ByError("获取用户信息失败", commutil.AppendStr("解析用户JSON失败:",
				"userjson:", userjson, ",err:", err.Error()), UserId)
			return sysmodel.SSOUser{}, errors.New("获取用户信息失败！")
		}
		rolelen := len(newUser.UserRoleIds)
		if rolelen == 0 {
			return newUser, errors.New("获取用户角色失败")
		}
		//检查成功
		newUser.UserID = newUser.LoginUid

		if g.IsEmpty(newUser.EntID) {
			return newUser, errors.New("企业ID获取失败")
		}

		return newUser, nil
	} else {
		return sysmodel.SSOUser{}, UserCheckErrorBySessionEmpty
	}
}

// 根据http获取用户对象
func GetUserFromHttp(Ctx *gin.Context) (sysmodel.SSOUser, error) {

	var user sysmodel.SSOUser
	var err error
	hasUserInfo := false
	sidCookie := Ctx.Request.Header.Get("authorization")
	if sidCookie == "" {
		// 历史参数。这里做兼容二次获取
		sidCookie = Ctx.Request.Header.Get("set-cookie")
	}
	var httpSid string
	if !g.IsEmpty(sidCookie) {
		arr := strings.Split(sidCookie, "=")
		if arr != nil && len(arr) > 0 {
			httpSid = arr[1]
			user, err = GetUserByToken(httpSid)
			hasUserInfo = true
		}
	} else {
		//再次通过SID 获取 一般针对于自定义事件内部服务器直接传SID
		httpSid = Ctx.PostForm("sid")
		if !g.IsEmpty(httpSid) {
			user, err = GetUserByToken(httpSid)
			hasUserInfo = true
		} else {
			err = errors.New("head获取cookie,请重新登录")
		}

	}

	if err != nil {
		if &user != nil && hasUserInfo {
			b := debug.Stack()
			stack := string(b)
			go common.InsertLoginLog(user.UserID, "easyfa服务GetUserFromHttp执行失败，userid="+user.UserID+" 程序stack="+stack+"http参数set-cookie="+sidCookie+"httpSid="+httpSid+"用户对象Sid="+user.SId+"未从redis获取到用户对象", user.SId, "")
		}
		return user, err
	} else {

		//如果是接口事件，因为系统共用一个接口用户，所以需要再次一次实际entid
		if strings.Contains(user.UserName, "接口账号") {
			entid := Ctx.PostForm("entid")
			if !g.IsEmpty(entid) && entid != "null" && entid != "undefined" {
				user.EntID = entid
			}
		}

		if g.IsEmpty(user.EntID) {
			return sysmodel.SSOUser{}, errors.New("未选择企业,请重新登录")
		}
		return user, nil
	}
}

func IsAdmin(user *sysmodel.SSOUser) bool {
	roleIds := user.UserRoleIds
	if roleIds == nil {
		return false
	}

	//如果是管理员则返回true
	if _, isok := roleIds["2"]; isok {
		return true
	}
	//超级管理员
	if _, isok := roleIds["3"]; isok {
		return true
	}

	return false
}

// 检查APPID ，如果有效就生成虚拟用户
func GetVirtualUserByAppId(ctx *gin.Context) (sysmodel.SSOUser, error) {

	err, appid, _, entid := byaccount.ChkExAppid(ctx)
	if err != nil {
		return sysmodel.SSOUser{}, err
	} else {
		user, err := GetUserFormUserId(appid)
		if err != nil {
			user = sysmodel.SSOUser{
				UserCode:  appid,
				IsEnc:     false,
				UserID:    appid,
				LoginUid:  appid,
				UserName:  "接口账号",
				EntID:     entid,
				FormEntId: entid,
				LoginTime: commutil.GetNowTime(),
				SId:       "interface_" + appid,
				AppId:     appid,
			}
			user.UserRoleIds = make(map[string]string)
			user.UserRoleIds["1"] = "普通用户"
			_, err = byaccount.GetUsersEnt(&user) //绑定企业属性 例;IsCostCenter
			if err != nil {
				return user, err
			}
			byaccount.SetUserSession(&user)
		}

		if user.EnList == nil || len(user.EnList) == 0 {
			// 做补救登录处理
			_, err = byaccount.GetUsersEnt(&user) //绑定企业属性 例;IsCostCenter
			if err != nil {
				return user, err
			}
			byaccount.SetUserSession(&user)
		}
		return user, nil
	}
}

// 检查APPID ，如果有效就生成虚拟用户
func GetVirtualUserByConfigIni(ctx *gin.Context) (sysmodel.SSOUser, error) {

	err, appid, secret := getAppId(ctx) //获取前台参数
	if err != nil {
		return sysmodel.SSOUser{}, err
	} else {

		request_key := encryptutil.StringToBase64(appid)
		request_secret := encryptutil.StringToBase64(secret)
		config_key := confighelper.GetIniConfig("exocr", "fcfk")
		config_secret := confighelper.GetIniConfig("exocr", "fcfs")

		if config_key == request_key && config_secret == request_secret {

			user := sysmodel.SSOUser{
				UserCode:  appid,
				IsEnc:     false,
				UserID:    appid,
				LoginUid:  appid,
				UserName:  "接口账号",
				EntID:     "1",
				FormEntId: "1",
				LoginTime: commutil.GetNowTime(),
				SId:       "interface_" + appid,
				AppId:     appid,
			}

			user.UserRoleIds = make(map[string]string)
			user.UserRoleIds["1"] = "普通用户"

			return user, nil

		}

		return sysmodel.SSOUser{}, errors.New("生成接口用户失败")
	}
}

func getAppId(ctx *gin.Context) (err error, appid, secret string) {

	appid = ctx.PostForm("appid")
	secret = ctx.PostForm("secret")

	//timestamp := ctx.PostForm("timestamp") //时间戳检查
	//sign := ctx.PostForm("sign")           //参数签名

	if g.IsEmpty(appid) {
		appid = ctx.Query("appid")
	}
	if g.IsEmpty(secret) {
		secret = ctx.Query("secret")
	}
	//从请求头获取参数
	if appid == "" {
		appid = ctx.Request.Header.Get("appid")
	}
	if secret == "" {
		secret = ctx.Request.Header.Get("secret")
	}

	if g.IsEmpty(appid) {
		return errors.New("未获取到appid"), "", ""
	} else {
		return nil, appid, secret
	}
}

// tokenID 和 文件 做关联处理
func GetFileOAuthTokenByAttach(tokenID, fid string) string {
	return tokenID //暂时直接返回tokenID，待优化
}

func GetFileOAuthTokenByUserID(userID string) string {
	if userID == "" {
		return ""
	}
	tokenID := ""
	userKey := CustomFileTokenUseridKey + userID
	authKey := ""
	val := rediscache.GetString(userKey)
	if val != "" {
		valList := strings.Split(val, ":")
		if len(valList) > 1 {
			if valList[1] == commutil.GetNowYYDDMM() {
				tokenID = val //相同一天,不更新
			}
		} else {
			tokenID = val
		}
	}
	if tokenID != "" {
		return tokenID
	}
	tokenID = commutil.GetUUID() + ":" + commutil.GetNowYYDDMM()
	authKey = CustomFileTokenOauthidKey + tokenID
	rediscache.SetStringExpire(userKey, tokenID, 60*60*24) //1天过期
	rediscache.SetStringExpire(authKey, userID, 60*60*24)  //1天过期
	return tokenID
}

func GetFileOAuthTokenByUser(user *sysmodel.SSOUser) string {
	return GetFileOAuthTokenByUserID(user.UserID)
}

// ValidateFileOAuthToken 文件访问token验证
// tokenID:文件访问token(用户token)
// ts:时间戳
func ValidateFileOAuthToken(tokenID string, ts string) (err error) {
	if tokenID == "" {
		return fmt.Errorf("未授权访问")
	}
	authKey := CustomFileTokenOauthidKey + tokenID
	val := rediscache.GetString(authKey)
	if val != "" {
		return
	}
	return fmt.Errorf("无效token")
}
