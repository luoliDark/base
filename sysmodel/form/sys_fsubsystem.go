package form

type Sys_fsubsystem struct {
	SubSysID     int     `xorm:"SubSysID" json:"SubSysID"`
	SubSysCode   string  `xorm:"SubSysCode" json:"SubSysCode"`
	SubSysName   string  `xorm:"SubSysName" json:"SubSysName"`
	ReleaseVer   string  `xorm:"ReleaseVer" json:"ReleaseVer"`
	LogoPath     string  `xorm:"LogoPath" json:"LogoPath"`
	PidPre       int     `xorm:"PidPre" json:"PidPre"`
	Memo         string  `xorm:"Memo" json:"Memo"`
	SortID       float64 `xorm:"sortid" json:"sortid"`
	IsOpen       int     `xorm:"IsOpen" json:"IsOpen"`
	AdminCode    string  `xorm:"AdminCode" json:"AdminCode"`
	AdminIsFrame int     `xorm:"AdminIsFrame" json:"AdminIsFrame"`
}

func (*Sys_fsubsystem) TableName() string {
	return "sys_fsubsystem"
}
