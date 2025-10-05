package byaccount

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
	"github.com/luoliDark/base/confighelper"
	"github.com/luoliDark/base/constant"
	"github.com/luoliDark/base/db/conn"
	"github.com/luoliDark/base/db/dbhelper"
	"github.com/luoliDark/base/loghelper"
	"github.com/luoliDark/base/redishelper"
	"github.com/luoliDark/base/redishelper/rediscache"
	"github.com/luoliDark/base/sysmodel"
	"github.com/luoliDark/base/sysmodel/eb"
	"github.com/luoliDark/base/util/commutil"
	"github.com/luoliDark/base/util/encryptutil"
	"github.com/xormplus/xorm"
)

var engine *xorm.Engine

func init() {
	//清除用户关联字段缓存，否则会一直存在。
	var enterpriseID = confighelper.GetEnterpriseID()
	var dbIndex = confighelper.GetSessionDbIndex()
	keyPreStr := commutil.AppendStr(enterpriseID, "_", constant.LoginExCol)
	//检查该历史版本是否存在数据，如果存在则模糊查找清除
	re := redishelper.GetString(enterpriseID, dbIndex, constant.LoginExCol)
	if !g.IsEmpty(re) {
		//clear
		redishelper.DeleleByLike(dbIndex, keyPreStr)
	}
}

//PC上登录方法
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

