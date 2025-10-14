package byaccount

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/luoliDark/base/confighelper"
	"github.com/luoliDark/base/constant"
	"github.com/luoliDark/base/db/conn"
	"github.com/luoliDark/base/db/dbhelper"
	"github.com/luoliDark/base/redishelper"
	"github.com/luoliDark/base/redishelper/rediscache"
	"github.com/luoliDark/base/sysmodel"
	"github.com/luoliDark/base/sysmodel/eb"
	"github.com/luoliDark/base/util/commutil"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
	"github.com/xormplus/xorm"
)

var engine *xorm.Engine

// PC上登录方法
// 验证码生成逻辑
// 打开登录页时，请求验证码接口返回验证码图片，前端传递的 JJPCID，后台生成一个带验证码图片，然后生成一组redis信息，JJPCID 对应验证码，，JJPCID -> capturecode,并且返回图片，
// 当用户点击图片刷新时，再次调用接口，生成新的图片和验证码，JJPCID 保持一致，同时验证码存在redis 设置一个失效时间；
// 用户打开登录页时，首先检查是否存在cookie JJPCID，如果有直接使用获取验证码，否则生成JJPCID 缓存再获取验证码
func Login(luser *sysmodel.SSOUser, password string, c *gin.Context) (*sysmodel.SSOUser, error) {

	engine, _ = conn.GetConnection("admin", true)

	// 验证登录账号密码
	ebuser, err := checkUserPwd(luser, password)
	if err != nil { // 密码验证失败
		return luser, err
	} else {
		//绑定用户信息
		_, err = BindSSOUser(ebuser, luser, luser.IsMobile)
		if err != nil {
			return luser, err
		}

		//写入redis
		if !SetUserSession(luser) {
			return luser, errors.New("登录成功，设置redis时失败")
		}
	}

	return luser, nil
}

func getParamByGetPostAndHeader(ctx *gin.Context, key string) string {
	v := ctx.PostForm(key)
	if v == "" {
		v = ctx.Request.Header.Get(key)
	}
	if v == "" {
		v = ctx.Query(key)
	}
	return v
}

// 检查AppID 是否有效,并返回虚拟用户
func ChkExAppid(ctx *gin.Context) (err error, appid, secret, entid string) {
	//access_token，entcode 只放在了请求头里面，
	access_token := ctx.Request.Header.Get("accessToken") //nginx服务转发后带下斜杠参数请求头无法正常获取
	entcode := ctx.Request.Header.Get("entcode")
	//其他参数兼容从post中再取一次
	timestamp := getParamByGetPostAndHeader(ctx, "timestamp")
	sign := getParamByGetPostAndHeader(ctx, "sign")
	if access_token != "" {
		//按照新的授权方式检查
		tokenInfo, err := CheckOAuthToken(access_token, entcode)
		if err != nil {
			return err, "", "", ""
		} else {
			appid = tokenInfo["appid"]
			appinfo := rediscache.GetHashMap(0, 0, "sys_appinfo", appid)
			return nil, appinfo["appid"], "", appinfo["entid"]
		}
	}
	appid = getParamByGetPostAndHeader(ctx, "appid")
	secret = getParamByGetPostAndHeader(ctx, "secret")

	fmt.Println(timestamp, sign) //todo 待开发 暂不检查
	if !commutil.IsNullOrEmpty(appid) && !commutil.IsNullOrEmpty(secret) {
		appinfo := rediscache.GetHashMap(0, 0, "sys_appinfo", appid)

		if appid == appinfo["appid"] && secret == appinfo["secret"] {
			return nil, appid, secret, appinfo["entid"]
		} else {
			return errors.New("appid或secret无效！"), "", "", ""
		}
	} else {
		return errors.New("参数appid,secret不全"), "", "", ""
	}
}

