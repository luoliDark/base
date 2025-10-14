package restentity

const (
	//状态：  submit 提交 ， discard 弃审， appreturn 退回，apppass 审批事件 , allpass 终审 ，delete 删除，save  保存，beforesave 保存前
	Submit     = "submit"
	DisCard    = "discard"
	AppReturn  = "appreturn"
	Apppass    = "apppass"
	Allpass    = "allpass"
	Delete     = "delete"
	Save       = "save"
	BeforeSave = "beforesave"
)

// 事件对象参数
type EventPar struct {
	Pid          int
	Primarykey   string
	Newguid      string
	EntId        string
	Appinion     string
	Action       string // 状态：  submit ， discard ， appreturn ， allpass ，beforesave
	IsCheck      bool   //是否检查事件
	StepId       string //当前审批节点
	StepAttr     string //当前审批属性类型
	NextStepAttr string //下一节点审批属性类型
	FlowId       string //当前审批流程ID
	Sid          string
	IsApi        bool
	UserID       string
	Params       string            //其他参数：例如保存前事件，params=保存前对象数据
	FormConfig   map[string]string // 表单的个性化配置，启用了某项检查或功能配置  注： key的编码都是 小写
}
