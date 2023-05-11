package form

type Sys_fbtnvsform struct {
	BtnVSFormID string `xorm:"BtnVSFormID" json:"btnvsformid"`
	BtnCode     string `xorm:"BtnCode" json:"btncode"`
	//BtnText     string  `xorm:"btntext" json:"btntext"`
	Pid        int     `xorm:"Pid" json:"pid"`
	IsEditPage int     `xorm:"IsEditPage" json:"iseditpage"`
	SortID     float32 `xorm:"SortID" json:"sortid"`
	IsPubChk   int     `xorm:"IsPubChk" json:"ispubchk"`
	UseChkSql  string  `xorm:"UseChkSql" json:"usechksql"`
	UseChkMsg  string  `xorm:"UseChkMsg" json:"usechkmsg"`
	UseChkStr  string  `xorm:"UseChkStr" json:"usechkstr"`
	IsOpen     int     `xorm:"IsOpen" json:"isopen"`
}

func (*Sys_fbtnvsform) TableName() string {
	return "sys_fbtnvsform"
}
