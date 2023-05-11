package wf

/**
 * @Author: weiyg
 * @Date: 2020/3/18 23:19
 * @describe:查看审批信息 -结构体
 */
type Flowcharts struct {
	WaiteID        string `xorm:"WaiteID" json:"WaiteID"`
	SourceStepType string `xorm:"sourcesteptype" json:"sourcesteptype"`
	Sourcestepname string `xorm:"sourcestepname" json:"sourcestepname"`
	Targetstepname string `xorm:"Targetstepname" json:"Targetstepname"`
	DynamicID      int    `xorm:"DynamicID" json:"DynamicID"`
	IsActive       string `xorm:"IsActive" json:"IsActive"`
	Sourcestepid   string `xorm:"Sourcestepid" json:"Sourcestepid"`
	Targetstepid   string `xorm:"Targetstepid" json:"Targetstepid"`
	StepAttr       string `xorm:"stepattr" json:"stepattr"`
	EntId          int    `xorm:"entid" json:"entid"`
	Approvelog     []Sys_wfapprovelog
}
