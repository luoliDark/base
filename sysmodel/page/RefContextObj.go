//单据上下文对象

package page

import (
	"base/sysmodel"
	"base/sysmodel/uservsent"
	"base/util/commutil"
	"reflect"
	"runingproject/services/public"
)

type RefContextObj struct {
	Pid            int              //单据类型
	PName          string           //单据类型名称
	PrimaryKey     string           //主健值
	BillNo         string           //编码 例：billno
	User           sysmodel.SSOUser //用户
	Enterprise     *uservsent.Eb_enterprise
	NewGuid        string //批次GUID
	FlowId         string //流程ID
	WaiteId        string //待审ID
	CurrStepId     string //当前节点id
	SaveSource     string //保存来源 pc or mobile
	TableName      string //表名
	PKCol          string //主健列名
	PkIds          string //多个主键值，针对于列表编辑保存使用，以, 分割
	NextAppUsers   string //手动选择的第一批审批人
	IsNextAppUsers string //是否包含手动审批人 二次提交
	ChooseStep     []map[string]string
	IsDataCenter   bool              //保存是否分发到数据平台
	IsShowPayState bool              //是否显示付款状态
	FormConfig     map[string]string // 单据个性化配置，非表单配置
	IsSign         string            // 是否加签
	SignList       string            // 加签人员名单
	IsCarbonCopy   string            // 是否抄送
	CarbonCopyList string            // 抄送人员名单
}

func (this *RefContextObj) SetFormConfig() {
	//获取单据个性化配置
	this.FormConfig = public.GetFormConfig(this.User, commutil.ToString(this.Pid))
}

func (this *RefContextObj) SetPid(Pid int) {
	//todo
	this.Pid = Pid
}
func (this *RefContextObj) SetUserId(UserId string) {
	//todo
	if reflect.DeepEqual(this.User, sysmodel.SSOUser{}) {
		this.User = sysmodel.SSOUser{}
	}
	this.User.UserID = UserId
}

func (this *RefContextObj) SetNewGuid(NewGuid string) {
	//todo
	this.NewGuid = NewGuid
}

func (this *RefContextObj) SetPkIds(pkIds string) {
	//todo
	this.PkIds = pkIds
}
