package eb

type Eb_user struct {
	UserID         string `xorm:"userid" json:"userid"`
	CWSoftInnerUid string `xorm:"cwsoftinneruid" json:"cwsoftinneruid"`
	UserCode       string `xorm:"usercode" json:"usercode"`
	UserPwd        string `xorm:"userpwd" json:"userpwd"`
	UserName       string `xorm:"username" json:"username"`
	UserPhone      string `xorm:"userphone" json:"userphone"`
	UserEmail      string `xorm:"useremail" json:"useremail"`
	UserOpenID     string `xorm:"useropenid" json:"useropenid"`
	UserDingtalk   string `xorm:"userdingtalk" json:"userdingtalk"`
	UserWeiXintalk string `xorm:"userweixintalk" json:"userweixintalk"`
	OA_Logincode   string `xorm:"oa_logincode" json:"oa_logincode"` //OA的登录账号
	BankAdd        string `xorm:"bankadd" json:"bankadd"`
	CardNo         string `xorm:"cardno" json:"cardno"`
	DeptID         string `xorm:"deptid" json:"deptid"`
	Sex            string `xorm:"sex" json:"sex"`
	UserLevel      string `xorm:"userlevel" json:"userlevel"`
	BankAccount    string `xorm:"bankaccount" json:"bankaccount"`
	IsLock         string `xorm:"islock" json:"islock"`
	LockCount      int    `xorm:"lockcount" json:"lockcount"`
	LockTime       string `xorm:"locktime" json:"locktime"`
	IsFirstLogin   int    `xorm:"isfirstlogin" json:"isfirstlogin"`
	CompID         string `xorm:"compid" json:"compid"`
	Memo           string `xorm:"memo" json:"memo"`
	CreateUID      string `xorm:"create_uid" json:"create_uid"`
	CreateDate     string `xorm:"create_date" json:"create_date"`
	UpdateUID      string `xorm:"update_uid" json:"update_uid"`
	UpdateDate     string `xorm:"update_date" json:"update_date"`
	IsDiscard      int    `xorm:"isdiscard" json:"isdiscard"`
	DisCardUID     string `xorm:"discard_uid" json:"discard_uid"`
	DisCardDate    string `xorm:"discard_date" json:"discard_date"`
	GLCode         string `xorm:"gl_code" json:"gl_code"`
	NewGuID        string `xorm:"newguid" json:"newguid"`
	IsModify       int    `xorm:"ismodify" json:"ismodify"`
	CurrPID        int    `xorm:"currpid" json:"currpid"`
	SaveSource     string `xorm:"savesource" json:"savesource"`
	ImgSrc         string `xorm:"imgsrc" json:"imgsrc"`
	Status         int    `xorm:"status" json:"status"`
	EntId          int    `xorm:"entid" json:"entid"`
	Didiuserid     string `xorm:"didiuserid" json:"didiuserid"`
	OfficeAddId    string `xorm:"officeaddid" json:"officeaddid"`
	OfficeAdd      string `xorm:"officeadd" json:"officeadd"`
}

func (*Eb_user) TableName() string {
	return "eb_user"
}
