package byOnlyView

import (
	"base/base/confighelper"
	"base/base/db/conn"
	"base/base/redishelper"
	"base/base/sso/ssologin/byaccount"
	"base/base/sysmodel"
	"base/base/sysmodel/eb"
	"errors"
)

type OnlyViewUser struct {
	FormUrl string
	User    sysmodel.SSOUser
}

func GetFormUrlByOnlyView(key string) string {

	//通过redis查询出实际URL，并返回接口传过来的loginuid 用户，前台如果没登录的 就用k3用户写cookie中

	formUrl := redishelper.GetString(confighelper.GetEnterpriseID(), confighelper.GetSessionDbIndex(), key)
	return formUrl

}

func GetUserByOnlyView(loginUid string) (sysmodel.SSOUser, error) {

	session := conn.GetSession(true)
	defer session.Close()

	ebuser := eb.Eb_user{}
	luser := sysmodel.SSOUser{IsMobile: 0}
	isok, err := session.Where("usercode=? or CWSoftInnerUid=? and (IsDiscard=0 or isdiscard is null )", loginUid, loginUid).Get(&ebuser)
	if isok {
		_, err = byaccount.BindSSOUser(&ebuser, &luser, luser.IsMobile)
		if err != nil {
			return luser, err
		}

		//写入redis
		if !byaccount.SetUserSession(&luser) {
			return luser, errors.New("登录成功，设置redis时失败")
		}

	}

	return luser, nil

}
