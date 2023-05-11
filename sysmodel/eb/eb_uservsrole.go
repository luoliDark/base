package eb

type Eb_UserVsRole struct {
	UserVsRoleId   string `xorm:"uservsroleid" json:"uservsroleid"`
	UserID         string `xorm:"userid" json:"userid"`
	RoleId         string `xorm:"roleid" json:"roleid"`
	EntId          string `xorm:"entid" json:"entid"`
	InsertDate     string `xorm:"insertdate" json:"insertdate"`
	IsFromDingDing string `xorm:"isfromdingding" json:"isfromdingding"`
}

func (*Eb_UserVsRole) TableName() string {
	return "eb_uservsrole"
}
