package model

type Redis_RefreshVer struct {
	RefreId     int    `xorm:"RefreId" json:"RefreId"`
	KeyPre      string `xorm:"KeyPre" json:"KeyPre"`
	EntId       string `xorm:"entId" json:"entId"`
	Pid         string `xorm:"Pid" json:"Pid"`
	CType       string `xorm:"cType" json:"cType"`
	Version     string `xorm:"version" json:"version"`
	RefreshTime string `xorm:"refreshTime" json:"refreshTime"`
}

func (*Redis_RefreshVer) TableName() string {
	return "Redis_RefreshVer"
}
