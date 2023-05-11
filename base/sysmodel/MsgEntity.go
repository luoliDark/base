package sysmodel

import "base/base/sysmodel/eb"

type MsgEntity struct {
	EntId             int          //企业ID
	Pid               int          //单据类型
	PrimaryKey        string       //主健
	PName             string       //单据类型
	BillNo            string       //流水号
	UserName          string       //提交人
	DeptName          string       //提交部门
	StepId            string       //审批节点ID
	TotalMoney        string       //总金额
	Memo              string       //备注
	IsWait            bool         //是否待办
	WaitTaskTmplateId string       //待办模板ID
	BillDate          string       //单据日期
	TemplateId        int          //模板ID
	Title             string       //标题
	Body              string       //文本
	WXBody            string       //微信文本
	LinkUrl           string       //连接地址
	PCLinkUrl         string       //PC端连接地址
	MsgType           string       //消息类型 cc抄送、push催办、info 普通消息
	ToUsers           []eb.Eb_user //收件人
}
