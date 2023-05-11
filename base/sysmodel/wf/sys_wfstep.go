package wf

import "base/base/util/commutil"

// NotGetUserAction
const (
	StopSubmit               = "StopSubmit"               //禁止提交
	Skip                     = "Skip"                     //跳过本节点
	ChoiceBySubmit           = "ChoiceBySubmit"           //提交人手选
	ApprovedCurrentStepByGet = "ApprovedCurrentStepByGet" //审批到该节点时重新计算审批人
)

type Sys_wfstep struct {
	StepID                  string `xorm:"stepid" json:"stepid"`
	StepName                string `xorm:"stepname" json:"stepname"`
	FlowID                  string `xorm:"flowid" json:"flowid"`
	Pid                     int    `xorm:"pid" json:"pid"`
	StepSort                int    `xorm:"stepsort" json:"stepsort"`
	PassType                string `xorm:"passtype" json:"passtype"`
	ReturnRule              string `xorm:"returnrule" json:"returnrule"`
	OverTimeRule            string `xorm:"overtimerule" json:"overtimerule"`
	StepType                string `xorm:"steptype" json:"steptype"`
	CheckStepMsg            string `xorm:"checkstepmsg" json:"checkstepmsg"`
	IsChild                 int    `xorm:"ischild" json:"ischild"`
	RefFlowID               string `xorm:"refflowid" json:"refflowid"`
	HtmlCssText             string `xorm:"htmlcsstext" json:"htmlcsstext"`
	IsOpenEdit              int    `xorm:"isopenedit" json:"isopenedit"`
	IsQueueSend             int    `xorm:"isqueuesend" json:"isqueuesend"`
	IsOpenZH                int    `xorm:"isopenzh" json:"isopenzh"`
	IsOpenPush              int    `xorm:"isopenpush" json:"isopenpush"`
	IsSelectNexAppUser      int    `xorm:"isselectnexappuser" json:"isselectnexappuser"`
	IsPreAppUsersLeader     int    `xorm:"ispreappusersleader" json:"ispreappusersleader"`
	IsUseHistoryAppUser     int    `xorm:"isusehistoryappuser" json:"isusehistoryappuser"`
	IsOpenSign              int    `xorm:"isopensign" json:"isopensign"`
	IsOpenConvert           int    `xorm:"isopenconvert" json:"isopenconvert"`
	IsOpenRecover           int    `xorm:"isopenrecover" json:"isopenrecover"`
	StepAttr                string `xorm:"stepattr" json:"stepattr"`
	StepMemo                string `xorm:"stepmemo" json:"stepmemo"`
	SignUserWhere           string `xorm:"signuserwhere" json:"signuserwhere"`
	ApproveSkipType         string `xorm:"approveskiptype" json:"approveskiptype"`
	NotGetUserAction        string `xorm:"notgetuseraction" json:"notgetuseraction"` // 审批人为获取到时，处理方式
	IsSubmitSkip            int    `xorm:"issubmitskip" json:"issubmitskip"`
	IsStaticStep            int    `xorm:"isstaticstep" json:"isstaticstep"`
	IsReSubmitToThis        int    `xorm:"isresubmittothis" json:"isresubmittothis"`
	HtmlDivId               string `xorm:"htmldivid" json:"htmldivid"`
	IsFirstSkipByRepeatUser int    `xorm:"isfirstskipbyrepeatuser" json:"isfirstskipbyrepeatuser"`
	IsDiscard               int    `xorm:"isdiscard" json:"isdiscard"`
	IsCycleFindNextUser     int    `xorm:"iscyclefindnextuser" json:"iscyclefindnextuser"`
	StopFindByTopLevel      int    `xorm:"stopfindbytoplevel" json:"stopfindbytoplevel"`
	StopFindByRole          string `xorm:"stopfindbyrole" json:"stopfindbyrole"`
	IsContainRole           int    `xorm:"iscontainrole" json:"iscontainrole"`
	StopFindByRole_Show     string `xorm:"-"  json:"stopfindbyrole_show"`
	CopyFrom                string `xorm:"copyfrom" json:"copyfrom"`
	EntId                   int    `xorm:"entid" json:"entid"`
	MapWfUser               string `xorm:"mapwfuser" json:"mapwfuser"`
	ResetFlowAction         string `xorm:"resetflowaction" json:"resetflowaction"`
}

func (*Sys_wfstep) TableName() string {
	return "sys_wfstep"
}

type StepEntity struct {
	Step   Sys_wfstep
	Access []Sys_wfstepaccessdynamic
	CCUser []Sys_wfccinfo
}

