package wf

type Wfstepaccess struct {
	Type  string `xorm:"type" json:"type"`
	Name  string `xorm:"name" json:"name"`
	Id    string `xorm:"id" json:"id"`
	EntId int    `xorm:"entid" json:"entid"`
}
