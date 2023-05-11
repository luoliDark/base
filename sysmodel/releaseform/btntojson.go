package releaseform

//主表按钮 用于生成json
type BtnJson struct {
	BtnCode       string `xorm:"BtnCode" json:"BtnCode"`
	BtnText       string `xorm:"BtnText" json:"BtnText"`
	PageStateShow string `xorm:"PageStateShow" json:"PageStateShow"`
	IsSysBtn      int    `xorm:"IsSysBtn" json:"IsSysBtn"`
	IsEditPage    int    `xorm:"IsEditPage" json:"IsEditPage"`
	Pid           int    `xorm:"Pid" json:"Pid"`
	IsChkVer      int    `xorm:"IsChkVer" json:"IsChkVer"`
}

//子表按钮 用于生成json
type GridBtns struct {
	GridID   int    `xorm:"GridID" json:"GridID"`
	Pid      int    `xorm:"Pid" json:"Pid"`
	Toolbars string `xorm:"Toolbars" json:"Toolbars"`
}

//子表某个按钮 用于生成json
type GridBtn struct {
	BtnCode string `xorm:"BtnCode" json:"BtnCode"`
	BtnText string `xorm:"BtnText" json:"BtnText"`
}
