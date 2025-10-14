package password

import (
	"fmt"

	"github.com/luoliDark/base/db/conn"
	"github.com/luoliDark/base/loghelper"
	"github.com/luoliDark/base/sysmodel"
	"github.com/luoliDark/base/sysmodel/eb"
	"github.com/luoliDark/base/util/encryptutil"
)

/**
 * @describe:
 *
 * @Author: YiXin
 * @Date: 2020/9/23
 */

func Modifypwd(ssouser sysmodel.SSOUser, oldpwd, newpwd string) *sysmodel.ResultBean {
	resultBean := &sysmodel.ResultBean{IsSuccess: true}

	if e := checkOldAndNewPwd(oldpwd, newpwd); e != nil {
		return resultBean.SetError("", e.Error(), "")
	}
	userid := ssouser.UserID
	oldencodePwd := encryptutil.EncryptSha256(oldpwd)
	user := &eb.Eb_user{}
	db, _ := conn.GetDB()
	isok, e := db.Where("userid=? and UserPwd = ?", userid, oldencodePwd).NoAutoCondition().Exist(user)
	if e != nil {
		return resultBean.SetError("", fmt.Sprint("修改密码失败 :", e), "")
	}
	if !isok {
		return resultBean.SetError("", "修改密码失败,原密码错误 !", "")
	}

	//修改密码
	user.UserPwd = encryptutil.EncryptSha256(newpwd)
	_, e = db.Where("userid=?", userid).Cols("UserPwd").Update(user)
	if e != nil {
		return resultBean.SetError("", fmt.Sprint("修改密码失败 :", e), "")
	}
	loghelper.ByInfo("修改密码", fmt.Sprintf("用户 %v 修改密码成功！", ssouser.UserName), ssouser.UserID)

	return resultBean
}

func checkOldAndNewPwd(oldpwd, newpwd string) error {
	if len(newpwd) < 6 || len(newpwd) > 16 {
		return fmt.Errorf("修改密码失败: 密码要求长度满足6-16位！")
	}
	if oldpwd == newpwd {
		return fmt.Errorf("新密码不能和原密码一样！")
	}

	return nil
}

func Resetpwd(ssoUser sysmodel.SSOUser, uid string) *sysmodel.ResultBean {
	resultBean := &sysmodel.ResultBean{IsSuccess: true}

	user := &eb.Eb_user{}
	db, _ := conn.GetDB()
	isok, e := db.Where("UserID=?", uid).Cols("usercode").Get(user)
	if e != nil || !isok {
		return resultBean.SetError("", fmt.Sprint("未获取到重置密码的用户:", e), "")
	}
	user.UserPwd = encryptutil.EncryptSha256(user.UserCode)

	_, e = db.Where("UserID=?", uid).Cols("UserPwd").Update(user)
	if e != nil {
		return resultBean.SetError("", fmt.Sprint("重置密码失败 :", e), "")
	}

	return resultBean
}