//绑定用户信息
func BindSSOUser(ebuser *eb.Eb_user, ssouser *sysmodel.SSOUser, isMobile int8) (*sysmodel.SSOUser, error) {

	//绑定基本信息
	ssouser.UserID = ebuser.UserID
	ssouser.UserCode = ebuser.UserCode
	ssouser.DeptID = ebuser.DeptID
	ssouser.UserName = ebuser.UserName
	ssouser.OpenID = ebuser.UserOpenID
	ssouser.CWSoftInnerUid = ebuser.CWSoftInnerUid
	ssouser.LoginTime = commutil.TimeFormat(time.Now(), commutil.Time_Fomat01)
	ssouser.IsMobile = isMobile
	ssouser.EntID = commutil.ToString(ebuser.EntId)
	ssouser.BankAccount = ebuser.BankAccount
	ssouser.BankAdd = ebuser.BankAdd
	ssouser.Sex = ebuser.Sex
	ssouser.UserLevel = ebuser.UserLevel
	ssouser.ImgSrc = ebuser.ImgSrc
	ssouser.UserPhone = ebuser.UserPhone
	ssouser.UserEmail = ebuser.UserEmail

	//绑定企业
	IsGetCompByUser, err := GetUsersEnt(ssouser)
	if err != nil {
		return ssouser, err
	}

	//员工公司
	if IsGetCompByUser {
		ssouser.CompID = ebuser.CompID //修改为使用部门对应的公司
	}

	//绑定服务器组
	ssouser.SGId = "1"

	//绑定用户登录扩展字段
	key := constant.LoginExCol
	//缓存放在 cacheDbIndex 下，清除缓存才会清除 ;
	m := redishelper.GetHashMap(confighelper.GetEnterpriseID(), confighelper.GetSessionDbIndex(), key)
	var deptExCol string
	var compExCol string
	var selfExCol string

	ssouser.ExLoginCol = make(map[string]string)

	if m != nil && len(m) > 0 {
		deptExCol = m["dept"]
		compExCol = m["comp"]
		selfExCol = m["self"]
	}

	var userM []map[string]string

	if !g.IsEmpty(selfExCol) {
		sql := "select " + selfExCol + " from eb_user where userid=?"
		userM, err = dbhelper.Query(ssouser.UserID, true, sql, ebuser.UserID)
		if err != nil {
			return nil, fmt.Errorf("查询用户信息拓展信息失败:%v", err)
		}
	}

	//查询部门名称及相关部门领导
	rowDepMap := rediscache.GetListMap(ebuser.EntId, 0, "eb_deptusercol", "20102")
	var sbdepuser bytes.Buffer
	for ind, _ := range rowDepMap {
		m := rowDepMap[ind]
		sqlcol := m["sqlcol"]
		sbdepuser.WriteString(strings.ToLower(sqlcol))
		sbdepuser.WriteString(",")
	}

	sql := "select b.compid, b.defdeptidbyuseradd,b.buddeptid,b.depttype," +
		"(select deptname from eb_dept a where a.deptid=b.DefDeptIdByUserAdd limit 1)as defdeptnamebyadd, b.deptname"
	if len(rowDepMap) > 0 {
		sql += "," + strings.TrimRight(sbdepuser.String(), ",")
	}

	//部门扩展字段
	if !g.IsEmpty(deptExCol) {
		sql += "," + deptExCol
	}

	sql += " from eb_dept b where b.deptid=?"
	ssouser.DeptManager = make(map[string]string)
	deptM, err := dbhelper.Query(ssouser.UserID, true, sql, ebuser.DeptID)
	if err != nil {
		return nil, fmt.Errorf("查询所在的部门信息:%v", err)
	}
	if deptM != nil && len(deptM) > 0 {

		m := deptM[0]
		ssouser.DeptType = m["depttype"]
		ssouser.DeptName = m["deptname"]
		ssouser.BudDeptId = m["buddeptid"]
		if !IsGetCompByUser {
			ssouser.CompID = m["compid"] //取部门档案公司做为登录人公司
		}

		ebuser.CompID = ssouser.CompID

		//新增默认带出的部门
		ssouser.DefDeptIdByAdd = m["defdeptidbyuseradd"]
		ssouser.DefDeptNameByAdd = m["defdeptnamebyadd"]
		//部门领导
		for col, val := range m {
			if col == "deptname" {
				continue
			} else {
				ssouser.DeptManager[col] = val
			}
		}
	} else {
		//未查询到部门，设置为空 防止前端设置默认值时，设置了部门ID未设置名称，造成必填检查时效，数据异常。
		ssouser.DeptID = ""
	}

	//查询我管理的部门列表
	qsql := "select deptid as k,deptname  as v from eb_dept where manageruid=?"
	lstm, err := dbhelper.Query(ssouser.UserID, true, qsql, ssouser.UserID)
	if err != nil {
		return nil, fmt.Errorf("查询管理的部门列表:%v", err)
	}
	ssouser.MyChildDept = commutil.RowsToMap(lstm)
	ssouser.MyChildDeptIDS = commutil.MapToIdStr(ssouser.MyChildDept)

	//查询我分管的部门列表
	qsql = "select deptid as k,deptname  as v from eb_dept where manager_fguid=?"
	lstFG, _ := dbhelper.Query(ssouser.UserID, true, qsql, ssouser.UserID)
	ssouser.MyFGChildDept = commutil.RowsToMap(lstFG)
	ssouser.MyFGChildDeptIds = commutil.MapToIdStr(ssouser.MyFGChildDept)

	//查询公司名称
	qcompSql := "select compname  "
	//公司扩展字段
	if !g.IsEmpty(compExCol) {
		qcompSql += "," + compExCol
	}

	qcompSql = qcompSql + "  from eb_company where compid=?"
	compM, _ := dbhelper.Query(ssouser.UserID, true, qcompSql, ssouser.CompID)
	if len(compM) > 0 {
		ssouser.CompName = compM[0]["compname"]
	}

	////查询头像
	//ssouser.Image, _ = dbhelper.QueryFirstCol(ssouser.UserID, true,
	//	"select image from eb_userimage where userid=?", ebuser.UserID)

	//绑定角色
	userRoles, _ := dbhelper.Query(ssouser.UserID, true, "select a.roleid,b.rolename from eb_uservsrole a "+
		"left join eb_role b on a.roleid=b.roleid where a.userid=? ", ssouser.UserID)
	roleids := map[string]string{}
	isHasNormalRole := false //是否有普通用户权限
	if len(userRoles) > 0 {
		for _, value := range userRoles {
			roleId := value["roleid"]
			if roleId == "1" {
				isHasNormalRole = true //如果有普通用户
			}
			roleids[roleId] = value["rolename"]
		}
	}

	if len(roleids) == 0 || !isHasNormalRole {
		//默认给普通用户
		roleids["1"] = "普通用户"
	}

	ssouser.UserRoleIds = roleids

	//出纳
	isClOrCw := ssouser.UserRoleIds["5"] != "" || ssouser.UserRoleIds["4"] != ""
	if isClOrCw {
		//查询我管理的公司 chuna caiwu
		qsql = `SELECT a.compid as k,a.compname as v from eb_company a 
 join eb_companyvsuser b on a.compid = b.compid where b.userid=?  `
		lstCMP, err := dbhelper.Query(ssouser.UserID, true, qsql, ssouser.UserID)
		if err != nil {
			return nil, fmt.Errorf("查询所管理公司信息失败:%v", err)
		}
		if len(lstCMP) > 0 {
			ssouser.MyChildComp = commutil.RowsToMap(lstCMP)
			ssouser.MyChildCompIds = commutil.MapToIdStr(ssouser.MyChildComp)
		} else {
			qsql = "SELECT compid as k,compname  as v from 	eb_company where Chuna_Uid=?"
			lstCMP, err = dbhelper.Query(ssouser.UserID, true, qsql, ssouser.UserID)
			if err != nil {
				return nil, fmt.Errorf("查询所管理公司信息失败:%v", err)
			}
			ssouser.MyChildComp = commutil.RowsToMap(lstCMP)
			ssouser.MyChildCompIds = commutil.MapToIdStr(ssouser.MyChildComp)
		}
	}

	//绑定岗位
	jobRows, _ := dbhelper.Query(ssouser.UserID, true,
		"select a.jobid,b.jobname from eb_uservsjob a "+
			"left join eb_job b on a.jobid=b.jobid where a.userid=?  ", ssouser.UserID)

	userJobs := map[string]string{}
	if len(jobRows) > 0 {
		for _, value := range jobRows {
			userJobs[value["jobid"]] = value["jobname"]
		}
	}
	ssouser.UserJobs = userJobs

	//登录扩展字段
	if !g.IsEmpty(deptExCol) && deptM != nil {
		m := deptM[0]
		colArr := strings.Split(deptExCol, ",")
		for _, c := range colArr {
			newVal, err := getExDataSourceColNameValue(ssouser.EntID, "dept", c, m[c])
			if err != nil {
				return ssouser, errors.New("获取登录用户扩展字段失败" + err.Error())
			}
			k := commutil.AppendStr("#user.dept.", c)
			ssouser.ExLoginCol[k] = newVal
		}
	}

	if !g.IsEmpty(compExCol) && compM != nil {
		m := compM[0]
		colArr := strings.Split(compExCol, ",")
		for _, c := range colArr {
			newVal, err := getExDataSourceColNameValue(ssouser.EntID, "comp", c, m[c])
			if err != nil {
				return ssouser, errors.New("获取登录用户扩展字段失败" + err.Error())
			}
			k := commutil.AppendStr("#user.comp.", c)
			ssouser.ExLoginCol[k] = newVal
		}
	}

	if !g.IsEmpty(selfExCol) && userM != nil {
		m := userM[0]
		colArr := strings.Split(selfExCol, ",")
		for _, c := range colArr {
			newVal, err := getExDataSourceColNameValue(ssouser.EntID, "self", c, m[c])
			if err != nil {
				return ssouser, errors.New("获取登录用户扩展字段失败" + err.Error())
			}
			k := commutil.AppendStr("#user.self.", c)
			ssouser.ExLoginCol[k] = newVal
		}
	}
	//自定义全局性配置，例如只需要指定客户启用或关闭的一些功能、信息在用户登录时，初始化在ExConfig参数
	configList, _ := dbhelper.Query(ssouser.UserID, false, " select lower(configcode) configcode,configname from sys_global_CustomConfig where entid = ? and isopen = 1 ", ssouser.EntID)
	ssouser.ExConfig = make(map[string]string, len(configList))
	for _, c := range configList {
		ssouser.ExConfig[c["configcode"]] = c["configname"]
	}
	return ssouser, nil
}

