package constant

import "runtime"

/*
 公共使用的常量
*/

var CPUNumber = runtime.NumCPU()

const EmptyString = ""

//空的结构体 xrom有些方法必须传结构体，可传此空对象
var EmptyStruct struct{}

// 流程动作类型 审批类型 , 全部使用小写,统一匹配
const (
	Save       = "save"      // 保存
	Submit     = "submit"    //提交 反写:在途增加
	AppPass    = "apppass"   //审批通过 反写:在途减少，已完成增加
	AllPass    = "allpass"   //终审 反写:在途减少，已完成增加
	AppReturn  = "appreturn" //退回 反写:在途减少
	DisCard    = "discard"   //弃审 反写:已完成减少
	ResetFlow  = "resetflow" //重置流程
	ImportData = "importdata"
	// 拷贝类型
	M_M = "M_M" //主表——主表
	M_D = "M_D" //主表——子表
	D_D = "D_D" //子表——子表
	D_M = "D_M" //主表——子表
)

var StatusNameMap = map[string]string{
	Save:       "保存",
	Submit:     "提交",
	AppPass:    "审批",
	AllPass:    "终审",
	AppReturn:  "退回",
	DisCard:    "弃审",
	ResetFlow:  "重置流程",
	ImportData: "导入数据",
}

// 审批状态
const (
	FlowstatusSaveTemp        = 0  // 暂存
	FlowstatusSave            = 1  // 保存/待提交
	FlowstatusApproveing      = 2  // 审批中
	FlowstatusAllpass         = 3  // 已审核
	FlowstatusDiscard         = -1 // 弃审
	FlowstatusAppReturn       = -3 // 退回
	FlowstatusAppReturnModify = -9 // 退回待修改
)

//不检查的信封号
const NotCheckRealNumber = "99999999"
const RealNumberLength = 20

/*    用户相关    */
const LoginExCol = "loginexcol"

// GetStatusName 获取状态名称
func GetStatusName(a string) string {
	return StatusNameMap[a]
}
