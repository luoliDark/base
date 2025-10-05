package byaccount

import (
	"errors"
	"fmt"

	"github.com/luoliDark/base/db/dbhelper"
	"github.com/luoliDark/base/loghelper"
	"github.com/luoliDark/base/redishelper/rediscache"
	ssomodel "github.com/luoliDark/base/sso/ssologin/model"
	"github.com/luoliDark/base/sysmodel"
	"github.com/luoliDark/base/sysmodel/eb"
	"github.com/luoliDark/base/util/commutil"
	"github.com/luoliDark/base/util/encryptutil"
)

func checkUserPwd(ssouser *sysmodel.SSOUser, password string) (*eb.Eb_user, error) {

	user := new(eb.Eb_user)

	// 用户登录检查  用户总表，
	if ssouser.UserCode == "" || password == "" {
		return user, errors.New("账号和密码不能为空")
	}

	has, err := engine.Where("usercode=? and ifnull(IsDiscard,0) = 0  ", ssouser.UserCode).Get(user)
	if err != nil {
		loghelper.ByError("登录失败", fmt.Sprint("登录失败,查询用户发生错误:", err.Error()), ssouser.UserID)
		return user, errors.New("系统异常请稍后再试")
	}
	if !has {
		// 没有账户
		return user, errors.New("账号不存在")
	}
	// 用户状态，0 表示被弃用账户，1 表示有效用户，2 表示用户由于连续登陆失败被锁定,3 表示用户被冻结， 4 表示用户已过期（涉及到会员制）
	// 检查账户是否有效
	if user.Status == 1 || user.IsDiscard == 1 {
		// 账户无效
		return user, errors.New("账号无效")
	} else if user.Status == 2 {
		return user, errors.New("账号被锁定")
	} else if user.Status == 3 {
		return user, errors.New("账号被冻结")
	} else if user.Status == 4 {
		return user, errors.New("账号已过期")
	} else {
		// 对比密码
		encodePwd := encryptutil.EncryptSha256(password)
		if encodePwd != user.UserPwd {
			// 密码错误
			return user, errors.New("密码错误")
		}
	}

	//验证成功
	return user, nil
}

// 检查用户企业信息
func GetUsersEnt(user *sysmodel.SSOUser) (IsGetCompByUser bool, err error) {

	lst, err := dbhelper.Query(user.UserID, true,
		"select  b.isnotctrwfnextuser, b.isgetcompbyuser, b.isopendingphone, b.entid,b.entname,b.logo,b.menuver,b.formentid "+
			" from eb_entvsuser a join eb_enterprise b on a.entid=b.EntID where a.userid=?", user.UserID)

	if err != nil {
		//查询菜单版本
		globalm := rediscache.GetHashMap(0, 0, "sys_global", "1")
		ent := make([]ssomodel.EB_EntVsUser, 1)
		ent[0] = ssomodel.EB_EntVsUser{EntID: "0", EntName: "未知企业", Mver: globalm["menuver"], FormEntId: "0"}
		user.EnList = ent
		user.FormEntId = "0"
		return IsGetCompByUser, err
	} else {

		entvsuser := make([]ssomodel.EB_EntVsUser, len(lst))
		for index, m := range lst {
			entvsuser[index] = ssomodel.EB_EntVsUser{EntID: m["entid"], EntName: m["entname"], Logo: m["logo"],
				Mver: m["menuver"], FormEntId: m["formentid"], IsOpenOcr: commutil.ToBool(m["isopenocr"]),
				IsCostCenter:       commutil.ToBool(m["iscostcenter"]),
				IsVatDetail:        commutil.ToBool(m["isvatdetail"]),
				IsExWf:             commutil.ToBool(m["isexwf"]),
				IsNotCtrWFNextUser: commutil.ToBool(m["isnotctrwfnextuser"]),
				SsoTime:            commutil.ToInt(m["SsoTime"]),
			}

			user.IsOpenDingPhone = commutil.ToInt(m["isopendingphone"]) //钉钉电话通知

		}
		entM := rediscache.GetHashMap(0, 0, "eb_enterprise", user.EntID)
		if len(entM) == 0 {
			return IsGetCompByUser, fmt.Errorf("当前登录企业未启用，请联系管理员！")
		}

		user.EnList = entvsuser
		user.IsOpenDingPhone = commutil.ToInt(entM["isopendingphone"]) //钉钉电话通知
		IsGetCompByUser = commutil.ToBool(entM["isgetcompbyuser"])     //登录人公司取员工表上的compid

		//如果从对照表没取到企业，就从用户档案entid 直接取该企业
		if len(user.EnList) == 0 {
			ent := make([]ssomodel.EB_EntVsUser, 1)
			ent[0] = ssomodel.EB_EntVsUser{EntID: user.EntID, EntName: entM["entname"], Mver: entM["menuver"],
				FormEntId: entM["formentid"], IsOpenOcr: commutil.ToBool(entM["isopenocr"]),
				IsCostCenter:       commutil.ToBool(entM["iscostcenter"]),
				IsVatDetail:        commutil.ToBool(entM["isvatdetail"]),
				IsExWf:             commutil.ToBool(entM["isexwf"]),
				IsNotCtrWFNextUser: commutil.ToBool(entM["isnotctrwfnextuser"]),
				SsoTime:            commutil.ToInt(entM["ssotime"]),
				FileView_Ver:       commutil.ToInt(entM["fileview_ver"]),
			}
			user.EnList = ent
			user.FormEntId = entM["formentid"]
			user.IsOpenOcr = commutil.ToBool(entM["isopenocr"])
			user.SsoTime = commutil.ToInt(entM["ssotime"])
			user.FileView_Ver = commutil.ToInt(entM["fileview_ver"])
			user.IsVatDetail = commutil.ToBool(entM["isvatdetail"])
			user.IsCostCenter = commutil.ToBool(entM["iscostcenter"])
			user.IsExWf = commutil.ToBool(entM["isexwf"])
		} else if len(user.EnList) == 1 {
			// 只有一家企业时默认选中
			user.EntID = user.EnList[0].EntID
			user.FormEntId = user.EnList[0].FormEntId
			user.IsOpenOcr = user.EnList[0].IsOpenOcr
			user.SsoTime = user.EnList[0].SsoTime
			user.FileView_Ver = user.EnList[0].FileView_Ver
			user.IsVatDetail = user.EnList[0].IsVatDetail
			user.IsCostCenter = user.EnList[0].IsCostCenter
			user.IsExWf = user.EnList[0].IsExWf
		} else if len(user.EnList) > 1 {
			//todo 临时代码 否则手机无法登录
			//补充：手机端现不支持选择多企业，登录会导致 餐琪登录到 成都等其他企业，且不能切换。
			// 不能正常使用。一般手机端使用自己企业的账号， 先默认登录到用户表维护的企业，
			user.FormEntId = entM["formentid"] //如果有2个以上企业默认fromid 就用员工本人的
			user.IsOpenOcr = commutil.ToBool(entM["isopenocr"])
			user.SsoTime = commutil.ToInt(entM["ssotime"])
			user.FileView_Ver = commutil.ToInt(entM["fileview_ver"])
			user.IsCostCenter = commutil.ToBool(entM["iscostcenter"])
			user.IsVatDetail = commutil.ToBool(entM["isvatdetail"])
			user.IsExWf = commutil.ToBool(entM["isexwf"])
		}

		if user.FormEntId == "" {
			user.FormEntId = user.EntID
		}
		return IsGetCompByUser, nil
	}
}