// 绑定用户信息
func BindSSOUser(ebuser *eb.Eb_user, ssouser *sysmodel.SSOUser, isMobile int8) (*sysmodel.SSOUser, error) {

	//绑定基本信息
	ssouser.UserID = ebuser.UserID
	ssouser.DeptID = ebuser.DeptID
	ssouser.CompID = ebuser.CompID
	ssouser.UserName = ebuser.UserName
	ssouser.OpenID = ebuser.UserOpenID
	ssouser.CWSoftInnerUid = ebuser.CWSoftInnerUid
	ssouser.LoginTime = commutil.TimeFormat(time.Now(), commutil.Time_Fomat01)
	ssouser.IsMobile = isMobile
	ssouser.EntID = commutil.ToString(ebuser.EntId)
	ssouser.BankAccount = ebuser.BankAccount
	ssouser.BankAdd = ebuser.BankAdd

	//绑定服务器组
	ssouser.SGId = "1"

	//查询部门名称及相关部门领导
	rowDepMap := rediscache.GetListMap(commutil.ToInt(ebuser.EntId), 0, "eb_deptusercol", "20102")
	var sbdepuser bytes.Buffer
	for ind, _ := range rowDepMap {
		m := rowDepMap[ind]
		sqlcol := m["sqlcol"]
		sbdepuser.WriteString(strings.ToLower(sqlcol))
		sbdepuser.WriteString(",")
	}

	sql := "select deptname"
	if len(rowDepMap) > 0 {
		sql += "," + strings.TrimRight(sbdepuser.String(), ",")
	}
	sql += " from eb_dept where deptid=?"
	ssouser.DeptManager = make(map[string]string)
	rowM, _ := dbhelper.Query(ssouser.UserID, true, sql, ebuser.DeptID)
	if len(rowM) > 0 {
		m := rowM[0]
		ssouser.DeptName = m["deptname"]
		//部门领导
		for col, val := range m {
			if col == "deptname" {
				continue
			} else {
				ssouser.DeptManager[col] = val
			}
		}
	}

	//查询我管理的部门列表
	qsql := "select deptid as k,deptname  as v from eb_dept where manageruid=?"
	lstm, _ := dbhelper.Query(ssouser.UserID, true, qsql, ssouser.UserID)
	ssouser.MyChildDept = commutil.RowsToMap(lstm)
	ssouser.MyChildDeptIDS = commutil.MapToIdStr(ssouser.MyChildDept)

	//查询我分管的部门列表
	qsql = "select deptid as k,deptname  as v from eb_dept where manager_fguid=?"
	lstFG, _ := dbhelper.Query(ssouser.UserID, true, qsql, ssouser.UserID)
	ssouser.MyFGChildDept = commutil.RowsToMap(lstFG)
	ssouser.MyFGChildDeptIds = commutil.MapToIdStr(ssouser.MyFGChildDept)

	//查询公司名称
	ssouser.CompName, _ = dbhelper.QueryFirstCol(ssouser.UserID, true,
		"select compname from eb_company where compid=?", ebuser.CompID)

	////查询头像
	//ssouser.Image, _ = dbhelper.QueryFirstCol(ssouser.UserID, true,
	//	"select image from eb_userimage where userid=?", ebuser.UserID)

	//绑定企业
	_, err := GetUsersEnt(ssouser)
	if err != nil {
		return ssouser, err
	}

	//绑定角色
	userRoles, _ := dbhelper.Query(ssouser.UserID, true, "select a.roleid,b.rolename from eb_uservsrole a "+
		"left join eb_role b on a.roleid=b.roleid where a.userid=? ", ssouser.UserID)
	roleids := map[string]string{}
	if len(userRoles) > 0 {
		for _, value := range userRoles {
			roleids[value["roleid"]] = value["rolename"]
		}
	}

	if len(roleids) == 0 {
		//默认给普通用户
		roleids["1"] = "普通用户"
	}

	ssouser.UserRoleIds = roleids

	//绑定岗位
	jobRows, _ := dbhelper.Query(ssouser.UserID, true,
		"select a.jobid,b.jobname from eb_uservsjob a "+
			"left join eb_job b on a.jobid=b.jobid where a.userid=?  ", ssouser.UserID)

	userJobs := map[string]string{}
	if len(jobRows) > 0 {
		for _, value := range jobRows {
			roleids[value["jobid"]] = value["jobname"]
		}
	}
	ssouser.UserJobs = userJobs

	return ssouser, nil
}

// 获取值
func GetUserExAttrColValue(defkeyword string, user *sysmodel.SSOUser) string {
	bean := GetUserExAttrCol(defkeyword, user)
	if val := commutil.ToString(bean.ResultData); val != "" {
		return strings.Split(val, "_")[0]
	}
	return ""
}

