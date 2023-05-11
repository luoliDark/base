package form

type Sys_billhelpmsgvsrole struct {
	MsgVsRoleId int    `xorm:"msgvsroleid" json:"msgvsroleid"`
	MsgVsBillId string `xorm:"msgvsbillid" json:"msgvsbillid"`
	RoleID      string `xorm:"roleid" json:"roleid"`
	CreateDate  string `xorm:"create_date" json:"create_date"`
	Memo        string `xorm:"memo" json:"memo"`
	CreateUid   string `xorm:"create_uid" json:"create_uid"`
}

func (*Sys_billhelpmsgvsrole) TableName() string {
	return "sys_billhelpmsgvsrole"
}
