package sso

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
	"paas/base/confighelper"
	"paas/base/db/dbhelper"
	"paas/base/loghelper"
	"paas/base/redishelper"
	"paas/base/sso/ssologin/byaccount"
	"paas/base/sso/ssologin/common"
	"paas/base/sysmodel"
	"paas/base/sysmodel/logtype"
	"paas/base/util/commutil"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

// 根据token获取用户信息
func GetUserByToken(token string) (sysmodel.SSOUser, error) {

	userId := redishelper.GetString(confighelper.GetEnterpriseID(), confighelper.GetSessionDbIndex(), token)

	if g.IsEmpty(userId) {
		sidArr := strings.Split(token, ":")
		if len(sidArr) > 2 {
			loginDate := sidArr[2] //取登录日期

			//如果是当天登录的过，说明没有过期， 就再重新从redis获取2次
			if loginDate == commutil.GetNowYYDDMM() {
				b := debug.Stack()
				stack := string(b)

				for i := 1; i <= 2; i++ {

					time.Sleep(time.Duration(1) * time.Second) //延时1秒

					//延时后重新获取
					userId = redishelper.GetString(confighelper.GetEnterpriseID(), confighelper.GetSessionDbIndex(), token)

					go common.InsertLoginLog(userId, "bpmserver服务GetUserByToken执行失败，userId="+userId+" Sid="+token+
						"当天登录信息获取失败，偿试第"+commutil.ToString(i)+"次重获redis信息"+
						" 程序stack="+stack, token, "")

					if !g.IsEmpty(userId) {
						break //获取成功后不再遍历
					}

				}

			}
		}
	}

	if g.IsEmpty(userId) {
		return sysmodel.SSOUser{}, errors.New("未取到缓存中用户ID信息")
	}

	user, err := GetUserFormUserId(userId)
	if err != nil {

		b := debug.Stack()
		stack := string(b)

		go common.InsertLoginLog(userId, "bpmserver服务GetUserByToken执行失败，userId="+userId+" Sid="+token+"未从redis获取到用户对象"+" 程序stack="+stack, token, "")
		return user, err
	} else {
		/// 取消Token不一致 检查，以支持多端登录 ， 用户信息的SID设置为传的Token ，例如：后续接口调用时才能保证SID为同一个。
		user.SId = token
		return user, nil //登录成功
	}
}

//根据userid 获取user对象
func GetUserFormUserId(UserId string) (sysmodel.SSOUser, error) {
	userjson := redishelper.GetString(confighelper.GetEnterpriseID(), confighelper.GetSessionDbIndex(), UserId)
	if userjson != "" {
		userbyte := []byte(userjson)
		newUser := sysmodel.SSOUser{}
		err := json.Unmarshal(userbyte, &newUser)
		if err != nil {
			return sysmodel.SSOUser{}, errors.New("获取用户Session为空")
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
		return sysmodel.SSOUser{}, errors.New("获取用户信息为空,请重新登录")
	}
}

//根据http获取用户对象
func GetUserFromHttp(Ctx *gin.Context) (sysmodel.SSOUser, error) {

	var user sysmodel.SSOUser
	var err error
	sidCookie := Ctx.Request.Header.Get("authorization")
	if sidCookie == "" {
		// 历史参数。做兼容二次获取
		sidCookie = Ctx.Request.Header.Get("set-cookie")
	}
	var httpSid string
	if !g.IsEmpty(sidCookie) {
		arr := strings.Split(sidCookie, "=")
		if arr != nil && len(arr) > 0 {
			httpSid = arr[1]
			user, err = GetUserByToken(httpSid)
		}
	} else {

		httpSid = Ctx.Request.Header.Get("sid")

		if g.IsEmpty(httpSid) {
			//再次通过SID 获取 一般针对于自定义事件内部服务器直接传SID
			httpSid = Ctx.PostForm("sid")
		}

		if !g.IsEmpty(httpSid) {
			user, err = GetUserByToken(httpSid)
		} else {
			err = errors.New("head获取cookie,请重新登录")
		}
	}

	if err != nil {

		if &user != nil {
			b := debug.Stack()
			stack := string(b)
			go common.InsertLoginLog(user.UserID, "bpmserver服务GetUserFromHttp执行失败，userid="+user.UserID+" 程序stack="+stack+"http参数set-cookie="+sidCookie+"httpSid="+httpSid+"用户对象Sid="+user.SId+"未从redis获取到用户对象", user.SId, "")
		}

		return user, err
	} else {
		if g.IsEmpty(user.EntID) {

			b := debug.Stack()
			stack := string(b)

			go common.InsertLoginLog(user.UserID, "bpmserver服务GetUserFromHttp执行失败，userid="+user.UserID+" 程序stack="+stack+"http参数set-cookie="+sidCookie+"httpSid="+httpSid+"用户对象Sid="+user.SId+"未从redis获取到用户对象", user.SId, "")

			return sysmodel.SSOUser{}, errors.New("未选择企业,请重新登录")
		} else {
			//记录用户在线数
			//	go SaveOnlinecnt(user)
		}
		return user, nil
	}
}

var mutex sync.RWMutex

//记录用户在线数(只保留一个月的数据)
func SaveOnlinecnt(user sysmodel.SSOUser) {
	mutex.Lock()
	defer mutex.Unlock()
	day := time.Now().Day()
	hour := time.Now().Hour()
	userid := user.UserID
	entid := user.EntID
	//每个月1号，删除上个月的用户在线记录
	if day == 1 {
		sql := "delete from sys_onlinecnt where date_format(InsertDate,'%Y-%m') = ?"
		_, err := dbhelper.ExecSql(userid, true, sql, commutil.TimeFormat(time.Now().AddDate(0, -1, 0), "2006-01"))
		if err != nil {
			loghelper.ByError(logtype.ExecSqlErr, err.Error()+" 删除上个月的用户在线记录失败SQL:"+sql, user.UserID)
		}
	}
	//验证该用户在当前时间段内是否登录过
	sql := "select count(1) from sys_onlinecnt where day = ? and time = ? and userid = ? and entid = ?"
	count, err := dbhelper.QueryFirstCol(userid, true, sql, day, hour, userid, entid)
	if err != nil {
		loghelper.ByError(logtype.QueryErr, err.Error()+" 查询用户在线记录sql:"+sql, user.UserID)
	}
	if commutil.ToInt(count) == 0 {
		//在当前时间段未登录则记录在线用户信息
		sql = "insert into sys_onlinecnt(OnlinecntId,EntId,UserId,Day,Time,InsertDate) values(?,?,?,?,?,now())"
		_, err = dbhelper.ExecSql(user.UserID, true, sql, commutil.GetUUID(), entid, userid, day, hour)
		if err != nil {
			loghelper.ByError(logtype.ExecSqlErr, err.Error()+" 记录用户在线数sql:"+sql, user.UserID)
		}
	}
}

//检查APPID ，如果有效就生成虚拟用户
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
			_, _ = byaccount.GetUsersEnt(&user) //绑定企业属性 例;IsCostCenter
			if err != nil {
				return user, err
			}
			byaccount.SetUserSession(&user)
		}

		if user.EnList == nil || len(user.EnList) == 0 {
			// 做补救登录处理
			_, _ = byaccount.GetUsersEnt(&user) //绑定企业属性 例;IsCostCenter
			if err != nil {
				return user, err
			}
			byaccount.SetUserSession(&user)
		}

		return user, nil
	}
}
