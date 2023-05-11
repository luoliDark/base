package constant

/**
 * @author : Yix
 * @date 2022-06-09
 * @desc:
 **/

const CreateUidColName = "create_uid"
const UseridColName = "userid"

//自动拷贝提交，下游创建人取值方式，
//0 默认当前登录人， 1 上游创建人 create_uid 2 userid 申请人
const (
	GetSubmitWayByLoginID = iota
	GetSubmitWayByCreate_Uid
	GetSubmitWayByUserid
)

//固定不更新字段
var NotUpdateColMapByMain = map[string]interface{}{
	"billid":       nil,
	"billno":       nil,
	"flowstatus":   nil,
	"update_uid":   nil,
	"update_date":  nil,
	"create_uid":   nil,
	"create_date":  nil,
	"isdiscard":    nil,
	"discard_date": nil,
	"discard_uid":  nil,
	"currpid":      nil,
	"savesource":   nil,
	"entid":        nil,
	"isimport":     nil,
	"isautocreate": nil,
	"isqc":         nil,
	"newguid":      nil,
}

//固定不更新字段
var NotUpdateColMapByMainVsForm = map[string]interface{}{
	"billid":       nil,
	"billno":       nil,
	"flowstatus":   nil,
	"update_uid":   nil,
	"update_date":  nil,
	"create_uid":   nil,
	"create_date":  nil,
	"isdiscard":    nil,
	"discard_date": nil,
	"discard_uid":  nil,
	"approve_uid":  nil, //审批人
	"approve_date": nil,
	"currpid":      nil,
	"savesource":   nil,
	"entid":        nil,
	"isimport":     nil,
	"isautocreate": nil,
	"isqc":         nil,
	"newguid":      nil,
}

//关联带出字段，基础档案和表单字段名不一致时，通过分割符号，映射
var TargetRefColSplitOp = ":"
