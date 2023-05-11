package wf

type Sys_wfstepaccessdynamic struct {
	DynamicID        int    `xorm:"DynamicID pk autoincr " json:"DynamicID"`
	WaiteID          string `xorm:"WaiteID" json:"WaiteID"`
	FlowID           string `xorm:"FlowID" json:"FlowID"`
	Pid              int    `xorm:"Pid" json:"Pid"`
	BillID           string `xorm:"BillID" json:"BillID"`
	StepID           string `xorm:"StepID" json:"StepID"`
	AccObjID         string `xorm:"AccObjID" json:"AccObjID"`
	AccType          string `xorm:"AccType" json:"AccType"`
	NewGuid          string `xorm:"NewGuid" json:"NewGuid"`
	AccessSourcePid  int    `xorm:"access_source_pid" json:"access_source_pid"`
	AccessSourceId   string `xorm:"access_source_id" json:"access_source_id"`
	AccessSourceCol  string `xorm:"access_source_col" json:"access_source_col"`
	IsSign           int    `xorm:"IsSign" json:"IsSign"`
	IsConvert        int    `xorm:"IsConvert" json:"IsConvert"`
	OpUid            string `xorm:"OpUid" json:"OpUid"`
	IsFromStatic     int    `xorm:"IsFromStatic" json:"IsFromStatic"`
	IsActive         int    `xorm:"IsActive" json:"IsActive"`
	WaiteSignUserApp int    `xorm:"waitesignuserapp" json:"waitesignuserapp"`
	EntId            int    `xorm:"entid" json:"entid"`
	TargetIsUser     int    `xorm:"targetisuser" json:"targetisuser"`
}

func (*Sys_wfstepaccessdynamic) TableName() string {
	return "sys_wfstepaccessdynamic"
}
