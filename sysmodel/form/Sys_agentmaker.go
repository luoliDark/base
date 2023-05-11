package form

/**
 * @describe:
 *
 * @Author: YiXin
 * @Date: 2021/2/2
 */

type Sys_agentmaker struct {
	ConsinID          int    `xorm:"consinid" json:"consinid"`
	AppUid            string `xorm:"appuid" json:"appuid"`                       //登录授权人
	ConsignTargentUid string `xorm:"consigntargentuid" json:"consigntargentuid"` //授权给谁
	ConsignStatus     string `xorm:"consignstatus" json:"consignstatus"`
	ConsignBeginTime  string `xorm:"consignbegintime" json:"consignbegintime"`
	ConsignEndTime    string `xorm:"consignendtime" json:"consignendtime"`
	EntId             int    `xorm:"entid" json:"entid"`
}

func (*Sys_agentmaker) TableName() string {
	return "sys_agentmaker"
}
