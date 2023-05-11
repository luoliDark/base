package design

/**
 * @Author: lvxuanye
 * @Date: 2020/4/23 17:29
 * @describe:表单创建:内嵌-助手授权
 */
type Loadhelpbyrole struct {
	MsgName     string `xorm:"MsgName" json:"MsgName"`
	MsgCode     string `xorm:"MsgCode" json:"MsgCode"`
	MsgVsBillId string `xorm:"MsgVsBillId" json:"MsgVsBillId"`
	Pid         int    `xorm:"Pid" json:"Pid"`
}