func getExDataSourceColNameValue(entid, ctype, col, colvalue string) (string, error) {
	//默认值关健字 例：user.deptid.depttype  /  user.compid.rate /  user.self.sex
	colMap := make(map[string]string)
	switch ctype {
	case "dept":
		colMap = rediscache.GetHashMap(commutil.ToInt(entid), 0, "sys_fpagefieldlogin", "20102"+col)
	case "comp":
		colMap = rediscache.GetHashMap(commutil.ToInt(entid), 0, "sys_fpagefieldlogin", "20103"+col)
	case "self":
		//表示查询用户自己其它自定义字段
		colMap = rediscache.GetHashMap(commutil.ToInt(entid), 0, "sys_fpagefieldlogin", "20101"+strings.ToLower(col))
	default:
		return "", errors.New("获取登录用户扩展字段失败请检查配置信息 " + ctype)
	}
	if len(colMap) == 0 {
		return colvalue, nil
	}
	var dsName string //数据源名称
	if len(colMap) > 0 {
		ds := strings.Trim(colMap["datasource"], " ")
		if !g.IsEmpty(ds) {
			dsM := rediscache.GetHashMap(commutil.ToInt(entid), commutil.ToInt(ds), "sys_fpage", ds)
			if len(dsM) > 0 {
				dsSqlTab := dsM["sqltablename"]
				pkcol := dsM["primarykey"]
				NameCol := dsM["namecol"]
				qdsSql := commutil.AppendStr("select ", NameCol, " from ", dsSqlTab, " where ", pkcol, " =?")
				dsName, _ = dbhelper.QueryFirstCol("", true, qdsSql, colvalue)
				if !g.IsEmpty(dsName) {
					colvalue = colvalue + "_" + dsName
				}
			}
		}
	}
	return colvalue, nil
}

