package wf

/**
 * @Author: lvxuanye
 * @Date: 2020/4/24
 * @describe:审批代理授权（全部，部分，条件）
 */
type Queryagent struct {
	ConsinID          int    `xorm:"ConsinID" json:"consinid"`
	AppUname          string `xorm:"appuname" json:"appuname"`
	AppId             string `xorm:"appuid" json:"appuid"`
	Memo              string `xorm:"Memo" json:"memo"`
	TargentUname      string `xorm:"TargentUname" json:"targentuname"`
	ConsignTargentUid string `xorm:"consigntargentuid" json:"consigntargentuid"`
	ConsignStatus     int    `xorm:"ConsignStatus" json:"consignstatus"`
	ConsignBeginTime  string `xorm:"ConsignBeginTime" json:"consignbegintime"`
	ConsignEndTime    string `xorm:"ConsignEndTime" json:"consignendtime"`
	IsByPid           int    `xorm:"IsByPid" json:"isbypid"`
	CreateDate        string `xorm:"Create_Date" json:"createdate"`
	EntId             int    `xorm:"entid" json:"entid"`
}
