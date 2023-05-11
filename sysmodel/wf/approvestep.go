package wf

// 审批通过struct
type ApproveStep struct {
	StepId                  string              //当前审批节点
	StepDynId               int                 //动态节点表主健ID
	NextStepLeader          []map[string]string //本节点审批人直属领导作为下一级节点审批人
	ApproveSkipType         string              //审批跳过类型  ByNear  相邻节点  ByOverStep  跨节点
	PassType                string              //通过类型  QZ或BS
	IsFirstSkipByRepeatUser string              //当和后面节点审批人重复时，先对本节点进行跳过
	IsSystemAutoApprove     string              //机器人审批节点 禁止人工审批
	IsAgent                 bool                //是否代理审批
	SourceAction            string              //审批权限来源 Sign   加签 Convert 转批 Agent 代理授权 TimeOut 超时转批
	OriginAppUid            string              //原始审批人
	NextSelectAppUser       string              //手工选择的下级审批人
	NextSelectStepId        string              //手工选择的下级节点ID
	IsSelectNexAppUser      int
	Appinion                string //意见
	Terminal                string //终端
	StepAttr                string //节点类型
	NextStepAttr            string //下一节点类型
	NextStepType            string //下一节点类型 cc,sumbmit,end等
	EntId                   int    `xorm:"entid" json:"entid"`
	TargetIsUser            int
	IsSubmitSkip            int
	NextIsEnd               bool   //下一步是结束节点 现在没有用 因为没取值
	IsByPreSelectUser       bool   //当前节点审批人是由上级审批人手选的
	PreSelectUsers          string //上级审批人选择的用户ID
}
