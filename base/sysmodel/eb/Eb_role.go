package eb

type Eb_role struct {
	RoleID         string `xorm:"roleid" json:"roleid"`
	SourceRoleID   string `xorm:"sourceroleid" json:"sourceroleid"`
	RoleCode       string `xorm:"rolecode" json:"rolecode"`
	RoleName       string `xorm:"rolename" json:"rolename"`
	CompID         string `xorm:"compid" json:"compid"`
	Memo           string `xorm:"memo" json:"memo"`
	CreateUID      string `xorm:"create_uid" json:"create_uid"`
	CreateDate     string `xorm:"create_date" json:"create_date"`
	UpdateUID      string `xorm:"update_uid" json:"update_uid"`
	UpdateDate     string `xorm:"update_date" json:"update_date"`
	IsDiscard      int    `xorm:"isdiscard" json:"isdiscard"`
	DisCardUID     string `xorm:"discard_uid" json:"discard_uid"`
	DisCardDate    string `xorm:"discard_date" json:"discard_date"`
	SaveSource     string `xorm:"savesource" json:"savesource"`
	Groupid        int    `xorm:"groupid" json:"groupid"`
	Groupname      string `xorm:"groupname" json:"groupname"`
	Ispub          int    `xorm:"ispub" json:"ispub"`
	Entid          int    `xorm:"entid" json:"entid"`
	IsFromDingDing int    `xorm:"isFromDingDing" json:"isFromDingDing"`
}

func (*Eb_role) TableName() string {
	return "eb_role"
}
