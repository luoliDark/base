package design

/**
 * @Author: lvxuanye
 * @Date: 2020/4/23 14:17
 * @describe:表单创建:功能-内嵌授权
 */
type Loadbtnbyrole struct {
	BtnText     string `xorm:"BtnText" json:"BtnText"`
	BtnCode     string `xorm:"BtnCode" json:"BtnCode"`
	BtnVSFormID string `xorm:"BtnVSFormID" json:"BtnVSFormID"`
	Pid         int    `xorm:"Pid" json:"Pid"`
}