// 据默认值 取登录人扩展字段信息 例：所在部门 的部门某属性 或所在公司的某属性 （获取后将该字段保存到redis 下个用户登录是直接取值)
func GetUserExAttrCol(defkeyword string, user *sysmodel.SSOUser) sysmodel.ResultBean {
	//默认值关健字 例：user.deptid.depttype  /  user.compid.rate /  user.self.sex
	defkeyword = strings.ReplaceAll(defkeyword, "{", "")
	defkeyword = strings.ReplaceAll(defkeyword, "#", "")
	defkeyword = strings.ReplaceAll(defkeyword, "}", "")
	defArr := strings.Split(defkeyword, ".")
	if len(defArr) != 3 {
		return sysmodel.ResultBean{IsSuccess: false, ErrorMsg: "取用户扩展字段格式不能为" + defkeyword + " 必须为user.deptid.depttype的格式"}
	}
	ctype := defArr[1]
	col := defArr[2]

	var tableName string
	var pkCol string
	var pkVal string
	colMap := make(map[string]string)
	switch ctype {
	case "dept":
		tableName = "eb_dept"
		pkCol = "deptid"
		pkVal = user.DeptID
		colMap = rediscache.GetHashMap(commutil.ToInt(user.EntID), 0, "sys_fpagefieldlogin", "20102"+col)
	case "comp":
		tableName = "eb_company"
		pkCol = "compid"
		pkVal = user.CompID
		colMap = rediscache.GetHashMap(commutil.ToInt(user.EntID), 0, "sys_fpagefieldlogin", "20103"+col)

	case "self":
		//表示查询用户自己其它自定义字段
		tableName = "eb_user"
		pkCol = "userid"
		pkVal = user.UserID
		colMap = rediscache.GetHashMap(commutil.ToInt(user.EntID), 0, "sys_fpagefieldlogin", "20101"+col)

	}
	if g.IsEmpty(tableName) {
		return sysmodel.ResultBean{IsSuccess: false, ErrorMsg: "获取登录用户扩展字段失败请检查配置信息" + defkeyword}
	}
	qsql := commutil.AppendStr("select ", col, " from ", tableName, " where ", pkCol, "=?")
	val, err := dbhelper.QueryFirstCol("", true, qsql, pkVal)
	var dsName string //数据源名称
	if err != nil {
		return sysmodel.ResultBean{IsSuccess: false, ErrorMsg: err.Error()}
	}

	if len(colMap) > 0 {
		ds := strings.Trim(colMap["datasource"], " ")
		if !g.IsEmpty(ds) {
			dsM := rediscache.GetHashMap(commutil.ToInt(user.EntID), commutil.ToInt(ds), "sys_fpage", ds)
			if len(dsM) > 0 {
				dsSqlTab := dsM["sqltablename"]
				pkcol := dsM["primarykey"]
				NameCol := dsM["namecol"]
				qdsSql := commutil.AppendStr("select ", NameCol, " from ", dsSqlTab, " where ", pkcol, " =?")
				dsName, _ = dbhelper.QueryFirstCol("", true, qdsSql, val)
				if !g.IsEmpty(dsName) {
					val = val + "_" + dsName
				}
			}
		}
	}

	//写到redis中，下次登录时直接将该字段 写到用户对象中，就不用再调本接口查询
	key := constant.LoginExCol
	m := redishelper.GetHashMap(confighelper.GetEnterpriseID(), confighelper.GetSessionDbIndex(), key)
	if m == nil || len(m) == 0 {
		m = make(map[string]string)
	}
	tmpCols, ok := m[ctype]
	if ok {
		if g.IsEmpty(tmpCols) {
			tmpCols = col
		} else if !strings.Contains(","+tmpCols+",", ","+col+",") {
			//两边加豆号 防止字段包包含重命
			tmpCols = commutil.AppendStr(tmpCols, ",", col)
		}
		m[ctype] = tmpCols
	} else {
		m[ctype] = col
	}

	//写到redis
	redishelper.SetHashMap(confighelper.GetEnterpriseID(), confighelper.GetSessionDbIndex(), key, m)

	return sysmodel.ResultBean{IsSuccess: true, ResultData: val}
}

// 写入cookie
func WriteCookie(luser *sysmodel.SSOUser, c *gin.Context) bool {
	origin := c.Request.Header.Get("Origin")
	time := 3600 * 24 * 30 //1个月
	//c.SetCookie("mver", luser.Mver, time, "/", origin, false, true)
	c.SetCookie("toid", luser.SId, time, "/", origin, false, true)
	c.SetCookie("sgid", luser.SGId, time, "/", origin, false, true)
	return true
}

// 选择企业ID
func SelectEnt(user *sysmodel.SSOUser) *sysmodel.ResultBean {

	//更新redis中信息
	jsonvalue := commutil.ObjectToJson(&user)
	isok := redishelper.SetString(confighelper.GetEnterpriseID(), confighelper.GetSessionDbIndex(), user.UserID, jsonvalue)
	return &sysmodel.ResultBean{IsSuccess: isok}

}
