package wf

// 审批通过struct
type AdjustStep struct {
	WaiteID      string
	FlowID       string
	Pid          int
	Primarykey   string
	Targetstepid string
	DynamicID    string
	NewGuid      string
	Appinion     string //调整说明
}
