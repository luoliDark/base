package sysmodel

type Eb_uservsrole struct {
	UserID         string `xorm:"userid" json:"userid"`
	RoleID         string `xorm:"roleid" json:"roleid"`
	EntId          int    `xorm:"entid" json:"entid"`
	IsFromDingDing int    `xorm:"isFromDingDing" json:"isFromDingDing"`
}

func (*Eb_uservsrole) TableName() string {
	return "eb_uservsrole"
}
