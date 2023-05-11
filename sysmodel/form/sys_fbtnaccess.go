package form

type Sys_fbtnaccess struct {
	BtnCid      int    `xorm:"BtnCid" json:"BtnCid"`
	BtnVSFormID string `xorm:"BtnVSFormID" json:"BtnVSFormID"`
	RoleID      string `xorm:"RoleID" json:"RoleID"`
	Pid         int    `xorm:"Pid" json:"Pid"`
	EntId       int    `xorm:"entid" json:"entid"`
}

func (*Sys_fbtnaccess) TableName() string {
	return "sys_fbtnaccess"
}
