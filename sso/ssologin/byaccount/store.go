package byaccount

import (
	"encoding/json"
	"fmt"

	"github.com/gogf/gf/frame/g"
	"github.com/luoliDark/base/confighelper"
	"github.com/luoliDark/base/redishelper"
	"github.com/luoliDark/base/sso/ssologin/common"
	"github.com/luoliDark/base/sysmodel"
	"github.com/luoliDark/base/util/commutil"
)

// 设置用户缓存， redis 以及cookie
func SetUserSession(user *sysmodel.SSOUser) bool {
	return SetUserSessionBySid(user, "")
}

// 设置用户缓存， redis 以及cookie
// 指定SID设置
func SetUserSessionBySid(user *sysmodel.SSOUser, sid string) bool {

	if user.UserRoleIds != nil && len(user.UserRoleIds) > 0 {
		//此角色用户做只能加密查看
		_, ok := user.UserRoleIds["7"]
		if ok {
			user.IsEnc = true
		}
	}

	//  设置SID
	if g.IsEmpty(sid) {
		//生成token
		if g.IsEmpty(user.AppId) {

			if user.SsoTime == -1 {
				//表示永久

				if commutil.HasSpecialCharacterByStr(user.UserCode) {
					user.SId = ":" + commutil.GetUUID()
				} else {
					//token中增加登录日期。 用于在检查登录时，若redis获取失败时，再根据登录日期判定是否为当天，是当天时 就重新延时从redis取值
					user.SId = user.UserCode + ":" + commutil.GetUUID()
				}

			} else {

				//表示只保存一天 所以拼上时间 用于判定是否是当天登录

				if commutil.HasSpecialCharacterByStr(user.UserCode) {
					user.SId = ":" + commutil.GetUUID() + ":" + commutil.GetNowYYDDMM()
				} else {
					//token中增加登录日期。 用于在检查登录时，若redis获取失败时，再根据登录日期判定是否为当天，是当天时 就重新延时从redis取值
					user.SId = user.UserCode + ":" + commutil.GetUUID() + ":" + commutil.GetNowYYDDMM()
				}

			}

		} else {
			//接口用户直接用APPID 以便长期有效
			user.SId = "interface_" + user.AppId
		}
	}

	// 取消删除 Token SID ，登录后生成新的 SID ，根据UserID，使用新用户对象覆盖原先的用户对象信息。  这样可以支持多端登录。
	user.LoginUid = user.UserID
	if user.IsMobile == 0 {
		// pc 端登录
		oldjson := redishelper.GetString(confighelper.GetEnterpriseID(), confighelper.GetSessionDbIndex(), user.UserID)

		//接口用户 不重复写入
		if !g.IsEmpty(oldjson) && !g.IsEmpty(user.AppId) {
			return true
		}

		if !g.IsEmpty(oldjson) && g.IsEmpty(user.AppId) {

			olduser := sysmodel.SSOUser{}
			oldbyte := []byte(oldjson)
			err := json.Unmarshal(oldbyte, &olduser)

			if err != nil {
				fmt.Println("用户反序列化失败")
			}
		}

		// 保存最新token
		jsonvalue := commutil.ObjectToJson(&user)

		//SsoTime = -1表示永久保存
		if g.IsEmpty(user.AppId) && user.SsoTime != -1 {
			//员工账号
			var timeout = 60 * 60 * 24 //一天后过期
			redishelper.SetStringExpire(confighelper.GetEnterpriseID(), confighelper.GetSessionDbIndex(), user.SId, user.UserID, timeout)
			redishelper.SetStringExpire(confighelper.GetEnterpriseID(), confighelper.GetSessionDbIndex(), user.UserID, jsonvalue, timeout)
		} else {

			//接口账号  或 永久账号
			redishelper.SetString(confighelper.GetEnterpriseID(), confighelper.GetSessionDbIndex(), user.SId, user.UserID)
			redishelper.SetString(confighelper.GetEnterpriseID(), confighelper.GetSessionDbIndex(), user.UserID, jsonvalue)

		}

		go common.InsertLoginLog(user.UserID, "PC登录成功", user.SId, "PC")

	} else {

		//移动端登录
		//在sid userid 的key前增加M_
		mUserid := "M_" + user.UserID

		oldjson := redishelper.GetString(confighelper.GetEnterpriseID(), confighelper.GetSessionDbIndex(), mUserid)
		if !g.IsEmpty(oldjson) {

			olduser := sysmodel.SSOUser{}
			oldbyte := []byte(oldjson)
			err := json.Unmarshal(oldbyte, &olduser)

			if err != nil {
				fmt.Println("用户反序列化失败")
			}
		}

		// 保存最新token
		jsonvalue := commutil.ObjectToJson(&user)

		if user.SsoTime != -1 {

			var timeout = 60 * 60 * 24 //一天后过期
			redishelper.SetStringExpire(confighelper.GetEnterpriseID(), confighelper.GetSessionDbIndex(), user.SId, mUserid, timeout)
			redishelper.SetStringExpire(confighelper.GetEnterpriseID(), confighelper.GetSessionDbIndex(), mUserid, jsonvalue, timeout)
		} else {

			//var timeout = 60 * 60 * 24 //一天后过期
			//redishelper.SetStringExpire(confighelper.GetEnterpriseID(), confighelper.GetSessionDbIndex(), user.SId, mUserid, timeout)
			//redishelper.SetStringExpire(confighelper.GetEnterpriseID(), confighelper.GetSessionDbIndex(), mUserid, jsonvalue, timeout)

			//长久保存
			redishelper.SetString(confighelper.GetEnterpriseID(), confighelper.GetSessionDbIndex(), user.SId, mUserid)
			redishelper.SetString(confighelper.GetEnterpriseID(), confighelper.GetSessionDbIndex(), mUserid, jsonvalue)

		}

		go common.InsertLoginLog(user.UserID, "手机登录成功", user.SId, "mobile")
	}

	//新增时默认部门（汪：redis存的是实际部门，前台VUE显示的是制单带出部门）
	if !g.IsEmpty(user.DefDeptIdByAdd) {
		user.DeptID = user.DefDeptIdByAdd
		user.DeptName = user.DefDeptNameByAdd
	}

	return true
}
