package sysmodel

import ssomodel "paas/base/sso/ssologin/model"

//登录用户对象信息 用于保存到redis sessionDB中
type SSOUser struct {
	UserCode       string `xorm:"usercode" json:"usercode"`
	IsEnc          bool   //是否做数据脱敏处理
	IsOpenOcr      bool   //是否启用OCR识别功能
	IsExWf         bool   //是否使用外部流程系统进行审批 例：外部BPM 外部OA
	SsoTime        int    //登录信息保存时间 -1表示永久，24表示保留一天
	FileView_Ver   int    //文件预览版本
	IsCostCenter   bool   //是否启用成本中心
	IsVatDetail    bool   //是否启用OCR增值专票明细识别及创建报销
	AppId          string //接口请求时对应的AppId
	UserID         string `xorm:"userid" json:"userid"`
	CWSoftInnerUid string `xorm:"cwsoftinneruid" json:"cwsoftinneruid"`
	LoginUid       string `xorm:"loginuid" json:"loginuid"`
	UserName       string `xorm:"username" json:"username"`
	UserPhone      string `xorm:"userphone" json:"userphone"` // 用户手机号
	UserEmail      string `xorm:"useremail" json:"useremail"` // 用户邮箱
	//7表示加密查看
	UserRoleIds      map[string]string       `xorm:"userroleids" json:"userroleids"`           //角色清单，用map原因是为判定时好取值
	UserJobs         map[string]string       `xorm:"userjobs" json:"userjobs"`                 //角色岗位，用map原因是为判定时好取值 key jobid jobvalue
	DeptID           string                  `xorm:"deptid" json:"deptid"`                     //所属部门
	DeptName         string                  `xorm:"deptname" json:"deptname"`                 //所属部门
	DeptType         string                  `xorm:"depttype" json:"depttype"`                 //所属部门
	DefDeptIdByAdd   string                  `xorm:"defdeptidbyadd" json:"defdeptidbyadd"`     //新增时默认部门
	DefDeptNameByAdd string                  `xorm:"defdeptnamebyadd" json:"defdeptnamebyadd"` //新增时默认部门
	DeptManager      map[string]string       `xorm:"deptmanager" json:"deptmanager"`           //所属部门的审批人
	MyChildDept      map[string]string       `xorm:"mychilddept" json:"mychilddept"`           //我直接管理的部门列表
	MyFGChildDept    map[string]string       `xorm:"myfgchilddept" json:"myfgchilddept"`       //我分管的部门列表
	MyChildComp      map[string]string       `xorm:"mychildcomp" json:"mychildcomp"`           //我管理的公司列表
	MyChildCompIds   string                  `xorm:"mychildcompids" json:"mychildcompids"`     //我管理的公司列表
	MyChildDeptIDS   string                  `xorm:"mychilddeptids" json:"mychilddeptids"`     //我直接管理的部门列表
	MyFGChildDeptIds string                  `xorm:"myfgchilddeptids" json:"myfgchilddeptids"` //我分管的部门列表
	CompID           string                  `xorm:"compid" json:"compid"`                     //所属公司
	CompName         string                  `xorm:"compname" json:"compname"`                 //所属公司
	EntID            string                  `xorm:"entid" json:"entid"`                       //登录时选择的企业ID
	EntName          string                  `xorm:"entname" json:"entname"`                   //登录时选择的企业ID
	FormEntId        string                  `xorm:"formentid" json:"formentid"`               //该企业使用的表单是那一套 可以多企业使用一个公用企业的模板
	OpenID           string                  `xorm:"openid" json:"openid"`                     //微信opid
	IsMobile         int8                    `xorm:"ismobile" json:"ismobile"`                 //是否为移动端登录
	LangID           string                  `xorm:"langid" json:"langid"`                     // 语言ID 例：cn,en等
	LoginTime        string                  `xorm:"logintime" json:"logintime"`               //登录时间
	SId              string                  `xorm:"sid" json:"sid"`                           // 登录成功后生成的token 令牌
	SGId             string                  `xorm:"sgid" json:"sgid"`                         // 服务器组ID 用于负载均衡使用
	EnList           []ssomodel.EB_EntVsUser `xorm:"enlist" json:"enlist"`                     //当前用户的企业信息
	BankAccount      string                  `xorm:"bankaccount" json:"bankaccount"`
	BankAdd          string                  `xorm:"bankadd" json:"bankadd"`
	MsgServerUrl     string                  `xorm:"msgserverurl" json:"msgserverurl"`
	UserLevel        string                  `xorm:"userlevel" json:"userlevel"`
	Sex              string                  `xorm:"sex" json:"sex"`
	ImgSrc           string                  `xorm:"imgsrc" json:"imgsrc"`
	IsDingDing       int                     `xorm:"isdingding" json:"isdingding"`
	IsOpenDingPhone  int                     `xorm:"isopendingphone" json:"isopendingphone"`
	//登录的扩展字段 例：user.self.sex  user.deptid.depttype
	ExLoginCol map[string]string `xorm:"exlogincol" json:"exlogincol"`
	//自定义全局性配置，例如只需要指定客户启用或关闭的一些功能、信息在用户登录时，初始化在ExConfig参数
	ExConfig  map[string]string `xorm:"-" json:"ExConfig"`
	BudDeptId string            `xorm:"-" json:"buddeptid"` //归口部门编码
}

type ExLoginUser struct {
	UserID   string `xorm:"userid" json:"userid"`
	UserCode string `xorm:"usercode" json:"usercode"`
	UserName string `xorm:"username" json:"username"`
	DeptID   string `xorm:"deptid" json:"deptid"`     //所司部门
	DeptName string `xorm:"deptname" json:"deptname"` //所司部门
	SId      string `xorm:"sid" json:"sid"`           // 登录成功后生成的token 令牌
}
