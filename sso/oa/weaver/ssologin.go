package weaver

import (
	"base/db/conn"
	"base/sso/ssologin/byaccount"
	"base/sysmodel"
	"base/sysmodel/eb"
	"errors"
)

//des对称解密 登录
func WeaverLogin(luser *sysmodel.SSOUser, usercode string) (*sysmodel.SSOUser, error) {
	engine, _ := conn.GetDB()
	user := new(eb.Eb_user)

	has, err := engine.Where("usercode=?", usercode).Get(user)
	if err != nil {
		return &sysmodel.SSOUser{}, err
	}
	if !has {
		//OA 跳转过来，loginid 不是hr系统的工号，usercode要使用HR的工号，所以OA的登录账号放到 拓展loginid,
		ul := []eb.Eb_user{}
		err = engine.Where("expandloginid=? and ifnull(isdiscard,0) = 0", usercode).Find(&ul)
		if err != nil {
			return &sysmodel.SSOUser{}, err
		}
		if len(ul) == 0 {
			// 没有账户
			return &sysmodel.SSOUser{}, errors.New("账号不存在")
		}
		if len(ul) > 1 {
			// 存在多个相同的 loginid
			return &sysmodel.SSOUser{}, errors.New("loginid 账号异常，账号编码重复。请联系管理员！")
		}
		user = &ul[0]
	}
	//绑定用户信息
	_, err = byaccount.BindSSOUser(user, luser, luser.IsMobile)
	if err != nil {
		return luser, err
	}
	//写入redis
	if !byaccount.SetUserSession(luser) {
		return luser, errors.New("登录成功，设置redis时失败")
	}

	return luser, nil
}