func (this *StepEntity) GetStep(stepMap map[string]string) {
	this.Step.Pid = commutil.ToInt(stepMap["pid"])
	this.Step.StepSort = commutil.ToInt(stepMap["stepsort"])
	this.Step.IsChild = commutil.ToInt(stepMap["ischild"])
	this.Step.IsOpenEdit = commutil.ToInt(stepMap["isopenedit"])
	this.Step.IsQueueSend = commutil.ToInt(stepMap["isqueuesend"])
	this.Step.IsOpenZH = commutil.ToInt(stepMap["isopenzh"])
	this.Step.IsOpenPush = commutil.ToInt(stepMap["isopenpush"])
	this.Step.IsSelectNexAppUser = commutil.ToInt(stepMap["isselectnexappuser"])
	this.Step.IsPreAppUsersLeader = commutil.ToInt(stepMap["ispreappusersleader"])
	this.Step.IsUseHistoryAppUser = commutil.ToInt(stepMap["isusehistoryappuser"])
	this.Step.IsOpenSign = commutil.ToInt(stepMap["isopensign"])
	this.Step.IsOpenConvert = commutil.ToInt(stepMap["isopenconvert"])
	this.Step.IsOpenRecover = commutil.ToInt(stepMap["isopenrecover"])
	this.Step.IsSubmitSkip = commutil.ToInt(stepMap["issubmitskip"])
	this.Step.IsStaticStep = commutil.ToInt(stepMap["isstaticstep"])
	this.Step.IsReSubmitToThis = commutil.ToInt(stepMap["isresubmittothis"])
	this.Step.IsFirstSkipByRepeatUser = commutil.ToInt(stepMap["isfirstskipbyrepeatuser"])
	this.Step.IsDiscard = commutil.ToInt(stepMap["isdiscard"])
	this.Step.StepID = stepMap["stepid"]
	this.Step.StepName = stepMap["stepname"]
	this.Step.FlowID = stepMap["flowid"]
	this.Step.PassType = stepMap["passtype"]
	this.Step.ReturnRule = stepMap["returnrule"]
	this.Step.OverTimeRule = stepMap["overtimerule"]
	this.Step.StepType = stepMap["steptype"]
	this.Step.RefFlowID = stepMap["refflowid"]
	this.Step.HtmlCssText = stepMap["htmlcsstext"]
	this.Step.StepAttr = stepMap["stepattr"]
	this.Step.StepMemo = stepMap["stepmemo"]
	this.Step.SignUserWhere = stepMap["signuserwhere"]
	this.Step.ApproveSkipType = stepMap["approveskiptype"]
	this.Step.NotGetUserAction = stepMap["notgetuseraction"]
	this.Step.HtmlDivId = stepMap["htmldivid"]
	this.Step.CopyFrom = stepMap["copyfrom"]
	this.Step.EntId = commutil.ToInt(stepMap["entid"])
	this.Step.IsCycleFindNextUser = commutil.ToInt(stepMap["iscyclefindnextuser"])
	this.Step.StopFindByTopLevel = commutil.ToInt(stepMap["stopfindbytoplevel"])
	this.Step.StopFindByRole = stepMap["stopfindbyrole"]
	this.Step.IsContainRole = commutil.ToInt(stepMap["iscontainrole"])
	this.Step.MapWfUser = stepMap["mapwfuser"]
	this.Step.ResetFlowAction = stepMap["resetflowaction"]
	this.Step.CheckStepMsg = stepMap["checkstepmsg"]
}

func (this *Sys_wfstep) GetWFStep(stepMap map[string]string) {
	this.Pid = commutil.ToInt(stepMap["pid"])
	this.StepSort = commutil.ToInt(stepMap["stepsort"])
	this.IsChild = commutil.ToInt(stepMap["ischild"])
	this.CheckStepMsg = commutil.ToString(stepMap["checkstepmsg"])
	this.IsOpenEdit = commutil.ToInt(stepMap["isopenedit"])
	this.IsQueueSend = commutil.ToInt(stepMap["isqueuesend"])
	this.IsOpenZH = commutil.ToInt(stepMap["isopenzh"])
	this.IsOpenPush = commutil.ToInt(stepMap["isopenpush"])
	this.IsSelectNexAppUser = commutil.ToInt(stepMap["isselectnexappuser"])
	this.IsPreAppUsersLeader = commutil.ToInt(stepMap["ispreappusersleader"])
	this.IsUseHistoryAppUser = commutil.ToInt(stepMap["isusehistoryappuser"])
	this.IsOpenSign = commutil.ToInt(stepMap["isopensign"])
	this.IsOpenConvert = commutil.ToInt(stepMap["isopenconvert"])
	this.IsOpenRecover = commutil.ToInt(stepMap["isopenrecover"])
	this.IsSubmitSkip = commutil.ToInt(stepMap["issubmitskip"])
	this.IsStaticStep = commutil.ToInt(stepMap["isstaticstep"])
	this.IsReSubmitToThis = commutil.ToInt(stepMap["isresubmittothis"])
	this.IsFirstSkipByRepeatUser = commutil.ToInt(stepMap["isfirstskipbyrepeatuser"])
	this.IsDiscard = commutil.ToInt(stepMap["isdiscard"])
	this.StepID = stepMap["stepid"]
	this.StepName = stepMap["stepname"]
	this.FlowID = stepMap["flowid"]
	this.PassType = stepMap["passtype"]
	this.ReturnRule = stepMap["returnrule"]
	this.OverTimeRule = stepMap["overtimerule"]
	this.StepType = stepMap["steptype"]
	this.RefFlowID = stepMap["refflowid"]
	this.HtmlCssText = stepMap["htmlcsstext"]
	this.StepAttr = stepMap["stepattr"]
	this.StepMemo = stepMap["stepmemo"]
	this.SignUserWhere = stepMap["signuserwhere"]
	this.ApproveSkipType = stepMap["approveskiptype"]
	this.NotGetUserAction = stepMap["notgetuseraction"]
	this.HtmlDivId = stepMap["htmldivid"]
	this.CopyFrom = stepMap["copyfrom"]
	this.EntId = commutil.ToInt(stepMap["entid"])
}
