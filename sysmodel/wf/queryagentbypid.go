package wf

/**
 * @Author: lvxuanye
 * @Date: 2020/4/26
 * @describe:审批代理授权（全部，部分，条件）
 */
type Queryagentbypid struct {
	ID           int    `xorm:"id" json:"id"`
	ConsinID     int    `xorm:"consinid" json:"consinid"`
	Pname        string `xorm:"pname" json:"pname"`
	PId          string `xorm:"pid" json:"pid"`
	Isbywhere    int    `xorm:"isbywhere" json:"isbywhere"`
	Create_Date  string `xorm:"create_date" json:"create_date"`
	Appuname     string `xorm:"appuname" json:"appuname"`
	Targentuname string `xorm:"targentuname" json:"targentuname"`
	EntId        int    `xorm:"entid" json:"entid"`
}