func GetSubmmitUserByCreate_uid(entid, pid int, primarykey string) (sysmodel.SSOUser, error) {
	return GetSubmmitUserByUserCol(entid, pid, primarykey, constant.CreateUidColName)
}

func GetSubmmitUserByUserCol(entid, pid int, primarykey string, col string) (sysmodel.SSOUser, error) {
	if col == "" {
		col = constant.CreateUidColName
	}
	session, _ := conn.GetSession()
	defer session.Close()
	hashMap := rediscache.GetHashMap(entid, pid, "sys_fpage", commutil.ToString(pid))
	sqltablename := hashMap["sqltablename"]
	create_uid, err := dbhelper.QueryFirstColByTran(session, "", true,
		fmt.Sprintf("select %v from %v where billid = ?", col, sqltablename), primarykey)
	if err != nil {
		return sysmodel.SSOUser{}, fmt.Errorf("未查询到制单人：%v", err)
	}
	user := &eb.Eb_user{}
	sso := &sysmodel.SSOUser{}
	_, e := session.Where("userid=?", create_uid).Get(user)
	if e != nil {
		return sysmodel.SSOUser{}, fmt.Errorf("查询用户信息失败：%v", err)
	}
	sso, err = BindSSOUser(user, sso, 1)
	if err != nil {
		return sysmodel.SSOUser{}, fmt.Errorf("查询用户信息失败：%v", err)
	}

	return *sso, nil
}

//写入cookie
func WriteCookie(luser *sysmodel.SSOUser, c *gin.Context) bool {
	origin := c.Request.Header.Get("Origin")
	time := 3600 * 24 * 30 //1个月
	//c.SetCookie("mver", luser.Mver, time, "/", origin, false, true)
	c.SetCookie("toid", luser.SId, time, "/", origin, false, true)
	c.SetCookie("sgid", luser.SGId, time, "/", origin, false, true)
	return true
}

