package wf

import "base/base/sysmodel/matrix"

type QuerystepCCuser struct {
	AccID      int    `xorm:"accid" json:"accid"`
	StepID     string `xorm:"stepid" json:"stepid"`
	CCUserID   string `xorm:"ccuserid" json:"ccuserid"`
	Flowid     string `xorm:"flowid" json:"flowid"`
	Ccusername string `xorm:"ccusername" json:"ccusername"`
	EntId      int    `xorm:"entid" json:"entid"`
}

type QuerystepCCInfo struct {
	CCUser    []QuerystepCCuser        `xorm:"ccuser" json:"ccuser"`
	MatrixOjb []matrix.Sys_matrixvspid `xorm:"matrixojb" json:"matrixojb"`
}
