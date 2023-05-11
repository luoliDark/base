package wf

type Sys_wfstepaccesspivot struct {
	AccPivotID        string `xorm:"accpivotid" json:"accpivotid"`
	Pid               int    `xorm:"pid" json:"pid"`
	GridId            int    `xorm:"gridid" json:"gridid"`
	IsMain            int    `xorm:"ismain" json:"ismain"`
	MainCols          string `xorm:"maincols" json:"maincols"`
	DSPid             int    `xorm:"dspid" json:"dspid"`
	DSSqlTable        string `xorm:"dssqltable" json:"dssqltable"`
	DSCols            string `xorm:"dscols" json:"dscols"`
	CType             string `xorm:"ctype" json:"ctype"`
	IsReadMain        int    `xorm:"isreadmain" json:"isreadmain"`
	Othertypemaincols string `xorm:"othertypemaincols" json:"othertypemaincols"`
	EntId             int    `xorm:"entid" json:"entid"`
}

func (*Sys_wfstepaccesspivot) TableName() string {
	return "sys_wfstepaccesspivot"
}