//选择企业ID
func SelectEnt(user *sysmodel.SSOUser) *sysmodel.ResultBean {

	//更新redis中信息
	jsonvalue := commutil.ObjectToJson(&user)
	isok := redishelper.SetString(confighelper.GetEnterpriseID(), confighelper.GetSessionDbIndex(), user.UserID, jsonvalue)
	return &sysmodel.ResultBean{IsSuccess: isok}

}

//验证是否需要修改密码
func CheckUpdatePassword(user sysmodel.SSOUser) sysmodel.ResultBean {
	userid := user.UserID
	entid := user.EntID
	if commutil.IsNullOrEmpty(userid) || commutil.IsNullOrEmpty(entid) {
		return sysmodel.ResultBean{IsSuccess: false, ErrorCode: "500", ErrorMsg: "当前登录信息为空，无法验证是否重置密码!"}
	}
	userSql := "select userpwd from eb_user where userid = '" + userid + "' limit 1"
	userInfo, err := dbhelper.Query(userid, false, userSql)
	if err != nil || len(userInfo) <= 0 {
		return sysmodel.ResultBean{IsSuccess: false, ErrorCode: "500", ErrorMsg: "查询当前登录用户失败！"}
	}

	defPwdSql := "select defpwd from eb_enterprise where entid = '" + entid + "' limit 1"
	entInfo, err := dbhelper.Query(userid, false, defPwdSql)
	if err != nil || len(entInfo) <= 0 {
		return sysmodel.ResultBean{IsSuccess: false, ErrorCode: "500", ErrorMsg: "查询当前企业失败！"}
	}
	defPwd := entInfo[0]["defpwd"]
	userpwd := userInfo[0]["userpwd"]
	if userpwd == encryptutil.EncryptSha256(defPwd) || userpwd == encryptutil.EncryptSha256(user.UserCode) {
		return sysmodel.ResultBean{IsSuccess: false, ErrorCode: "402"}
	} else {
		return sysmodel.ResultBean{IsSuccess: true}
	}
}

