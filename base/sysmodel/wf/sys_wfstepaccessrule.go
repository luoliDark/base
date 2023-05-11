package wf

type Sys_wfstepaccessrule struct {
	Accid               string `xorm:"Accid" 'AccID' json:"AccID"`
	Stepid              string `xorm:"Stepid" json:"StepID"`
	Iscodeget           int    `xorm:"Iscodeget" json:"IsCodeGet"`
	Isruleget           int    `xorm:"Isruleget" json:"IsRuleGet"`
	Sqlcode             string `xorm:"Sqlcode" json:"SqlCode"`
	Ruletype            string `xorm:"Ruletype" json:"RuleType"`
	Appusercol          string `xorm:"Appusercol" json:"AppUserCol"`
	Deptrule            string `xorm:"Deptrule" json:"DeptRule"`
	Deptrule2           string `xorm:"Deptrule2" json:"DeptRule2"`
	Deptappusercol      string `xorm:"Deptappusercol" json:"DeptAppUserCol"`
	Othercolrule        string `xorm:"Othercolrule" json:"OtherColRule"`
	Otherappusercol     string `xorm:"Otherappusercol" json:"OtherAppUserCol"`
	Deptrule2name       string `xorm:"Deptrule2name" json:"DeptRule2Name"`
	Appusercolname      string `xorm:"Appusercolname" json:"AppUserColName"`
	Deptappusercolname  string `xorm:"Deptappusercolname" json:"DeptAppUserColName"`
	Othercolrulename    string `xorm:"Othercolrulename" json:"OtherColRuleName"`
	Otherappusercolname string `xorm:"Otherappusercolname" json:"OtherAppUserColName"`
	FullName            string `xorm:"FullName" json:"FullName"`
	EntId               int    `xorm:"entid" json:"entid"`
}

func (*Sys_wfstepaccessrule) TableName() string {
	return "Sys_wfstepaccessrule"
}