//据默认值 取登录人扩展字段信息 例：所在部门 的部门某属性 或所在公司的某属性 （获取后将该字段保存到redis 下个用户登录是直接取值)
func GetUserExAttrCol(entid, defkeyword string, user *sysmodel.SSOUser) sysmodel.ResultBean {
	//默认值关健字 例：user.deptid.depttype  /  user.compid.rate /  user.self.sex
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
		colMap = rediscache.GetHashMap(commutil.ToInt(user.EntID), 0, "sys_fpagefieldlogin", "20101"+strings.ToLower(col))
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
			dsM := rediscache.GetHashMap(commutil.ToInt(entid), commutil.ToInt(ds), "sys_fpage", ds)
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

func ExGetUserByUid(userId string) (*sysmodel.SSOUser, error) {

	user := eb.Eb_user{}
	ssouser := &sysmodel.SSOUser{}

	eng, _ := conn.GetDB()

	_, err := eng.Where(" userId=? and (IsDiscard = 0 or IsDiscard is null)  ", userId).Get(&user)
	if err != nil {
		loghelper.ByError("登录失败", fmt.Sprintf("登录失败,查询用户%v发生错误:%v", userId, err.Error()), "")
		return ssouser, errors.New("系统异常请稍后再试")
	}

	ssouser.UserID = user.UserID

	//绑定用户信息
	ssouser, err = BindSSOUser(&user, ssouser, 0)

	return ssouser, nil
}

/**
 * 解密后，获取用户信息
 * add by bxn-liaq
 */
func GetUserCode(usk string) (string, error) {
	k := confighelper.GetDesKey()
	usercode, _ := encryptutil.DesDecrypt(usk, []byte(k))
	if g.IsEmpty(usercode) {
		return "", errors.New("对称解密失败,请联系管理员 usk=" + usk)
	} else {
		return usercode, nil
	}
}

//des对称解密 登录
func DesLogin(luser *sysmodel.SSOUser, usk string) (*sysmodel.SSOUser, error) {
	//des对称解密
	k := confighelper.GetDesKey()
	usercode, _ := encryptutil.DesDecrypt(usk, []byte(k))

	if g.IsEmpty(usercode) {
		loghelper.ByError("对称解密失败", "usk="+usk, "")
		return &sysmodel.SSOUser{}, errors.New("对称解密失败,请联系管理员")
	}
	return LoginByUserCode(luser, usercode)
}
func LoginByUserCode(luser *sysmodel.SSOUser, userCode string) (*sysmodel.SSOUser, error) {

	engine, _ := conn.GetConnection("admin", true)
	user := new(eb.Eb_user)
	has, err := engine.Where("usercode=?", userCode).Get(user)
	if err != nil {
		return &sysmodel.SSOUser{}, errors.New("系统异常请稍后再试")
	}
	if !has {
		return &sysmodel.SSOUser{}, errors.New("账号不存在")
	}
	//绑定用户信息
	_, err = BindSSOUser(user, luser, luser.IsMobile)
	if err != nil {
		return luser, err
	}
	//写入redis
	if !SetUserSession(luser) {
		return luser, errors.New("登录成功，设置redis时失败")
	}
	return luser, nil
}

//对appid 进行验证
func ExLoginByAppID(ctx *gin.Context) (sysmodel.SSOUser, error) {
	ssouser := sysmodel.SSOUser{}
	eb_user := new(eb.Eb_user)
	// 检查 appid
	err, appid, _, _ := ChkExAppid(ctx)
	if err != nil {
		return ssouser, err
	}

	userCode := ctx.PostForm("logininusercode")
	ismobile := ctx.PostForm("ismobile")
	ssouser.UserCode = userCode
	engine, _ := conn.GetDB()
	_, err = engine.Where(" (usercode=? or userid=? ) and (isdiscard=0 or isdiscard is null)",
		userCode, userCode).Get(eb_user)
	if err != nil || commutil.IsNullOrEmpty(eb_user.UserCode) {
		return ssouser, errors.New("查无此用户")
	}
	//移动登录
	if commutil.ToInt(ismobile) == 1 {
		ssouser.IsMobile = 1
	} else {
		ssouser.IsMobile = 0
	}
	//绑定用户信息
	_, err = BindSSOUser(eb_user, &ssouser, 0)
	if err != nil {
		return ssouser, errors.New("验证APPID，绑定用户失败")
	}
	ssouser.AppId = appid
	//写入redis
	if !SetUserSession(&ssouser) {
		return ssouser, errors.New("登录成功，设置redis时失败")
	}
	return ssouser, nil

}

//检查AppID 是否有效,并返回虚拟用户
func ChkExAppid(ctx *gin.Context) (err error, appid, secret, entid string) {

	appid = ctx.PostForm("appid")
	secret = ctx.PostForm("secret")
	timestamp := ctx.PostForm("timestamp") //时间戳检查
	sign := ctx.PostForm("sign")           //参数签名
	fmt.Println(timestamp, sign)           //todo 待开发 暂不检查

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

//检查AppID 是否有效,并返回虚拟用户
func ChkAppIdFromRedis(appid, secret string) (err error, entid string) {

	if !commutil.IsNullOrEmpty(appid) && !commutil.IsNullOrEmpty(secret) {

		appinfo := rediscache.GetHashMap(0, 0, "sys_appinfo", appid)

		if appid == appinfo["appid"] && secret == appinfo["secret"] {
			return nil, appinfo["entid"]
		} else {
			return errors.New("appid或secret无效！"), ""
		}
	} else {
		return errors.New("参数appid,secret不全"), ""
	}

}
